package storage

import (
	"encoding/csv"
	"go-project/internal/models"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var testDomains = []*models.Domain{
	{ID: 1, Name: "example.com", ExpireAt: "2023-12-31"},
	{ID: 2, Name: "test.com", ExpireAt: "2024-01-31"},
}

func TestSaveDomains(t *testing.T) {

	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "api.csv")

	domainStorage := NewStorageDomain()
	err := domainStorage.Save(testDomains, filePath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatalf("error closing file: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	expectedRecords := [][]string{
		{strconv.Itoa(testDomains[0].ID), testDomains[0].Name, testDomains[0].ExpireAt, time.DateTime},
		{strconv.Itoa(testDomains[1].ID), testDomains[1].Name, testDomains[1].ExpireAt, time.DateTime},
	}

	if !reflect.DeepEqual(records, expectedRecords) {
		t.Errorf("expected %v, got %v", expectedRecords, records)
	}
}
