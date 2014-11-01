package statlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataset_Histogram(t *testing.T) {
	d := NewDataset(1, 2, 3, 4, 4, 4, 7, 8)
	assert.Equal(t, "  *     *         \n"+
		"  *     *     *   \n"+
		"  *     *     *   \n"+
		" 1.00  3.33  5.67 ",
		d.Histogram(20))
}
