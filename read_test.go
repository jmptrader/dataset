package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDataset(t *testing.T) {
	dataset, err := ReadDataset([]byte("\n12.3\n\t12.4    7"))
	if assert.NoError(t, err) {
		assert.Equal(t, &Dataset{
			records: []Record{12.3, 12.4, 7},
		}, dataset)
	}
}
