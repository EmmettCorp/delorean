package storage

import (
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

func (r *SnapshotRepo) Put(v domain.Snapshot) error {
	return nil
}

func (r *SnapshotRepo) Get(id string) (domain.Snapshot, error) {
	v := domain.Snapshot{}
	return v, nil
}

func (r *SnapshotRepo) List(activeOnly bool) []domain.Snapshot {
	return nil
}
