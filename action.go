package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/codegangsta/cli"
	"gopkg.in/pipe.v2"
	"gopkg.in/yaml.v2"
)

// Config struct for android-update.yml
type Config struct {
	Sdk      string
	Packages []string
}

// main action
func doAction(c *cli.Context) {
	fmt.Printf("use %v as config file\n", c.String("config"))

	configPath := c.String("config")
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	_sdkPath = c.String("sdk")
	if config.Sdk != "" {
		_sdkPath = config.Sdk
	}

	fmt.Printf("sdk path is %v\n", _sdkPath)
	printTargets(config.Packages)

	for i, len := 0, len(config.Packages); i < len; i++ {
		execInstall2(config.Packages[i])
	}
}

// read config file from provided path
func readConfig(path string) (Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	output := Config{}
	if err != nil {
		return output, err
	}

	err = yaml.Unmarshal(yamlFile, &output)
	if err != nil {
		return output, err
	}

	return output, nil
}

// print target packages to install
func printTargets(packages []string) {
	fmt.Printf("----- Install following %v packages -----\n", len(packages))
	for i, len := 0, len(packages); i < len; i++ {
		fmt.Println(packages[i])
	}
	fmt.Println("-----------------------------------------")
}

// execute actual install command
func execInstall(target string) {
	sayYes := exec.Command("echo y")
	doInstall := exec.Command("android", "update", "sdk", "-a", "-u", "-t", target)

	sayYesOut, _ := sayYes.StdoutPipe()
	sayYes.Start()
	doInstall.Stdin = sayYesOut
	out, _ := doInstall.CombinedOutput()

	fmt.Println(string(out))
}

func execInstall2(target string) {
	androidExec := "android"
	if _sdkPath != "" {
		androidExec = _sdkPath + "/tools/" + androidExec
	}
	androidExec, _ = filepath.Abs(androidExec)
	fmt.Printf("command path: %v", androidExec)

	p := pipe.Line(
		pipe.Exec("echo", "y"),
		pipe.Exec(androidExec, "update", "sdk", "-a", "-u", "-t", target),
	)

	output, err := pipe.CombinedOutput(p)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println(string(output))
}
