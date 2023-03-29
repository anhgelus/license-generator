package utils

import (
	"github.com/BurntSushi/toml"
	"os"
)

func FileContent(path string, file os.DirEntry) []byte {
	content, err := os.ReadFile(path + "/" + file.Name())
	HandleError(err)
	return content
}

func DecodeToml(content []byte, t any) {
	_, err := toml.Decode(string(content), t)
	HandleError(err)
}

// RelativeToAbsolute convert relative path into absolute path
//
// path: the path to convert
//
// contextPath: the contextual path, with a slash (/) at the end
func RelativeToAbsolute(path string, contextPath string) string {
	println(path, contextPath)
	switch string(path[0]) {
	case "/":
		return path
	case "~":
		return path
	case ".":
		return contextPath + path[2:]
	default:
		return contextPath + path
	}
}
