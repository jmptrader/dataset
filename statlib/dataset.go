package statlib

import (
	"fmt"
	"sort"
)

// Datasets are IMMUTABLE
type Dataset struct {
	records []float64
	len     int
}

// NewDataset constructs a dataset from the provided records. The records can be
// provided in any order and are automatically sorted (but NOT de-duplicated).
// If no records are provided, this function will return nil.
func NewDataset(records ...float64) *Dataset {
	if len(records) == 0 {
		return nil
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

// Range returns the maximum minus the minimum record.
func (d *Dataset) Range() float64 {
	return d.records[d.len-1] - d.records[0]
}

// Blocks splits the dataset into `n` blocks containing an approximately the
// same number of elements in each block. The blocks are sorted from smallest
// value to largest value.
func (d *Dataset) Blocks(n int) []*Dataset {
	blocks := make([]*Dataset, n)
	for i := range blocks {
		blocks[i] = NewDataset(d.records[i*d.len/n : (i+1)*d.len/n]...)
	}
	return blocks
}

// Buckets splits the dataset into `n` buckets. Some buckets may contain more
// elements than others because each bucket will contain elements in the same
// size of range as every other bucket.
func (d *Dataset) Buckets(n int) []*Dataset {
	if n <= 1 {
		return []*Dataset{d}
	}
	buckets := make([][]float64, n)
	min := d.records[0]
	bucketSize := d.Range() / float64(n)
	for _, record := range d.records {
		// how far is this record from the bottom of the distribution?
		difference := record - min

		index := 0
		if difference != 0 {
			index = int((record - min) / bucketSize)
			// this conditional is here because the final record
			// will be exactly equal to the topmost bracket (a.k.a.
			// the start of the where the next bucket would be).
			if index >= len(buckets) {
				index--
			}
		}
		buckets[index] = append(buckets[index], record)
	}

	// convert our slices of float64s to datasets and return
	bucketDatasets := make([]*Dataset, n)
	for i, bucket := range buckets {
		bucketDatasets[i] = NewDataset(bucket...)
	}
	return bucketDatasets
}

// Len returns the count of elements in the Dataset
func (d *Dataset) Len() int       { return d.len }
func (d *Dataset) String() string { return fmt.Sprintf("Dataset%v", d.records) }
