package main

import "license-generator/src/args"

func main() {
	args.GenerateLicenseMap()
	arg := args.ParseCliArgs()
	arg.HandleArgs()
}
