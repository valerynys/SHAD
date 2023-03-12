//go:build !solution

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:] // получаем список URL'ов из аргументов командной строки

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Ошибка при запросе %s: %s\n", url, err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Ошибка при чтении содержимого страницы %s: %s\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("Содержимое страницы %s:\n%s\n", url, body)
	}
}
