package main

import (
	"fmt"
	"os"
	"regexp"

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
	cliInfo := utils.ShowMenuModel{
		Prompt: "Choose an option:",
		Choices: []string{
			utils.JoinAndConvertPathToOSFormat("docs", "tasks_docs"),
			utils.JoinAndConvertPathToOSFormat("docs", "app_docs"),
			"issues",
		},
	}
	docLocation := utils.ShowMenu(cliInfo, nil)
	docLocation = utils.JoinAndConvertPathToOSFormat(docLocation)
	targetName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"enter the name of the entity PLEASE USE DASHES OR UNDERLINE FOR SPACES"},
		},
	)
	pattern := `\s+`
	matched := regexp.MustCompile(pattern).MatchString(targetName)
	if matched == true {
		fmt.Printf("The document name cannot contain any speaces PLEASE USE DASHES OR UNDERLINE FOR SPACES !!!!!!!!!     :)")
		return
	}
	targetPath := utils.JoinAndConvertPathToOSFormat(docLocation, targetName)
	templatePath := utils.JoinAndConvertPathToOSFormat(docLocation, "template")
	utils.CopyDir(templatePath, targetPath)
	utils.RunCommand("code", []string{targetPath})
}
