package main

import (
	"embed"
	"license-generator/src/args"
	"license-generator/src/config"
	"license-generator/src/utils"
	"os"
	"strings"
	"text/template"
)

//go:embed resources/template/license
var staticContent embed.FS

func main() {
	args.GenerateLicenseMap()
	arg := args.ParseCliArgs()
	if arg.Info {
		os.Exit(0)
	}
	// import custom licenses if needed
	if arg.ConfigPath != "" {
		licenses, err := config.GetLicenseConfigs(arg.ConfigPath)
		utils.HandleError(err)
		contextPath, err := os.Getwd()
		utils.HandleError(err)
		println(contextPath)
		for _, license := range licenses {
			println("Path in main.go:", license.Path)
		}
		config.AddLicensesToMap(licenses, contextPath)
	}
	arg.HandleArgs()

	l := findLicense(arg.LicenseType)
	file := parseLicense(arg, l)
	err := os.WriteFile("LICENSE", []byte(file), 0666)
	utils.HandleError(err)
	println("The LICENSE was successfully created!")
}

func findLicense(license args.License) string {
	if license.File == "~" {
		content, err := staticContent.ReadFile("resources/template/license/" + license.Name)
		utils.HandleError(err)
		return string(content)
	}
	return ""
}

func parseLicense(arg *args.Arguments, license string) string {
	t, err := template.New("index.html").Parse(license)
	utils.HandleError(err)

	tempWriter := new(strings.Builder)
	err = t.Execute(tempWriter, map[string]interface{}{
		"Year":    arg.Year,
		"Authors": utils.StringArrayToString(arg.Authors),
		"AppName": arg.AppName,
	})
	utils.HandleError(err)
	return tempWriter.String()
}
