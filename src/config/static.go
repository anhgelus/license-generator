package config

import (
	"license-generator/src/utils"
	"os"
	"errors"
)

func ImportStaticConfig() (*[]*LicenseConfig, error) {
	dirPath, err := os.UserHomeDir()
	utils.HandleError(err)
	configPath := dirPath + "/.config/license-generator/"
	err = os.Mkdir(configPath, 777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}
	l := new([]*LicenseConfig)
	err = GetLicenseConfigs(configPath, l)
	return l, err
}
