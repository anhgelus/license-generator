package args

import (
	"flag"
	"github.com/anhgelus/license-generator/src/utils"
	"os"
)

var (
	Name             string
	LicenseSelected  string
	Authors          string
	Years            string
	CustomConfigPath string
	Help             bool
	List             bool
)

func ParseCliArgs() *Arguments {
	flag.Parse()
	// if the user don't give basic information through args
	if len(os.Args) < 2 {
		return &Arguments{Question: true}
	}
	arguments := &Arguments{Info: false}

	if Help || List {
		arguments.Info = true
		if Help {
			arguments.InfoText = HelpArg.textGenerator
		} else if List {
			arguments.InfoText = LicenseListArg.textGenerator
		}
		return arguments
	}

	arguments.AppName = Name
	arguments.LicenseType.Name = LicenseSelected
	arguments.Authors = parseAuthors(Authors)
	arguments.Years = Years
	arguments.Question = len(Name) == 0 || len(Years) == 0 || len(Authors) == 0 || len(LicenseSelected) == 0

	return arguments
}

func (arg *Arguments) HandleArgs() {
	if arg.Question {
		arg.handleQuestion()
	}
	m := make(map[string]string)
	m["App Name"] = arg.AppName
	m["License"] = arg.LicenseType.Name
	m["Author(s)"] = utils.StringArrayToString(arg.Authors)
	m["Years(s)"] = arg.Years
	utils.GenerateSumeUp("Options", m, "-")
}

func (arg *Arguments) handleQuestion() {
	if arg.AppName == "" {
		name := ""
		print("Name of the application: ")
		err := utils.Scan(&name)
		utils.HandleError(err)
		arg.AppName = name
	}
	if arg.LicenseType.Name == "" {
		oldLicense := ""
		print("License: ")
		err := utils.Scan(&oldLicense)
		utils.HandleError(err)
		license, found := GetLicense(oldLicense)
		if !found {
			println("Unknown license type. Check available licenses with -l.")
			os.Exit(1)
		}
		println("License found: " + license.Name)
		arg.LicenseType = license
	}
	if len(arg.Authors) == 0 {
		unparsedAuthor := ""
		print("Author(s) of this program (separate each author with a coma (,)): ")
		err := utils.Scan(&unparsedAuthor)
		utils.HandleError(err)
		arg.Authors = parseAuthors(unparsedAuthor)
	}
	if arg.Years == "" {
		year := ""
		print("Years of the copyright: ")
		err := utils.Scan(&year)
		utils.HandleError(err)
		arg.Years = year
	}
}
