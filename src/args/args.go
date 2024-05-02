package args

import (
	"errors"
	"github.com/anhgelus/license-generator/src/utils"
	"strings"
)

type Arguments struct {
	AppName     string
	LicenseType License
	Year        string
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
	gpl        License = generateBasicLicense("GPLv3")
	agpl       License = generateBasicLicense("AGPLv3")
	lgpl       License = generateBasicLicense("LGPLv3")
	mpl        License = generateBasicLicense("MPL")
	mit        License = generateBasicLicense("MIT")
	bsd        License = generateBasicLicense("BSD")
	freebsd    License = generateBasicLicense("FreeBSD")
	licenseMap         = make(map[string]License)
)

func generateBasicLicense(name string) License {
	return License{
		Name: name,
		File: "~",
	}
}

func GenerateLicenseMap() {
	licenseMap["gpl"] = gpl
	licenseMap["agpl"] = agpl
	licenseMap["lgpl"] = lgpl
	licenseMap["mpl"] = mpl
	licenseMap["mit"] = mit
	licenseMap["bsd"] = bsd
	licenseMap["freebsd"] = freebsd
}

func GetLicense(name string) (License, bool) {
	lic, found := licenseMap[strings.ToLower(name)]
	return lic, found
}

func AddLicense(license License, name string) {
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
	argLists     = [5]AvailableArgument{AppNameArg, LicenseArg, YearArg, AuthorsArg, ConfigPath}
	infoArgLists = [1]InfoArgument{LicenseListArg}
)

// GenerateParameter Generate the full parameter
func (arg AvailableArgument) GenerateParameter() string {
	return "--" + arg.Parameter
}

func (arg InfoArgument) GenerateParameter() string {
	return "-" + arg.Parameter
}

func (arg InfoArgument) GenerateText() string {
	return arg.textGenerator()
}

// AssignValueToArguments Assign the value of the argument passed through the cli inside the Arguments struct
func (arg *Arguments) assignValueToArguments(argument *AvailableArgument, v string) error {
	switch argument.Parameter {
	case "name":
		arg.AppName = v
	case "license":
		license, found := GetLicense(v)
		if !found {
			return errors.New("invalid license type, available license type: " + mapLicenseToString(licenseMap))
		}
		arg.LicenseType = license
	case "year":
		arg.Year = v
	case "authors":
		arg.Authors = parseAuthors(v)
	case "config-path":
		arg.ConfigPath = utils.RelativeToAbsolute(v, utils.ContextPath)
	default:
		return errors.New("unknown argument, use -h to see every arguments")
	}
	return nil
}

func mapLicenseToString(m map[string]License) string {
	str := ""
	for _, license := range m {
		str = str + ", " + license.Name
	}
	return str
}

func parseAuthors(s string) []string {
	return strings.Split(s, ",")
}

func helpText() string {
	str := ""
	for _, arg := range argLists {
		str += arg.GenerateParameter() + " " + arg.Argument + " - " + arg.Description + "\n"
	}
	for _, arg := range infoArgLists {
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
