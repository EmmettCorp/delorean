/*
Package storage keeps logic for volumes and snapshots info repositories.
*/
package storage

import (
	"path"

	"github.com/EmmettCorp/delorean/pkg/domain"
	bolt "go.etcd.io/bbolt"
)

const dbName = "delorean.db"

// New creates a new connection to bolt database.
func New(ph string) (*bolt.DB, error) {
	return bolt.Open(path.Join(ph, dbName), domain.RWFileMode, nil)
}
