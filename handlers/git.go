package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/arschles/goprox/githttp"
	"github.com/arschles/goprox/storage"
	s3 "github.com/minio/minio-go"
)

func gitClone(repoName, repoDir string) error {
	cmd := exec.Command("git", "clone", fmt.Sprintf("https://%s", repoName), repoDir)
	cmd.Dir = repoDir
	log.Printf("executing %s in %s", strings.Join(cmd.Args, " "), cmd.Dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("error cloning (%s)", err)
		return err
	}
	log.Printf("done filling %s", repoName)
	return nil
}

// NewGit returns the handler to be used for the Git server
func NewGit(s3Client *s3.Client, bucketName, tmpDir string) http.Handler {
	hdl := githttp.New(tmpDir)
	// hdl.UploadPack = false
	hdl.FillRepo = func(repoDir string) error {
		if !strings.HasPrefix(repoDir, tmpDir) {
			return fmt.Errorf("invalid repoDir in FillRepo (%s)", repoDir)
		}
		packageName := repoDir[len(tmpDir)+1:]
		if err := os.MkdirAll(repoDir, os.ModePerm); err != nil {
			log.Printf("error creating %s (%s)", repoDir, err)
			return err
		}
		defer func() {
			if err := os.RemoveAll(repoDir); err != nil {
				log.Printf("Error removing repository for %s in %s (%s)", packageName, repoDir, err)
			}
		}()

		objInfo, err := storage.PackageExists(s3Client, bucketName, packageName)
		needsDownload := true
		if err != nil {
			if err := gitClone(packageName, repoDir); err != nil {
				log.Printf("Error git cloning %s (%s)", packageName, err)
				return err
			}
			if err := storage.UploadPackage(s3Client, bucketName, packageName, repoDir); err != nil {
				log.Printf("Error uploading package %s from %s (%s)", packageName, repoDir, err)
			}
			needsDownload = false
		}

		if needsDownload {
			obj, err := s3Client.GetObject(bucketName, objInfo.Key)
			if err != nil {
				log.Printf("Error downloading %s from bucket %s (%s)", objInfo.Key, bucketName, err)
				return err
			}
			if err := storage.UntarObjectToDisk(obj, repoDir); err != nil {
				log.Printf("Error untarring %s to %s (%s)", objInfo.Key, repoDir, err)
				return err
			}
		}
		return nil
	}
	return hdl
}
