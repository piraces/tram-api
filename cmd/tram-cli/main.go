package main

import (
	tramcli "github.com/piraces/tram-api/internal"
	"github.com/piraces/tram-api/internal/cli"
	"github.com/piraces/tram-api/internal/storage/api"

	"github.com/spf13/cobra"
)

func main() {
	var repo tramcli.TramStopRepo
	repo = api.NewApiRepository()

	rootCmd := &cobra.Command{Use: "tram-cli"}
	rootCmd.AddCommand(cli.InitTramStopsCmd(repo))
	rootCmd.Execute()
}
