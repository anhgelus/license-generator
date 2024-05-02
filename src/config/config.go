package config

import (
	"errors"
	"fmt"
	"github.com/anhgelus/license-generator/src/args"
	"github.com/anhgelus/license-generator/src/utils"
	"os"
	"regexp"
)

type MainConfig struct {
	ListConfig     []string `toml:"customLicenses"`
	PathToLicenses string   `toml:"pathToLicenses"`
}

type LicenseConfig struct {
	Path       string `toml:"path"`
	Name       string `toml:"name"`
	Identifier string `toml:"identifier"`
}

var (
	regex = regexp.MustCompile(".*.toml")
)

func GetLicenseConfigs(path string) ([]*LicenseConfig, error) {
	files, err := os.ReadDir(path)
	utils.HandleError(err)

	var config MainConfig
	found := false
	for _, file := range files {
		// remove directories and non toml files
		if file.IsDir() || !regex.MatchString(file.Name()) {
			continue
		}
		// remove other file than config .toml
		if file.Name() != "config.toml" || found {
			continue
		}
		content := utils.FileContent(path, file)
		utils.DecodeToml(content, &config)
		println("Found the config.toml file")
		found = true
	}
	if !found {
		return nil, nil
	}
	if config.PathToLicenses == "" {
		return nil, errors.New("impossible to find the config file with this path: " + path)
	}
	return config.parseLicensesFile(utils.RelativeToAbsolute(config.PathToLicenses, path))
}

// parseLicensesFile return every LicenseConfig
func (config *MainConfig) parseLicensesFile(path string) ([]*LicenseConfig, error) {
	var licenses []*LicenseConfig
	files, err := os.ReadDir(path)
	utils.HandleError(err)
	for _, file := range files {
		// remove directories and non toml files
		if file.IsDir() || !regex.MatchString(file.Name()) {
			continue
		}
		// reject every config
		if file.Name() == "config.toml" {
			continue
		}
		content := utils.FileContent(path, file)
		var license LicenseConfig
		utils.DecodeToml(content, &license)
		licenses = append(licenses, &license)
		fmt.Printf("Imported %s (%s)\n", license.Name, license.Identifier)
	}
	AddLicensesToMap(licenses, path+"/")
	return licenses, nil
}

func (license *LicenseConfig) AddToMap(contextPath string) {
	args.AddLicense(license.ToLicense(contextPath), license.Identifier)
}

func AddLicensesToMap(licenses []*LicenseConfig, contextPath string) {
	for _, license := range licenses {
		license.AddToMap(contextPath)
	}
}

func (license *LicenseConfig) ToLicense(contextPath string) *args.License {
	return &args.License{
		Name: license.Name,
		File: utils.RelativeToAbsolute(license.Path, contextPath),
	}
}
