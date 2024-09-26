package storage

import "os"

func (s *DomainStorage) Get(filepath string) (*os.File, error) {
	file, errOpen := os.Open(filepath)

	return file, errOpen
}
