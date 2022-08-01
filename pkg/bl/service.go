/*
Package bl acts as a connector between database and application and keeps all the business logic.
*/
package bl

import (
	"fmt"
	"os"
	"path"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/domain"
)

type (
	snapshotRepo interface {
		Put(sn domain.Snapshot) error
		List(vUIDs []string) ([]domain.Snapshot, error)
		Delete(ph string) error
	}

	garbageRepo interface {
		Put(ph string) error
		List() ([]string, error)
		Delete(ph string) error
	}

	Service struct {
		snapshotRepo snapshotRepo
		garbageRepo  garbageRepo
		config       *config.Config
	}
)

func New(sr snapshotRepo, gr garbageRepo, cfg *config.Config) *Service {
	return &Service{
		snapshotRepo: sr,
		garbageRepo:  gr,
		config:       cfg,
	}
}

func (s *Service) DeleteSnapshot(snap domain.Snapshot) error {
	err := btrfs.DeleteSnapshot(snap.Path)
	if err != nil {
		return fmt.Errorf("can't delete from btrfs: %v", err)
	}

	err = s.snapshotRepo.Delete(snap.Path)
	if err != nil {
		return fmt.Errorf("can't delete info from storage: %v", err)
	}

	return nil
}

func (s *Service) Restore(snap domain.Snapshot) error {
	if !s.config.VolumeInRootFs(snap.Volume) {
		return fmt.Errorf("volume %s is not a child subvolume top level subvolume", snap.Volume.Label)
	}

	err := s.CreateSnapshotForVolume(snap.Volume, snap.Type)
	if err != nil {
		return err
	}

	subvolumeDelorianMountPoint := path.Join(domain.DeloreanMountPoint, snap.Volume.Subvol)
	oldFsDelorianMountPoint := path.Join(domain.DeloreanMountPoint, fmt.Sprintf("%s.old", snap.Volume.Subvol))

	err = os.Rename(subvolumeDelorianMountPoint, oldFsDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't rename directory %s: %v", oldFsDelorianMountPoint, err)
	}

	err = btrfs.Restore(snap.Path, subvolumeDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't create snapshot for %s: %v", snap.Volume.Device.MountPoint, err)
	}

	err = s.garbageRepo.Put(oldFsDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't put old filesystem path to garbage storage %s: %v", oldFsDelorianMountPoint, err)
	}

	// force reboot logic

	return nil
}

func (s *Service) CreateSnapshotForVolume(vol domain.Volume, snapType domain.SnapshotType) error {
	snap := domain.NewSnapshot(vol.SnapshotsPath, s.config.KernelVersion, vol, snapType)

	err := btrfs.CreateSnapshot(vol.Device.MountPoint, snap)
	if err != nil {
		return fmt.Errorf("can't create snapshot for %s: %v", snap.Path, err)
	}

	err = s.snapshotRepo.Put(snap)
	if err != nil {
		return fmt.Errorf("can't put snapshot to storage with path %s: %v", snap.Path, err)
	}

	return nil
}

func (s *Service) SnapshotsList(vUIDs []string) ([]domain.Snapshot, error) {
	return s.snapshotRepo.List(vUIDs)
}
