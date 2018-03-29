package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Lang struct {
	Name string
	Year int
	Url  string
}

func parse(f func(lang Lang)) {
	input, err := os.Open("D:\\go\\src\\test\\tour\\json\\lang.json")
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(input)

	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}

func count(name, url string, ch chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	defer r.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("%s: %s", name, err)
		return
	}

	n, _ := io.Copy(ioutil.Discard, r.Body)

	ch <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}

func main() {
	start := time.Now()
	c := make(chan string)
	n := 0

	parse(func(lang Lang) {
		fmt.Printf("%#v\n", lang)
		n++
		go func() {
			count(lang.Name, lang.Url, c)
		}()
	})

	for i := 0; i < n; i++ {
		fmt.Print(<-c)
	}

	fmt.Printf("%s", time.Since(start))
}
