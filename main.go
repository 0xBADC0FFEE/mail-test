package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/0xBADC0FFEE/mail-test/pkg/manager"
	"github.com/0xBADC0FFEE/mail-test/pkg/sources"
)

func main() {
	var (
		typeFlag   InputType
		stringFlag string
		kFlag      int
	)

	flag.Var(&typeFlag, "type", "Source type")
	flag.StringVar(&stringFlag, "string", "Go", "String to count")
	flag.IntVar(&kFlag, "k", 5, "Maximum workers")
	flag.Parse()

	inputReader := bufio.NewScanner(os.Stdin)

	jobsManager := manager.New(kFlag)
	jobsManager.Wake()

	results := make(chan Result)

	jobs := 0
	completed := 0

	go func() {

		for inputReader.Scan() {
			jobs++
			path := inputReader.Text()

			jobsManager.Add(func() {
				var source sources.Source

				switch typeFlag {
				case URL:
					source = sources.NewUrlReader(path)
				case FILE:
					source = sources.NewFileReader(path)
				}

				data, err := source.Get()
				if err != nil {
					println(fmt.Sprintf("Error getting data from %s: %v", path, err))
				}

				results <- Result{
					Source: path,
					Count:  bytes.Count(data, []byte(stringFlag)),
					Error:  err,
				}

				completed++

				if completed >= jobs {
					close(results)
				}

			})

		}

	}()

	total := 0

	for result := range results {

		if result.Error != nil {
			log.Printf("Failed counting strings for %s: %v", result.Source, result.Error)
			continue
		}

		log.Printf("Count for %s: %d", result.Source, result.Count)
		total += result.Count
	}

	log.Printf("Total: %d", total)
}

type Result struct {
	Source string
	Count  int
	Error  error
}
