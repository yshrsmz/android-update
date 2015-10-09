package main

import (
	"github.com/codegangsta/cli"
)

// Flags for command
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Value: "./android-update.yml",
		Usage: "set config file.",
	},
	cli.StringFlag{
		Name:   "sdk",
		Value:  "",
		Usage:  "set Android SDK directory.",
		EnvVar: "ANDROID_HOME",
	},
}
