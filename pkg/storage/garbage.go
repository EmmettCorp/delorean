package storage

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

const garbageBucketName = "garbage"

// GarbageRepo keeps paths to old subvolumes after restore operation.
// Usually path looks like path.Join(domain.DeloreanMountPoint, fmt.Sprintf("%s.old", vol.Subvol)).
// We can remove old subvolume after reboot only.
// The solution is to keep path in storage and remove on the next application run after reboot.
// This logic works only with force reboot.
type GarbageRepo struct {
	db     *bolt.DB
	bucket []byte
}

// NewGarbageRepo creates a new snapshot repository.
func NewGarbageRepo(db *bolt.DB) (*GarbageRepo, error) {
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

	return &GarbageRepo{
		db:     db,
		bucket: []byte(garbageBucketName),
	}, nil
}

// Put saves a path for old file system to data storage.
func (r *GarbageRepo) Put(ph string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)

		return b.Put([]byte(ph), nil)
	})
}

// List returns the list paths to old file systems.
func (r *GarbageRepo) List() ([]string, error) {
	paths := []string{}
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)

		return b.ForEach(func(k, _ []byte) error {
			ph := string(k)

			paths = append(paths, ph)

			return nil
		})
	})

	return paths, err
}

// Delete path from database.
func (r *GarbageRepo) Delete(ph string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(r.bucket)

		return b.Delete([]byte(ph))
	})
}
