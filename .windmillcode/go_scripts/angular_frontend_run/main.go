package main

import (
	"fmt"
	"os"
	"strings"
	"main/shared"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	shared.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	utils.SetGlobalVars(
		utils.SetGlobalVarsOptions{
			NonInteractive :settings.ExtensionPack.ProcessIfDefaultIsPresent,
		},
	)
	angularFrontend := settings.ExtensionPack.AngularFrontend
	var appPort string
	settingsAppPort := utils.IntToStr(settings.ExtensionPack.Ports.AngularRun0)
	if settingsAppPort != "0" {
		appPort = settingsAppPort
	} else {
		appPort = utils.IntToStr(4200)
	}

	appPort = utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The port to use"},
			Default: appPort,
		},
	)

	cliInfo := utils.ShowMenuModel{
		Prompt:  "run with cache?",
		Choices: []string{ "FALSE","TRUE"},
		Default: "FALSE",
	}
	runWithCache := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "run concurently with the scss server",
		Choices: []string{"TRUE", "FALSE"},
		Default: "TRUE",
	}
	concurrentWithScss := utils.ShowMenu(cliInfo, nil)

	shared.SetNodeJSEnvironment(settings.ExtensionPack.NodeJSVersion0)

	defaultConfig := "development"
	if utils.IsRunningInDocker() {
		defaultConfig = strings.Replace(defaultConfig, "development", "docker-development", 1)
		for key, val := range angularFrontend.Configurations {
			angularFrontend.Configurations[key] = strings.Replace(val, "development", "docker-development", 1)
		}
	}

	cliInfo = utils.ShowMenuModel{
		Prompt:  "the configuration to run",
		Choices: angularFrontend.Configurations,
		Default: defaultConfig,
		Other:   true,
	}
	serveConfiguration := utils.ShowMenu(cliInfo, nil)


	utils.CDToAngularApp()
	if runWithCache == "FALSE" {
		if err := os.RemoveAll(utils.JoinAndConvertPathToOSFormat(".", ".angular")); err != nil {
			fmt.Println("Error:", err)
			return
		}
		if err := os.RemoveAll(utils.JoinAndConvertPathToOSFormat(".", ".nx")); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	if concurrentWithScss == "FALSE" {
		utils.RunCommand("npx", []string{
			"ng", "serve", "-c",
			serveConfiguration, "--ssl=true", fmt.Sprintf("--ssl-key=%s", os.Getenv("WML_CERT_KEY0")),
			fmt.Sprintf("--ssl-cert=%s", os.Getenv("WML_CERT0")) + " --port " + appPort,
			"sass", "--watch", "src/assets/styles/turn_to_css/entry.scss:src/assets/styles/compiled.css",
		})
	} else {
		utils.RunCommand("npm", []string{"run", "build:scss"})
		utils.RunCommand("npx", []string{
			"concurrently",
			"ng serve -c " + serveConfiguration + " --ssl=true --ssl-key=" + os.Getenv("WML_CERT_KEY0") + " --ssl-cert=" + os.Getenv("WML_CERT0") + " --port " + appPort,
			fmt.Sprintf("chokidar \"%s\" -c \"%s\"",
				"src/assets/styles/turn_to_css/**/*.{scss,css}",
				"npm run build:scss",
			),
		})
	}

}
