package args

import (
	"github.com/anhgelus/license-generator/src/utils"
	"os"
	"strings"
)

func ParseCliArgs() *Arguments {
	args := os.Args
	// if the user don't give basic information through args
	if len(args) < 2 {
		return &Arguments{Question: true}
	}
	arguments := Arguments{Info: false}
	l := 0
	tb := false

ar:
	for i, arg := range args {
		if i == 0 {
			continue
		}
		if arg == "-h" {
			arguments.InfoText = HelpArg.textGenerator
			tb = true
		}
		for _, av := range argLists {
			if arg != av.GenerateParameter() {
				continue
			}
			arguments.assignValueToArguments(&av, args[i+1])
			l++
		}
		for _, av := range infoArgLists {
			if arg != av.GenerateParameter() {
				continue
			}
			arguments.InfoText = av.textGenerator
			l++
			tb = true
		}
		if tb {
			break ar
		}
	}
	if l != 4 {
		arguments.Question = true
	}
	if tb {
		arguments.Info = true
	}
	return &arguments
}

func (arg *Arguments) HandleArgs() {
	if arg.Question {
		arg.handleQuestion()
	}
	m := make(map[string]string)
	m["App Name"] = arg.AppName
	m["License"] = arg.LicenseType.Name
	m["Author(s)"] = utils.StringArrayToString(arg.Authors)
	m["Year(s)"] = arg.Year
	utils.GenerateSumeUp("Options", m, "-")
}

func (arg *Arguments) handleQuestion() {
	if arg.AppName == "" {
		name := ""
		print("Name of the application: ")
		err := utils.Scan(&name)
		utils.HandleError(err)
		println("The name is: " + name)
		arg.AppName = name
	}
	if arg.LicenseType.Name == "" {
		oldLicense := ""
		print("License: ")
		err := utils.Scan(&oldLicense)
		utils.HandleError(err)
		license, found := GetLicense(oldLicense)
		if !found {
			println("Unknown license type. Aborted.")
			os.Exit(2)
		}
		println("The license is: " + license.Name)
		arg.LicenseType = license
	}
	if len(arg.Authors) == 0 {
		unparsedAuthor := ""
		print("Author(s) of this program (separate each author with a coma (,)): ")
		err := utils.Scan(&unparsedAuthor)
		utils.HandleError(err)
		println("Authors: " + strings.ReplaceAll(unparsedAuthor, ",", ", "))
		arg.Authors = parseAuthors(unparsedAuthor)
	}
	if arg.Year == "" {
		year := ""
		print("Year of the copyright: ")
		err := utils.Scan(&year)
		utils.HandleError(err)
		println("The year is: " + year)
		arg.Year = year
	}
}
