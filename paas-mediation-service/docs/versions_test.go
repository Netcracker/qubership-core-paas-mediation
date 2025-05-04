package docs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVersion(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(2, MajorVersion)
	assertions.Equal(0, MinorVersion)
	assertions.Equal([]int{2}, SupportedMajors)
}
