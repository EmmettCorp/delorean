package findmnt

import (
	"testing"

	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/stretchr/testify/require"
)

func TestGetSubvol(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		res := getSubvol("/")
		rq.Equal(domain.Subvol5, res)

		res = getSubvol("/@")
		rq.Equal("@", res)

		res = getSubvol("/root")
		rq.Equal("root", res)
	})
}
