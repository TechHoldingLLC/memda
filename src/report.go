package main

import (
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"
)

type lambdaMemory struct {
	requestID    string
	max          string
	used         string
	limitReached bool
}

func parseLog(line string) lambdaMemory {

	reID := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	reqID := reID.FindString(line)
	reMax := regexp.MustCompile("Memory Size: ([0-9]+ MB)")
	max := reMax.FindStringSubmatch(line)
	reUsed := regexp.MustCompile("Memory Used: ([0-9]+ MB)")
	used := reUsed.FindStringSubmatch(line)

	limitReached := max[1] == used[1]
	mem := lambdaMemory{reqID, max[1], used[1], limitReached}
	return mem
}

func report(f string, logs []string) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorCyan := "\033[36m"
	colorGreen := "\033[32m"

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s%s\t%s\t%s\t%s", string(colorCyan), f, "Max", "Used", string(colorReset))

	count := 0
	for _, l := range logs {
		mem := parseLog(l)
		if mem.limitReached {
			count++
			fmt.Fprintf(w, "\n %s%s\t%s\t%s\t%s", string(colorRed), mem.requestID, mem.max, mem.used, string(colorReset))
		} else {
			fmt.Fprintf(w, "\n %s%s\t%s\t%s\t%s", string(colorReset), mem.requestID, mem.max, mem.used, string(colorReset))
		}
	}

	defer fmt.Println("\n")

	if count > 0 {
		defer fmt.Printf(string(colorRed)+f+" reached its memory limit %d times"+string(colorReset)+"\n", count)
	} else {
		defer fmt.Println(string(colorGreen) + f + " never reached its memory limit" + string(colorReset))
	}

	defer fmt.Println("\n")
	defer w.Flush()
}
