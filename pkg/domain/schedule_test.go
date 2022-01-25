package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncrease(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	sched := Schedule{}

	t.Run("monthly", func(t *testing.T) {
		t.Parallel()

		updated := sched.Increase(0)
		rq.True(updated)
		sched.Monthly = maxSnapshotAmount

		updated = sched.Increase(0)
		rq.False(updated)
	})

	t.Run("weekly", func(t *testing.T) {
		t.Parallel()

		updated := sched.Increase(1)
		rq.True(updated)
		sched.Weekly = maxSnapshotAmount

		updated = sched.Increase(1)
		rq.False(updated)
	})

	t.Run("daily", func(t *testing.T) {
		t.Parallel()

		updated := sched.Increase(2)
		rq.True(updated)
		sched.Daily = maxSnapshotAmount

		updated = sched.Increase(2)
		rq.False(updated)
	})

	t.Run("hourly", func(t *testing.T) {
		t.Parallel()

		updated := sched.Increase(3)
		rq.True(updated)
		sched.Hourly = maxSnapshotAmount

		updated = sched.Increase(3)
		rq.False(updated)
	})

	t.Run("boot", func(t *testing.T) {
		t.Parallel()

		updated := sched.Increase(4)
		rq.True(updated)
		sched.Boot = maxSnapshotAmount

		updated = sched.Increase(4)
		rq.False(updated)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		updated := sched.Increase(5)
		rq.False(updated)
	})
}

func TestDecrease(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	sched := Schedule{}

	t.Run("monthly", func(t *testing.T) {
		t.Parallel()

		updated := sched.Decrease(0)
		rq.False(updated)
		sched.Monthly = maxSnapshotAmount

		updated = sched.Decrease(0)
		rq.True(updated)
	})

	t.Run("weekly", func(t *testing.T) {
		t.Parallel()

		updated := sched.Decrease(1)
		rq.False(updated)
		sched.Weekly = maxSnapshotAmount

		updated = sched.Decrease(1)
		rq.True(updated)
	})

	t.Run("daily", func(t *testing.T) {
		t.Parallel()

		updated := sched.Decrease(2)
		rq.False(updated)
		sched.Daily = maxSnapshotAmount

		updated = sched.Decrease(2)
		rq.True(updated)
	})

	t.Run("hourly", func(t *testing.T) {
		t.Parallel()

		updated := sched.Decrease(3)
		rq.False(updated)
		sched.Hourly = maxSnapshotAmount

		updated = sched.Decrease(3)
		rq.True(updated)
	})

	t.Run("boot", func(t *testing.T) {
		t.Parallel()

		updated := sched.Decrease(4)
		rq.False(updated)
		sched.Boot = maxSnapshotAmount

		updated = sched.Decrease(4)
		rq.True(updated)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		updated := sched.Decrease(5)
		rq.False(updated)
	})
}
