package config

import (
	"errors"
	"os"
)

func ImportStaticConfig() ([]*LicenseConfig, error) {
	dirPath := os.Getenv("XDG_CONFIG_HOME")
	if dirPath == "" {
		dirPath = os.Getenv("HOME") + "/.config"
	}
	configPath := dirPath + "/license-generator/"
	err := os.Mkdir(configPath, 777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}
	l, err := GetLicenseConfigs(configPath)
	return l, err
}
