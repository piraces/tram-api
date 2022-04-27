package csv

import (
	"encoding/csv"
	"os"

	tramcli "github.com/piraces/tram-api/internal"
)

type csvRepo struct {
	separator rune
}

func NewCsvRepo() tramcli.CSVRepo {
	return &csvRepo{','}
}

var headers = []string{"TramStopId", "Title", "Icon", "Description", "LastUpdated", "URI"}

func (c *csvRepo) SaveCSV(tramStops *[]tramcli.TramStop, filename string) error {
	csvFile, err := os.Create(filename)

	if err != nil {
		csvFile.Close()
		return err
	}

	csvWriter := csv.NewWriter(csvFile)

	if err = csvWriter.Write(headers); err != nil {
		csvWriter.Flush()
		csvFile.Close()
		return err
	}

	for _, tramStop := range *tramStops {
		if err = csvWriter.Write([]string{tramStop.TramStopId, tramStop.Title, tramStop.Icon, tramStop.Description,
			tramStop.LastUpdated, tramStop.URI}); err != nil {
			csvWriter.Flush()
			csvFile.Close()
			return err
		}
	}

	csvWriter.Flush()
	csvFile.Close()
	return nil
}
