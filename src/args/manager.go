package args

import (
	"os"
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
			arguments.AssignValueToArguments(&av, args[i+1])
			l++
		}
	}
	if l != 4 {
		arguments.Question = true
	}
	return &arguments
}
