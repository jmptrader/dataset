package statlib

import (
	"bytes"
	"fmt"
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
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Five-number Summary         \n")
	fmt.Fprintf(buf, "----------------------------\n")
	fmt.Fprintf(buf, "Min               %10g\n", s.Min)
	fmt.Fprintf(buf, "Lower Quartile    %10g\n", s.LowerQuartile)
	fmt.Fprintf(buf, "Median            %10g\n", s.Median)
	fmt.Fprintf(buf, "Upper Quartile    %10g\n", s.UpperQuartile)
	fmt.Fprintf(buf, "Max               %10g", s.Max)
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
