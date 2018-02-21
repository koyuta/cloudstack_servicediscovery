package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	endpoint  = flag.String("endpoint", "", "Endpoint url of cloudstack")
	apiKey    = flag.String("api-key", "", "API key of cloudstack")
	secretKey = flag.String("secret-key", "", "Secret key of cloudstack")
	groups    = flag.String("groups", "", "List of groups separated by comma")
	labels    = flag.String("labels", "", "List of labels (e.g. \"job:mysql,zone:eu-east\")")
	filename  = flag.String("filename", "", "Output json file name that specified \"file_sd_config\"")
	port      = flag.Int("port", 9090, "Suffix port number")

	help = flag.Bool("help", false, "Print this help message and exit")
)

func main() {
	flag.Parse()
	flag.Usage = printUsage
	if *help {
		printUsage()
	}
	os.Exit(printOnError(run()))
}

func printOnError(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return 1
	}
	return 0
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage:
  %s [OPTIONS]

OPTIONS:
`, os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
