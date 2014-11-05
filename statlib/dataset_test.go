package statlib

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

func TestDataset_Blocks(t *testing.T) {
	d := NewDataset(1, 2, 3, 4, 5, 6, 7, 8)
	assert.Equal(t, []*Dataset{
		NewDataset(1, 2, 3, 4),
		NewDataset(5, 6, 7, 8),
	}, d.Blocks(2))

	assert.Equal(t, []*Dataset{
		NewDataset(1, 2),
		NewDataset(3, 4),
		NewDataset(5, 6),
		NewDataset(7, 8),
	}, d.Blocks(4))

	assert.Equal(t, []*Dataset{
		nil,
		NewDataset(1),
		NewDataset(2),
		NewDataset(3),
		NewDataset(4),
		nil,
		NewDataset(5),
		NewDataset(6),
		NewDataset(7),
		NewDataset(8),
	}, d.Blocks(10))
}

func TestDataset_Buckets(t *testing.T) {
	d := NewDataset(1.2, 3.7, 3.9, 4.0)

	assert.Equal(t, []*Dataset{
		NewDataset(1.2),
		NewDataset(3.7, 3.9, 4.0),
	}, d.Buckets(2))

	assert.Equal(t, []*Dataset{
		NewDataset(1.2),
		NewDataset(),
		NewDataset(),
		NewDataset(3.7, 3.9, 4.0),
	}, d.Buckets(4))

	d = NewDataset(1)
	assert.Equal(t, []*Dataset{
		NewDataset(1),
		NewDataset(),
		NewDataset(),
	}, d.Buckets(3))

	d = NewDataset(7, 7, 7, 7, 7)
	assert.Equal(t, []*Dataset{
		NewDataset(7, 7, 7, 7, 7),
		NewDataset(),
		NewDataset(),
	}, d.Buckets(3))
}

func TestDataset_Average(t *testing.T) {
	d := NewDataset(1, 2, 3)
	assert.Equal(t, 2, d.Average())
	d = NewDataset(1, 2, 6)
	assert.Equal(t, 3, d.Average())
	d = NewDataset(0.5, 1.5)
	assert.Equal(t, 1, d.Average())
	d = NewDataset(-5, 5)
	assert.Equal(t, 0, d.Average())
}

func TestDataset_Spread(t *testing.T) {
	d := NewDataset(1, 2, 3)
	assert.Equal(t, 2, d.Spread())
	d = NewDataset(1, 2, 6)
	assert.Equal(t, 5, d.Spread())
	d = NewDataset(0.5, 1.5)
	assert.Equal(t, 1, d.Spread())
	d = NewDataset(-5, 5)
	assert.Equal(t, 10, d.Spread())
}
