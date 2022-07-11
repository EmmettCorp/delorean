package storage

import (
	"encoding/json"

	"github.com/EmmettCorp/delorean/pkg/domain"
	bolt "go.etcd.io/bbolt"
)

type SnapshotRepo struct {
	db     *bolt.DB
	bucket []byte
}

func NewSnapshotRepo(db *bolt.DB) *SnapshotRepo {
	return &SnapshotRepo{
		db:     db,
		bucket: []byte("snapshots"),
	}
}

func (r *SnapshotRepo) Put(sn domain.Snapshot) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		dt, err := json.Marshal(sn)
		if err != nil {
			return err
		}

		return b.Put([]byte(sn.ID), dt)
	})
}

func (r *SnapshotRepo) Get(id string) (domain.Snapshot, error) {
	v := domain.Snapshot{}
	return v, nil
}

func (r *SnapshotRepo) List(activeOnly bool) []domain.Snapshot {
	return nil
}
