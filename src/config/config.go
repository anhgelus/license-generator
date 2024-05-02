package config

import (
	"errors"
	"fmt"
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

func GetLicenseConfigs(path string, licenses *[]*LicenseConfig) error {
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
		return nil
	}
	if config.PathToLicenses == "" {
		return errors.New("impossible to find the config file with this path: " + path)
	}
	return config.parseLicensesFile(utils.RelativeToAbsolute(config.PathToLicenses, path), licenses)
}

// parseLicensesFile return every LicenseConfig
func (config *MainConfig) parseLicensesFile(path string, licenses *[]*LicenseConfig) error {
	mapLicenses := make(map[string]LicenseConfig)
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
		mapLicenses[license.Identifier] = license
		fmt.Printf("Imported %s (%s)\n", license.Name, license.Identifier)
	}
	final := make([]*LicenseConfig, len(mapLicenses))
	if len(config.ListConfig) == 0 {
		i := 0
		for _, licenseConfig := range mapLicenses {
			final[i] = &licenseConfig
			i++
		}
		*licenses = append(final, *licenses...)
		return nil
	}
	i := 0
	for _, s := range config.ListConfig {
		v, found := mapLicenses[s]
		if !found {
			return errors.New("the license with the identifier " + s + " was not found")
		}
		final[i] = &v
		i++
	}
	*licenses = append(final, *licenses...)
	return nil
}

func (license *LicenseConfig) AddToMap(contextPath string) {
	args.AddLicense(license.ToLicense(contextPath), license.Identifier)
}

func AddLicensesToMap(licenses *[]*LicenseConfig, contextPath string) {
	for _, license := range *licenses {
		license.AddToMap(contextPath)
	}
}

func (license *LicenseConfig) ToLicense(contextPath string) args.License {
	return args.License{
		Name: license.Name,
		File: utils.RelativeToAbsolute(license.Path, contextPath),
	}
}
