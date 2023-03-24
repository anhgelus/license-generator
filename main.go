package main

import (
	"embed"
	"license-generator/src/args"
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
	arg.HandleArgs()

	l := findLicense(arg.LicenseType)
	file := parseLicense(arg, l)
	err := os.WriteFile("LICENSE", []byte(file), 0666)
	utils.HandleError(err)
	println("The LICENSE was successfully created!")
}

func findLicense(license args.License) string {
	content, err := staticContent.ReadFile("resources/template/license/" + string(license))
	utils.HandleError(err)
	return string(content)
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
