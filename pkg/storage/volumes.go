package storage

import (
	"encoding/json"

	"github.com/EmmettCorp/delorean/pkg/domain"
	bolt "go.etcd.io/bbolt"
)

// VolumeRepo is repository for subvolumes.
type VolumeRepo struct {
	db     *bolt.DB
	bucket []byte
}

// NewVolumeRepo creates a new VolumeRepo.
func NewVolumeRepo(db *bolt.DB) *VolumeRepo {
	return &VolumeRepo{
		db:     db,
		bucket: []byte("volumes"),
	}
}

// Put creates a new volume object in storage or updates an
// existing one.
func (r *VolumeRepo) Put(v domain.Volume) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		dt, err := json.Marshal(v)
		if err != nil {
			return err
		}

		return b.Put([]byte(v.ID), dt)
	})
}

// Get returns volume by id if exists or error.
func (r *VolumeRepo) Get(id string) (domain.Volume, error) {
	v := domain.Volume{}

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		v := b.Get([]byte(id))
		err := json.Unmarshal(v, &v)
		if err != nil {
			return err
		}

		return nil
	})

	return v, err
}

// List returns the list of volumes in database.
func (r *VolumeRepo) List(activeOnly bool) ([]domain.Volume, error) {
	var volumes []domain.Volume
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)

		return b.ForEach(func(k, v []byte) error {
			vol := domain.Volume{}
			err := json.Unmarshal(v, &vol)
			if err != nil {
				return err
			}
			if activeOnly && !vol.Active {
				return nil
			}
			volumes = append(volumes, vol)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return volumes, nil
}
