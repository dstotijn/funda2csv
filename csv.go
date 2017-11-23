package funda

import (
	"encoding/csv"
	"fmt"
	"io"
)

// ObjectsToCSV writes Funda objects in CSV format.
func ObjectsToCSV(objects Objects, w io.Writer) error {
	csvWriter := csv.NewWriter(w)
	csvWriter.Write([]string{"id", "address", "url"})
	for _, o := range objects {
		if err := csvWriter.Write([]string{
			o.ID,
			o.Address,
			o.URL.String(),
		}); err != nil {
			return fmt.Errorf("funda: could not write CSV record: %v", err)
		}

		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			return err
		}
	}

	return nil
}
