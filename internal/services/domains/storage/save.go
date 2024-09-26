package storage

import (
	"encoding/csv"
	"fmt"
	"go-project/internal/models"
	"os"
	"strconv"
	"time"
)

func (s *DomainStorage) Save(domains []*models.Domain, filePath string) error {
	file, err := os.Create(filePath)
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
			[]string{strconv.Itoa(domain.GetId()), domain.Name, domain.ExpireAt, time.DateTime},
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
