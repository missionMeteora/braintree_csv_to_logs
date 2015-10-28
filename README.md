# braintree_csv_to_logs [![GoDoc](https://godoc.org/github.com/missionMeteora/braintree_csv_to_logs?status.svg)](https://godoc.org/github.com/missionMeteora/braintree_csv_to_logs) ![Status](https://img.shields.io/badge/status-beta-yellow.svg)

braintree_csv_to_logs is a library which will process a braintree report in CSV format. It will output the report as a tar containing logs split up by year and month. 

## Usage - Command line
``` bash
# Input file and output file
./braintree_csv_to_logs -i report.csv -o archive.tar

# Stdin and Stdout
cat report.csv | ./braintree_csv_to_logs > archive.tar
```