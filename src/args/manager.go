package args

import (
	"license-generator/src/utils"
	"os"
	"strings"
)

func ParseCliArgs() *Arguments {
	args := os.Args
	// if the user don't give basic information through args
	if len(args) < 2 {
		return &Arguments{Question: true}
	}
	arguments := Arguments{}
	l := 0
	for i, arg := range args {
		if i == 0 {
			continue
		}
		for _, av := range argLists {
			if arg != av.GenerateParameter() {
				continue
			}
			arguments.assignValueToArguments(&av, args[i+1])
			l++
		}
	}
	if l != 4 {
		arguments.Question = true
	}
	return &arguments
}

func (arg *Arguments) HandleArgs() {
	if arg.Question {
		arg.handleQuestion()
	}
	m := make(map[string]string)
	m["App Name"] = arg.AppName
	m["License"] = string(arg.LicenseType)
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
	if arg.LicenseType == "" {
		oldLicense := ""
		print("License: ")
		err := utils.Scan(&oldLicense)
		utils.HandleError(err)
		license := GetLicense(oldLicense)
		if license == "" {
			println("Unknown license type. Aborted.")
			os.Exit(2)
		}
		println("The license is: " + license)
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
