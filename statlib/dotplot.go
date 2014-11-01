package statlib

import (
	"bytes"
	"fmt"
)

// Dotplot will make an attempt to fit within the specified `width` in columns.
func (d *Dataset) Dotplot(width int) string {
	bucketCount := width / 6 // we need about 6 columns per dotplot
	buckets := d.Buckets(bucketCount)
	bucketSize := d.Range() / float64(bucketCount)

	buf := new(bytes.Buffer)

	maxHeight := 0
	bottomLine := ``
	for i, bucket := range buckets {
		// TODO: need a good way to force this string to be below six
		// characters without truncating and ending up with a result
		// that looks like the right number but could be off by an order
		// of magnitude
		s := fmt.Sprintf("%5.2f ", d.records[0]+(bucketSize*float64(i)))
		bottomLine = bottomLine + s
		if bucket != nil && bucket.len > maxHeight {
			maxHeight = bucket.len
		}
	}

	for i := 0; i < maxHeight; i++ {
		line := ``
		for _, bucket := range buckets {
			if bucket != nil && bucket.len >= maxHeight-i {
				line = line + `  .   `
			} else {
				line = line + `      `
			}
		}
		fmt.Fprintf(buf, "%s\n", line)
	}
	fmt.Fprintf(buf, "%s", bottomLine)
	return buf.String()
}
