package storage

import (
	"encoding/json"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/domain"
	bolt "go.etcd.io/bbolt"
)

const snapBucketName = "snapshots"

type SnapshotRepo struct {
	db     *bolt.DB
	bucket []byte
}

func NewSnapshotRepo(db *bolt.DB) (*SnapshotRepo, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(snapBucketName))
		if err != nil {
			return fmt.Errorf("could not create bucket `%s`: %v", snapBucketName, err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &SnapshotRepo{
		db:     db,
		bucket: []byte(snapBucketName),
	}, nil
}

func (r *SnapshotRepo) Put(sn domain.Snapshot) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		dt, err := json.Marshal(sn)
		if err != nil {
			return err
		}

		return b.Put([]byte(sn.Path), dt)
	})
}

func (r *SnapshotRepo) List(vIDs []string) ([]domain.Snapshot, error) {
	snaps := []domain.Snapshot{}
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)
		return b.ForEach(func(_, v []byte) error {
			sn := domain.Snapshot{}
			err := json.Unmarshal(v, &sn)
			if err != nil {
				return err
			}

			if !inVolumes(vIDs, sn.VolumeID) {
				return nil
			}

			snaps = append(snaps, sn)

			return nil
		})
	})

	return snaps, err
}

func (r *SnapshotRepo) Delete(ph string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)

		return b.Delete([]byte(ph))
	})
}

func inVolumes(vIDs []string, v string) bool {
	for i := range vIDs {
		if vIDs[i] == v {
			return true
		}
	}

	return false
}
