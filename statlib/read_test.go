package statlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDataset(t *testing.T) {
	dataset, err := ReadDataset([]byte("\n1.23e2\n\t12.4    7"))
	if assert.NoError(t, err) {
		assert.Equal(t, NewDataset(1.23e2, 12.4, 7), dataset)
	}

	dataset, err = ReadDataset([]byte("test1.23e2\ntest\t12.4  1 hi 7"))
	if assert.NoError(t, err) {
		assert.Equal(t, NewDataset(12.4, 1, 7), dataset)
	}

	dataset, err = ReadDataset([]byte("12.4,1;7"))
	if assert.NoError(t, err) {
		assert.Equal(t, NewDataset(12.4, 1, 7), dataset)
	}
}
