package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Data is any data parsed from a Yaml file
type Data interface {
	GetData(env string) error
}

// GetExecPath gets the Path of the executable file
func GetExecPath() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	execPath, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return "", err
	}
	return execPath, nil
}

// GetPath gets the full path of the file
func GetPath(filename string) (string, error) {
	execPath, err := GetExecPath()
	if err != nil {
		return "", err
	}
	PathList := []string{filepath.Dir(execPath), filename}
	Path := strings.Join(PathList, "")
	return Path, nil
}

// parseYamlFile parses the file into the config struct
func parseFile(env string, ylD Data) error {

	fileName := fmt.Sprint("/assets/env/" + env + ".yaml")
	
	Path, err := GetPath(fileName)
	if err != nil {
		return fmt.Errorf("Cannot get %v path", err)
	}

	File, err := ioutil.ReadFile(Path)

	if err = yaml.Unmarshal(File, ylD); err != nil {
		return fmt.Errorf("Cannot parse config, %v", err)
	}
	return nil
}
