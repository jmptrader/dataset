package stat

import "sort"

// Datasets are IMMUTABLE
type Dataset struct {
	name    string
	records []float64
	len     int
}

func NewDataset(records ...float64) *Dataset {
	if len(records) == 0 {
		records = []float64{0}
	}
	slice := sort.Float64Slice(records)
	slice.Sort()
	d := &Dataset{
		records: []float64(slice),
		len:     len(records),
	}
	return d
}

// Summary is struct used to hold the classic "five number summary".
type Summary struct {
	Min           float64
	LowerQuartile float64
	Median        float64
	UpperQuartile float64
	Max           float64
}

// Percentile retrieves the number at the `index`-th percentile of the
// distribution. `index` must be between 0 and 1.
func (d *Dataset) Percentile(index float64) float64 {
	intIndex := int(index * float64(d.len-1))
	record := d.records[intIndex]
	if d.len%2 == 0 {
		record = (record + d.records[intIndex+1]) / 2
	}
	return record
}

// FiveNumberSummary calculates and return a Summary object representing the
// five number summary of the dataset.
func (d *Dataset) FiveNumberSummary() Summary {
	summary := Summary{}

	if d.len > 0 {
		summary.Min = d.records[0]
		summary.LowerQuartile = d.Percentile(.25)
		summary.Median = d.Percentile(.5)
		summary.UpperQuartile = d.Percentile(.75)
		summary.Max = d.records[len(d.records)-1]
	}

	return summary
}
