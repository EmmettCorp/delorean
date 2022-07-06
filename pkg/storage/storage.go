package storage

import (
	"path"

	bolt "go.etcd.io/bbolt"
)

const dbName = "delorean.db"

func New(ph string) (*bolt.DB, error) {
	return bolt.Open(path.Join(ph, dbName), 0600, nil)
}
