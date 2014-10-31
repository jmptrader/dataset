package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataset_Percentile(t *testing.T) {
	d := NewDataset(12, 13, 14)
	assert.Equal(t, 12, d.Percentile(0))
	assert.Equal(t, 12, d.Percentile(.25))
	assert.Equal(t, 13, d.Percentile(.5))
	assert.Equal(t, 13, d.Percentile(.75))
	assert.Equal(t, 14, d.Percentile(1))

	assert.Panics(t, func() {
		d.Percentile(1.2)
	})
	assert.Panics(t, func() {
		d.Percentile(-0.1)
	})

	d = NewDataset(12, 13, 14, 15)
	assert.Equal(t, 12, d.Percentile(0))
	assert.Equal(t, 12.5, d.Percentile(.25))
	assert.Equal(t, 13.5, d.Percentile(.5))
	assert.Equal(t, 14.5, d.Percentile(.75))
	assert.Equal(t, 15, d.Percentile(1))
}
