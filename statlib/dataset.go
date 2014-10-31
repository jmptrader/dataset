package statlib

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

// Percentile retrieves the number at the `index`-th percentile of the
// distribution. `index` must be between 0 and 1. An `index` of 0.5 retrieves
// the median. If the dataset has an even number of records, the number is the
// average of the two nearest samples to the percentile.
func (d *Dataset) Percentile(index float64) float64 {
	if index > 1 || index < 0 {
		panic("Index provided to percentile must be between 0 and 1")
	}

	intIndex := int(index * float64(d.len-1))
	record := d.records[intIndex]
	if d.len%2 == 0 && intIndex < d.len-1 {
		record = (record + d.records[intIndex+1]) / 2
	}
	return record
}
