package source_test

import (
	"clover-data-processor/app/source"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileDataSource(t *testing.T) {
	dataSource := source.FileDataSource{}

	t.Run("GetSpecPathSuccess", func(t *testing.T) {
		paths := dataSource.GetSpecPath("../../test/specs")

		assert.Equal(t, 2, len(paths))
		for _, p := range paths {
			assert.True(t, strings.HasSuffix(p, ".csv"))
		}

	})

	t.Run("GetDataPathSuccess", func(t *testing.T) {
		paths := dataSource.GetDataPath("../../test/data", "testformat1")

		assert.Equal(t, 1, len(paths))
		for _, p := range paths {
			assert.True(t, strings.HasSuffix(p, ".txt"))
		}
	})
}
