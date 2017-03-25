package deps

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

const (
	// DefaultFileName is the default name of the dependencies file in the current
	// working directory.
	DefaultFileName = "goprox.yaml"
)

// File is the yaml-tagged struct reprentation of a dependencies file
type File struct {
	Package  string   `yaml:"package"`
	Homepage string   `yaml:"homepage"`
	License  string   `yaml:"license"`
	Owners   []Owner  `yaml:"owners"`
	Import   []Import `yaml:"import"`
}

// Owner is the yaml-tagged struct representation of the owner portion of a dependencies file
type Owner struct {
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
	Homepage string `yaml:"homepage"`
}

// Import is the yaml-tagged struct representation of an import item in a dependencies file
type Import struct {
	Package string `yaml:"package"`
	Version string `yaml:"version"`
}

// ParseFile parses the file at fName into a File, or returns a non-nil error if the file
// was unparseable
func ParseFile(fName string) (*File, error) {
	fileBytes, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	ret := new(File)
	if err := yaml.Unmarshal(fileBytes, ret); err != nil {
		return nil, err
	}
	return ret, nil

	return nil, nil
}
