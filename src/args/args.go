package args

import (
	"strings"
)

type Arguments struct {
	AppName     string
	LicenseType *License
	Years       string
	Authors     []string
	ConfigPath  string
	Question    bool
	Info        bool
	InfoText    func() string
}

type License struct {
	Name string
	File string
}

var (
	gpl        = generateBasicLicense("GPLv3")
	agpl       = generateBasicLicense("AGPLv3")
	lgpl       = generateBasicLicense("LGPLv3")
	mpl        = generateBasicLicense("MPL")
	mit        = generateBasicLicense("MIT")
	bsd        = generateBasicLicense("BSD")
	freebsd    = generateBasicLicense("FreeBSD")
	licenseMap = make(map[string]*License)
)

func init() {
	licenseMap["gpl"] = &gpl
	licenseMap["agpl"] = &agpl
	licenseMap["lgpl"] = &lgpl
	licenseMap["mpl"] = &mpl
	licenseMap["mit"] = &mit
	licenseMap["bsd"] = &bsd
	licenseMap["freebsd"] = &freebsd
}

func generateBasicLicense(name string) License {
	return License{
		Name: name,
		File: "~",
	}
}

func GetLicense(name string) (*License, bool) {
	lic, found := licenseMap[strings.ToLower(name)]
	return lic, found
}

func AddLicense(license *License, name string) {
	licenseMap[name] = license
}

type AvailableArgument struct {
	Parameter   string
	Description string
	Argument    string
}

type InfoArgument struct {
	Parameter     string
	textGenerator func() string
	Description   string
}

type OtherArgument struct {
	Parameter   string
	Description string
}

var (
	AppNameArg = AvailableArgument{
		Parameter:   "name",
		Description: "Set the name of the project",
		Argument:    "<string>",
	}
	LicenseArg = AvailableArgument{
		Parameter:   "license",
		Description: "Set the license of the project",
		Argument:    "<string>",
	}
	YearArg = AvailableArgument{
		Parameter:   "year",
		Description: "Set the year",
		Argument:    "<int|string>",
	}
	AuthorsArg = AvailableArgument{
		Parameter:   "authors",
		Description: "Set the authors of the project, separate them with the coma (,)",
		Argument:    "[string,]",
	}
	ConfigPath = AvailableArgument{
		Parameter:   "config-path",
		Description: "Set a path to the config for using custom licenses",
		Argument:    "string",
	}
	HelpArg = InfoArgument{
		Parameter:     "h",
		textGenerator: helpText,
		Description:   "Show the help",
	}
	LicenseListArg = InfoArgument{
		Parameter:     "l",
		textGenerator: listLicense,
		Description:   "List every available license",
	}
	VerboseArg = OtherArgument{
		Parameter:   "v",
		Description: "Verbose mode",
	}
	argLists      = [5]AvailableArgument{AppNameArg, LicenseArg, YearArg, AuthorsArg, ConfigPath}
	infoArgLists  = [1]InfoArgument{LicenseListArg}
	otherArgLists = [1]OtherArgument{VerboseArg}
)

// GenerateParameter Generate the full parameter
func (arg *AvailableArgument) GenerateParameter() string {
	return "--" + arg.Parameter
}

func (arg *InfoArgument) GenerateParameter() string {
	return "-" + arg.Parameter
}

func (arg *OtherArgument) GenerateParameter() string {
	return "-" + arg.Parameter
}

func (arg *InfoArgument) GenerateText() string {
	return arg.textGenerator()
}

func parseAuthors(s string) []string {
	var authors []string
	for _, a := range strings.Split(s, ",") {
		authors = append(authors, strings.Trim(a, " "))
	}
	return authors
}

func helpText() string {
	str := ""
	for _, arg := range argLists {
		str += arg.GenerateParameter() + " " + arg.Argument + " - " + arg.Description + "\n"
	}
	for _, arg := range infoArgLists {
		str += arg.GenerateParameter() + " - " + arg.Description + "\n"
	}
	for _, arg := range otherArgLists {
		str += arg.GenerateParameter() + " - " + arg.Description + "\n"
	}
	str = str + "-h - Show the help"
	return str
}

func listLicense() string {
	str := ""
	for id, license := range licenseMap {
		str += "- " + license.Name + " (" + id + ")\n"
	}
	return str[:len(str)-1]
}
