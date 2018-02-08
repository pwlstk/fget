package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	source = flag.String("source", "source.txt", "A file with URLs to files. Each URL should be in a separate line.")
	format = flag.String("format", "mp3", "Files format.")
)

func init() {

	flag.Parse()
}

func main() {

	file, err := os.Open(*source)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()

	i := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		download(scanner.Text(), fmt.Sprintf("%03d", i))
		i++
	}
}

func download(url, filename string) {

	log.Printf("Downloading %s", url)

	file, err := os.Create(fmt.Sprintf("%s.%s", filename, *format))
	if err != nil {

		log.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {

		log.Printf("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		log.Printf("Error: %v", fmt.Errorf("Unexpected HTTP status %s.", resp.Status))
		return
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {

		log.Printf("Error: %v", err)
		return
	}
}
