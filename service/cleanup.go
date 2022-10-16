package service

import (
	"github.com/abdukhashimov/golang-hex-architecture/storage"
)

type CleanupI interface {
	Cleanup() error
}

type cleanup struct {
	db storage.StorageI
}

func NewCleanupService(db storage.StorageI) CleanupI {
	return &cleanup{
		db: db,
	}
}

func (c *cleanup) Cleanup() error {
	err := c.db.Cleanup().Cleanup()

	return err
}
