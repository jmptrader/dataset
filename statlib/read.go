package statlib

import (
	"errors"
	"strconv"
)

// ReadDataset takes a slice of bytes and returns the corresponding Dataset,
// splitting the bytes by whitespace.
func ReadDataset(bytes []byte) (*Dataset, error) {
	var (
		records      = []float64{} // list of parsed records
		recordString = []byte{}    // current record
	)

	for _, b := range bytes {
		switch b {
		case ' ', '\n', '\t':
			if len(recordString) < 1 {
				continue
			}

			f, err := strconv.ParseFloat(string(recordString), 64)
			if err != nil {
				return nil, err
			}
			records = append(records, f)
			recordString = []byte{}
		default:
			recordString = append(recordString, b)
		}
	}

	if len(recordString) > 0 {
		f, err := strconv.ParseFloat(string(recordString), 64)
		if err != nil {
			return nil, err
		}
		records = append(records, f)
	}

	if len(records) < 1 {
		return nil, errors.New(`no records found in dataset`)
	}

	return NewDataset(records...), nil
}
