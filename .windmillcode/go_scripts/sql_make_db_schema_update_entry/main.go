package main

import (
	"fmt"
	"os"
	"time"

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
	var defaultDatabaseOptions = []string{"postgres", "mysql", "mssql", "oraclesql", "mariadb", "firestore"}
	var databaseOptions []string
	if settings.ExtensionPack.DatabaseOptions != nil && len(settings.ExtensionPack.DatabaseOptions) > 0 {
		databaseOptions = settings.ExtensionPack.DatabaseOptions
	} else {
		databaseOptions = defaultDatabaseOptions
	}
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Enter the database script location (refer to the folder in apps\\database for your project)",
		Choices: databaseOptions,
		Default: "mysql",
	}
	databaseToBackup := utils.ShowMenu(cliInfo, nil)
	cliInfo = utils.ShowMenuModel{
		Prompt:  "open the update.sql file",
		Choices: []string{"YES", "NO"},
	}
	openUpdateFile := utils.ShowMenu(cliInfo, nil)
	databaseBackupLocation := utils.JoinAndConvertPathToOSFormat("apps", "database", databaseToBackup, "schema_entries")

	utils.CDToLocation(workspaceRoot)
	utils.CDToLocation(databaseBackupLocation)
	currentDay := time.Now().Format("06-1-02_15-04-05")
	err = os.MkdirAll(currentDay, 0755)
	if err != nil {
		fmt.Printf("unable to make %s : \n Err msg: %s", currentDay, err.Error())
	}
	utils.CopyDir(
		"template", currentDay,
	)
	if openUpdateFile == "YES" {
		utils.RunCommand("code", []string{utils.JoinAndConvertPathToOSFormat(currentDay, "update.sql")})
	}

}
