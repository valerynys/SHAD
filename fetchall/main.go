//go:build !solution

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Println("Usage: fetchall URL1 [URL2...]")
		os.Exit(1)
	}

	start := time.Now()
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			start = time.Now()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
				return
			}
			defer resp.Body.Close()
			_, err = io.Copy(ioutil.Discard, resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
				return
			}
			elapsed := time.Since(start).Seconds()
			fmt.Printf("%-5s %7d  %s\n", strconv.FormatFloat(elapsed, 'f', 2, 64)+"s", resp.ContentLength, url)
		}(i, url)
	}
	wg.Wait()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
