package config

import (
	"errors"
	"license-generator/src/args"
	"license-generator/src/utils"
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

func GetLicenseConfigs(path string) ([]LicenseConfig, error) {
	files, err := os.ReadDir(path)
	utils.HandleError(err)

	var config MainConfig
	for _, file := range files {
		// remove directories and non toml files
		if file.IsDir() || !regex.MatchString(file.Name()) {
			continue
		}
		// remove other file than config .toml
		if file.Name() != "config.toml" {
			continue
		}
		content := utils.FileContent(path, file)
		utils.DecodeToml(content, &config)
		println("Found the config.toml file")
	}
	if config.PathToLicenses == "" {
		return nil, errors.New("impossible to find the config file with this path: " + path)
	}
	return config.parseLicensesFile(utils.RelativeToAbsolute(config.PathToLicenses, path))
}

// parseLicensesFile return every LicenseConfig
func (config *MainConfig) parseLicensesFile(path string) ([]LicenseConfig, error) {
	licenses := make(map[string]LicenseConfig)
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
		licenses[license.Identifier] = license
		println("Imported", license.Name)
	}
	final := make([]LicenseConfig, len(licenses))
	for _, l := range licenses {
		println("Name:", l.Name, l.Identifier)
	}
	if len(config.ListConfig) == 0 {
		i := 0
		for _, licenseConfig := range licenses {
			final[i] = licenseConfig
			i++
		}
		return final, nil
	}
	i := 0
	for _, s := range config.ListConfig {
		v, found := licenses[s]
		if !found {
			return nil, errors.New("the license with the identifier " + s + " was not found")
		}
		final[i] = v
		i++
	}
	return final, nil
}

func (license *LicenseConfig) AddToMap(contextPath string) {
	args.AddLicense(license.ToLicense(contextPath), license.Identifier)
}

func AddLicensesToMap(licenses []LicenseConfig, contextPath string) {
	for _, license := range licenses {
		license.AddToMap(contextPath)
	}
}

func (license *LicenseConfig) ToLicense(contextPath string) args.License {
	return args.License{
		Name: license.Name,
		File: utils.RelativeToAbsolute(license.Path, contextPath),
	}
}
