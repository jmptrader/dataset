
Dataset
=======

Experimental Go utility to analyze and quickly plot a series of floating point
numbers on the command line.

Install:

```sh
	go install github.com/alexandercampbell/dataset
```

Usage:

```sh
	dataset sample-data/gauss_01.txt
	dataset sample-data/gauss_01.txt
```

Dataset works like many shell utilities.

```sh
	# Read from file
	dataset <some-file>

	# Read from stdin
	dataset
	dataset -

	# Flag messages
	dataset --help
```

Dataset ignores everything except numbers, so you can use it on CSVs or whatever
other format you have.

```sh
	echo "1234, 123, 143" | dataset
```



