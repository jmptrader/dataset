package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDataset(t *testing.T) {
	dataset, err := ReadDataset([]byte("\n1.23e2\n\t12.4    7"))
	if assert.NoError(t, err) {
		assert.Equal(t, NewDataset(1.23e2, 12.4, 7), dataset)
	}
}
