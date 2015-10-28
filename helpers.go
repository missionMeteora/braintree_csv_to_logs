package main

import (
	"flag"
	"os"
	"strconv"
)

func lineParseErr(pfx string, err error, i int) {
	stderr(pfx+" on line #"+strconv.Itoa(i)+": ", err)
}

func stderr(pfx string, err error) {
	os.Stderr.WriteString(pfx + err.Error() + "\n")
}

func getFlagLocs() (i, o, t string) {
	flag.StringVar(&i, "i", "", "Location of input file to process (Stdin by default)")
	flag.StringVar(&o, "o", "", "Location of output file (Stdout by default)")
	flag.Parse()
	return i, o, i + ".tmp"
}

func closeFile(f *os.File, err error, o, tmp string) {
	if f != nil {
		f.Close()

		if err == nil {
			os.Rename(tmp, o)
		} else {
			os.Remove(tmp)
		}
	}
}

func reportErrors(err error) {
	if err != nil {
		stderr("", err)
	}
}
