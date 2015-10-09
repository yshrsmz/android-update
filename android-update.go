package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var _sdkPath string

func main() {
	app := cli.NewApp()

	// configure app
	app.Name = "android-update"
	app.Version = Version
	app.Usage = "Update Android SDKs & Plugins"
	app.Author = "yshrsmz"
	app.Email = "the.phantom.bane@gmail.com"
	app.Flags = Flags
	app.Action = doAction

	app.Run(os.Args)
}
