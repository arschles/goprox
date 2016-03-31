package githttp

import (
	"regexp"
)

// Routing regexes
var (
	rpcUploadPath         = regexp.MustCompile("(.*?)/git-upload-pack$")
	rpcReceivePath        = regexp.MustCompile("(.*?)/git-receive-pack$")
	getInfoRefsPath       = regexp.MustCompile("(.*?)/info/refs$")
	getHeadPath           = regexp.MustCompile("(.*?)/HEAD$")
	getAlternatesPath     = regexp.MustCompile("(.*?)/objects/info/alternates$")
	getHTTPAlternatesPath = regexp.MustCompile("(.*?)/objects/info/http-alternates$")
	getInfoPacksPath      = regexp.MustCompile("(.*?)/objects/info/packs$")
	getInfoFilePath       = regexp.MustCompile("(.*?)/objects/info/[^/]*$")
	getLooseObjectPath    = regexp.MustCompile("(.*?)/objects/[0-9a-f]{2}/[0-9a-f]{38}$")
	getPackFilePath       = regexp.MustCompile("(.*?)/objects/pack/pack-[0-9a-f]{40}\\.pack$")
	getIdxFilePath        = regexp.MustCompile("(.*?)/objects/pack/pack-[0-9a-f]{40}\\.idx$")
)

// Service is the service for a git RPC
type Service struct {
	Method  string
	Handler func(HandlerReq) error
	RPC     string
}

func services(g *GitHTTP) map[*regexp.Regexp]Service {
	return map[*regexp.Regexp]Service{
		rpcUploadPath:         Service{"POST", g.serviceRPC, uploadPack},
		rpcReceivePath:        Service{"POST", g.serviceRPC, receivePack},
		getInfoRefsPath:       Service{"GET", g.getInfoRefs, ""},
		getHeadPath:           Service{"GET", g.getTextFile, ""},
		getAlternatesPath:     Service{"GET", g.getTextFile, ""},
		getHTTPAlternatesPath: Service{"GET", g.getTextFile, ""},
		getInfoPacksPath:      Service{"GET", g.getInfoPacks, ""},
		getInfoFilePath:       Service{"GET", g.getTextFile, ""},
		getLooseObjectPath:    Service{"GET", g.getLooseObject, ""},
		getPackFilePath:       Service{"GET", g.getPackFile, ""},
		getIdxFilePath:        Service{"GET", g.getIdxFile, ""},
	}
}

// getServiceForPath return's the service corresponding to the current http.Request's URL
// as well as the name of the repo
func getServiceForPath(g *GitHTTP, path string) (string, *Service) {
	for re, service := range services(g) {
		if m := re.FindStringSubmatch(path); m != nil {
			return m[1], &service
		}
	}

	// No match
	return "", nil
}
