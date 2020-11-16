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

func parseLogs(w *tabwriter.Writer, f string, logs []string) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	count := 0
	total := 0
	memMax := ""
	for _, l := range logs {
		mem := parseLog(l)
		memMax = mem.max
		if mem.limitReached {
			count++
		}
		total++
	}

	if count > 0 {
		fmt.Fprintf(w, "\n %s%s\t%s\t%s\t%s", string(colorRed), f, memMax, fmt.Sprintf("%d / %d", count, total), string(colorReset))
	} else {
		fmt.Fprintf(w, "\n %s%s\t%s\t%s\t%s", string(colorGreen), f, memMax, fmt.Sprintf("%d / %d", count, total), string(colorReset))
	}
}

func report(logs []lambdaLogs) {
	colorReset := "\033[0m"
	colorCyan := "\033[36m"

	w := new(tabwriter.Writer)
	flags := tabwriter.AlignRight
	w.Init(os.Stdout, 8, 8, 2, ' ', flags)

	fmt.Fprintf(w, "\n %s%s\t%s\t%s\t%s", string(colorCyan), "Lambda", "Max", "OOM", string(colorReset))
	fmt.Fprintf(w, "\n %s%s\t%s\t%s\t%s", string(colorCyan), "------", "---", "---", string(colorReset))

	for _, lambda := range logs {
		parseLogs(w, lambda.name, lambda.logs)
	}

	defer fmt.Println("\n")
	defer w.Flush()
}
