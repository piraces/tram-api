package cli

import (
	tramcli "github.com/piraces/tram-api/internal"
	"github.com/spf13/cobra"
	"log"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

func InitTramStopsCmd(apiRepository tramcli.TramStopRepo, csvRepository tramcli.CSVRepo) *cobra.Command {
	tramStopsCmd := &cobra.Command{
		Use:   "tram-stops",
		Short: "Saves current tram stops status to a CSV file",
		Run:   runTramStopsFn(apiRepository, csvRepository),
	}

	tramStopsCmd.Flags().StringP("output", "o", "output.csv", "Filename of output file")

	return tramStopsCmd
}

func runTramStopsFn(apiRepository tramcli.TramStopRepo, csvRepository tramcli.CSVRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatalf("Error while retrieving CLI flag: %s", err)
		}

		tramStops, err := apiRepository.GetTramStops()
		if err != nil {
			log.Fatalf("Error while retrieving tram stops: %s", err)
		}

		err = csvRepository.SaveCSV(&tramStops, output)
		if err != nil {
			log.Fatalf("Error while saving output to csv: %s", err)
		}
	}
}
