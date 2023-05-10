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

	process(file)
}

func open() *os.File {
	file, err := os.Open("data/valid.csv")
	if err != nil {
		panic(err)
	}
	return file
}

func process(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var wg sync.WaitGroup
	for scanner.Scan() {
		data := scanner.Text()
		wg.Add(1)
		go func(data string) {
			fmt.Println(data)
			wg.Done()
		}(data)
	}
	wg.Wait()
}
