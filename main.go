package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	file := open()
	defer file.Close()

	lines := make(chan string)
	validate(lines)
	process(file, lines)
	defer close(lines)
}

func validate(lines <-chan string) {
	go func() {
		line := <-lines
		fmt.Println(line)
	}()
}

func process(file *os.File, lines chan<- string) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var wg sync.WaitGroup
	for scanner.Scan() {
		data := scanner.Text()
		wg.Add(1)
		go func(data string) {
			// fmt.Println(data)
			lines <- data
			wg.Done()
		}(data)
	}
	wg.Wait()
}

func open() *os.File {
	file, err := os.Open("data/valid.csv")
	if err != nil {
		panic(err)
	}
	return file
}
