# log-filter

## Installation

Build the binary:
```
go build -o log-filter main.go
```

## Usage

Filter a log file by time range:
```
./log-filter -input=path/to/logfile.log \
             -start="2025/04/25 10:27:00" \
             -end="2025/04/25 10:43:00" \
             -output=filtered.log
```

- input: path to the source log file
- start: inclusive start time (UTC) in format `YYYY/MM/DD HH:MM:SS`
- end: inclusive end time (UTC) in format `YYYY/MM/DD HH:MM:SS`
- output: (optional) path to write the filtered log; defaults to stdout