package storage

import (
	"github.com/EmmettCorp/delorean/pkg/domain"
	bolt "go.etcd.io/bbolt"
)

type VolumeRepo struct {
	db *bolt.DB
}

func NewVolumeRepo(db *bolt.DB) *VolumeRepo {
	return &VolumeRepo{
		db: db,
	}
}

func (r *VolumeRepo) Put(v domain.Volume) error {
	return nil
}

func (r *VolumeRepo) Get(id string) (domain.Volume, error) {
	v := domain.Volume{}
	return v, nil
}

func (r *VolumeRepo) List(activeOnly bool) []domain.Volume {
	return nil
}
