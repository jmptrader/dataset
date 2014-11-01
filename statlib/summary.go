package statlib

import (
	"bytes"
	"fmt"
	"strings"
)

// Summary is struct used to hold the classic "five number summary".
type Summary struct {
	Min           float64
	LowerQuartile float64
	Median        float64
	UpperQuartile float64
	Max           float64
}

func (s Summary) String() string {
	return s.AsTable(28)
}

// AsTable returns a formatted, humanreadable string that attempts to fit within
// `width` columns.
func (s Summary) AsTable(width int) string {
	clampWidth(&width)
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Five-number Summary\n")
	fmt.Fprintf(buf, "%s\n", strings.Repeat("-", width))
	extraSpace := width - 18
	fmtStr := fmt.Sprintf("%%%dg", extraSpace)
	fmt.Fprintf(buf, "Min               "+fmtStr+"\n", s.Min)
	fmt.Fprintf(buf, "Lower Quartile    "+fmtStr+"\n", s.LowerQuartile)
	fmt.Fprintf(buf, "Median            "+fmtStr+"\n", s.Median)
	fmt.Fprintf(buf, "Upper Quartile    "+fmtStr+"\n", s.UpperQuartile)
	fmt.Fprintf(buf, "Max               "+fmtStr, s.Max)
	return buf.String()
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
