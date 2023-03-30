package config

import (
	"license-generator/src/utils"
	"os"
)

func ImportStaticConfig() ([]LicenseConfig, error) {
	dirPath, err := os.UserHomeDir()
	utils.HandleError(err)
	configPath := dirPath + "/.config/license-generator/"
	return GetLicenseConfigs(configPath)
}
