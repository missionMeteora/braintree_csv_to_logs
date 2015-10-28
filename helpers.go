package main

import (
	"os"
	"strconv"
)

func lineParseErr(pfx string, err error, i int) {
	stderr(pfx+" on line #"+strconv.Itoa(i)+": ", err)
}

func stderr(pfx string, err error) {
	os.Stderr.WriteString(pfx + err.Error() + "\n")
}
