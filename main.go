package main

import (
	"embed"
	"flag"
	"github.com/anhgelus/license-generator/src/args"
	"github.com/anhgelus/license-generator/src/config"
	"github.com/anhgelus/license-generator/src/utils"
	"os"
	"strings"
	"text/template"
)

//go:embed resources/template/license
var staticContent embed.FS

func init() {
	flag.StringVar(&args.Name, "name", "", "Set the name of the project")
	flag.StringVar(&args.LicenseSelected, "license", "", "Set the license of the project\"")
	flag.StringVar(&args.Years, "years", "", "Set the year")
	flag.StringVar(
		&args.Authors,
		"authors",
		"",
		"Set the authors of the project, separate them with the coma (,)",
	)
	flag.StringVar(
		&args.CustomConfigPath,
		"config-path",
		"",
		"Set the path to the config to use custom licenses",
	)
	flag.BoolVar(&args.List, "l", false, "List every available license")
	flag.BoolVar(&args.Help, "h", false, "Show the help")
	flag.BoolVar(&args.Verbose, "v", false, "Verbose")
}

func main() {
	// set the global context path
	var err error
	utils.ContextPath, err = os.Getwd()
	utils.HandleError(err)
	utils.ContextPath += "/"

	flag.Parse()

	// import the basic config
	_, err = config.ImportStaticConfig()
	if err != nil {
		println(err.Error())
	}

	// parse args
	arg := args.ParseCliArgs()
	// import custom licenses if needed
	if arg.ConfigPath != "" {
		l, err := config.GetLicenseConfigs(arg.ConfigPath)
		utils.HandleError(err)
		config.AddLicensesToMap(l, arg.ConfigPath)
		if args.Verbose {
			println("")
		}
	}

	// show the info arguments and exist
	if arg.Info {
		println(arg.InfoText())
		os.Exit(0)
	}

	arg.HandleArgs()

	l := findLicense(arg.LicenseType)
	file := parseLicense(arg, l)
	err = os.WriteFile("LICENSE", []byte(file), 0766)
	utils.HandleError(err)
	println("The LICENSE was successfully created!")
}

func findLicense(license *args.License) string {
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
		"Years":   arg.Years,
		"Authors": utils.StringArrayToString(arg.Authors),
		"AppName": arg.AppName,
	})
	utils.HandleError(err)
	return tempWriter.String()
}
