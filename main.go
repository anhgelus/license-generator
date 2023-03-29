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
	// import custom licenses if needed
	if arg.ConfigPath != "" {
		licenses, err := config.GetLicenseConfigs(arg.ConfigPath)
		utils.HandleError(err)
		// set the global context path
		utils.ContextPath, err = os.Getwd()
		utils.HandleError(err)
		utils.ContextPath += "/"
		config.AddLicensesToMap(licenses, arg.ConfigPath)
		println("")
	}
	// show the info arguments and exist
	if arg.Info {
		println(arg.InfoText())
		os.Exit(0)
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
	b, err := os.ReadFile(license.File)
	utils.HandleError(err)
	return string(b)
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
