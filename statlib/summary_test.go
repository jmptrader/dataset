package statlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataset_FiveNumberSummary(t *testing.T) {
	// Test the case of a single data point.
	d := NewDataset(12)
	assert.Equal(t, Summary{
		Min:           12,
		LowerQuartile: 12,
		Median:        12,
		UpperQuartile: 12,
		Max:           12,
	}, d.FiveNumberSummary())

	// Test a more complicated summary
	d = NewDataset(12, 14.3, 19, 25)
	assert.Equal(t, Summary{
		Min:           12,
		LowerQuartile: 13.15,
		Median:        16.65,
		UpperQuartile: 22,
		Max:           25,
	}, d.FiveNumberSummary())
}
