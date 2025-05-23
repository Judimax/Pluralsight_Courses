package main

import (
	"fmt"
	"os"
	"strings"

	"main/shared"

	"github.com/iancoleman/strcase"
	"github.com/windmillcode/go_cli_scripts/v6/utils"
)

func main() {

	scriptLocation, err := os.Getwd()
	if err != nil {
		return
	}
	templateLocation := utils.JoinAndConvertPathToOSFormat(scriptLocation, "template")
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
	utils.CDToFlutterApp()
	flutterApp, err := os.Getwd()
	if err != nil {
		return
	}
	pageName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The name you would like to give to the page"},
			ErrMsg: "You must provide a value",
		},
	)
	entityName := pageName
	snakeCasePageName := strcase.ToSnake(pageName)
	providerLocation := utils.JoinAndConvertPathToOSFormat(flutterApp, "lib", "pages", snakeCasePageName)
	newTemplatePath := utils.JoinAndConvertPathToOSFormat(providerLocation, fmt.Sprintf("%s_page.dart", snakeCasePageName))
	newRiverPodProviderPath := utils.JoinAndConvertPathToOSFormat(providerLocation, fmt.Sprintf("%s_page_riverpod_provider.dart", snakeCasePageName))
	utils.CopyDir(templateLocation, providerLocation)
	os.Rename(
		utils.JoinAndConvertPathToOSFormat(providerLocation, "template_page.dart"),
		newTemplatePath,
	)
	os.Rename(
		utils.JoinAndConvertPathToOSFormat(providerLocation, "template_riverpod_provider.dart"),
		newRiverPodProviderPath,
	)

	for _, path := range []string{newTemplatePath, newRiverPodProviderPath} {
		fileString, err := utils.ReadFile(path)
		if err != nil {
			return
		}
		fileString = strings.ReplaceAll(fileString, "WMLTemplate", strcase.ToCamel(entityName))
		fileString = strings.ReplaceAll(fileString, "Wml", "WML")
		fileString = strings.ReplaceAll(fileString, "template", snakeCasePageName)
		utils.OverwriteFile(path, fileString)
	}

}
