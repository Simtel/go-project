package domains

import (
	"encoding/csv"
	"fmt"
	"go-project/models"
	"os"
	"strconv"
)

func SaveDomains(domains []*models.Domain) error {
	file, err := os.Create("var/domains.csv")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var records [][]string

	for _, domain := range domains {
		records = append(
			records,
			[]string{strconv.Itoa(domain.ID), domain.Name, domain.ExpireAt},
		)
	}

	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("error writing record to file: %v", err)
		}
	}
	return nil
}
