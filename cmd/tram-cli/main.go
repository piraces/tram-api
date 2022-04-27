package main

import (
	tramcli "github.com/piraces/tram-api/internal"
	"github.com/piraces/tram-api/internal/cli"
	"github.com/piraces/tram-api/internal/storage/api"
	"github.com/piraces/tram-api/internal/storage/csv"

	"github.com/spf13/cobra"
)

func main() {
	var apiRepo tramcli.TramStopRepo
	var csvRepo tramcli.CSVRepo
	apiRepo = api.NewApiRepository()
	csvRepo = csv.NewCsvRepo()

	rootCmd := &cobra.Command{Use: "tram-cli"}
	rootCmd.AddCommand(cli.InitTramStopsCmd(apiRepo, csvRepo))
	rootCmd.Execute()
}
