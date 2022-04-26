package cli

import (
	"fmt"
	tramcli "github.com/piraces/tram-api/internal"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

func InitTramStopsCmd(repository tramcli.TramStopRepo) *cobra.Command {
	tramStopsCmd := &cobra.Command{
		Use:   "tram-stops",
		Short: "Saves current tram stops status to a CSV file",
		Run:   runTramStopsFn(repository),
	}

	return tramStopsCmd
}

func runTramStopsFn(repository tramcli.TramStopRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		tramStops, err := repository.GetTramStops()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(tramStops)
		}
	}
}
