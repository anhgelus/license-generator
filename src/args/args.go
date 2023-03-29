package args

import (
	"errors"
	"strings"
)

type Arguments struct {
	AppName     string
	LicenseType License
	Year        string
	Authors     []string
	Question    bool
	Info        bool
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
	HelpArg = InfoArgument{
		Parameter:     "h",
		textGenerator: helpText,
		Description:   "Show the help",
	}
	argLists     = [4]AvailableArgument{AppNameArg, LicenseArg, YearArg, AuthorsArg}
	infoArgLists = [0]InfoArgument{}
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
		str = str + arg.GenerateParameter() + " " + arg.Argument + " - " + arg.Description + "\n"
	}
	for _, arg := range infoArgLists {
		str = str + arg.GenerateParameter() + " - " + arg.Description + "\n"
	}
	str = str + "-h - Show the help"
	return str
}
