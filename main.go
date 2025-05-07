package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	input := flag.String("input", "", "Path to input log file")
	output := flag.String("output", "", "Path to output file (optional)")
	startStr := flag.String("start", "", "Start time (format: 2006/01/02 15:04:05, assumed UTC)")
	endStr := flag.String("end", "", "End time (format: 2006/01/02 15:04:05, assumed UTC)")
	flag.Parse()

	if *input == "" || *startStr == "" || *endStr == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Layouts for parsing timestamps
	const lineTimeLayout = "2006/01/02 15:04:05.000 -07:00"
	const flagTimeLayout = "2006/01/02 15:04:05"

	// Parse start and end times in UTC
	startTime, err := time.ParseInLocation(flagTimeLayout, *startStr, time.UTC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid start time: %v\n", err)
		os.Exit(1)
	}
	endTime, err := time.ParseInLocation(flagTimeLayout, *endStr, time.UTC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid end time: %v\n", err)
		os.Exit(1)
	}

	// Open input file
	file, err := os.Open(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Setup output writer
	var writer io.Writer = os.Stdout
	if *output != "" {
		outFile, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot create output file: %v\n", err)
			os.Exit(1)
		}
		defer outFile.Close()
		writer = outFile
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Extract timestamp between brackets
		endIdx := strings.Index(line, "]")
		if endIdx <= 1 {
			continue
		}
		tsStr := line[1:endIdx]
		ts, err := time.Parse(lineTimeLayout, tsStr)
		if err != nil {
			continue
		}
		// Compare in UTC
		ts = ts.UTC()
		if (ts.Equal(startTime) || ts.After(startTime)) && (ts.Equal(endTime) || ts.Before(endTime)) {
			fmt.Fprintln(writer, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}
}
