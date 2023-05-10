package main

import (
	"bufio"
	"os"
	"time"
)

func main() {
	file := open()
	defer file.Close()
	lines := process(file)
	validate(lines)
}

func validate(lines <-chan string) {
	go func() {
		// line := <-lines
		// fmt.Println(line)
	}()
}

func process(file *os.File) <-chan string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make(chan string)
	defer close(lines)
	for scanner.Scan() {
		data := scanner.Text()
		lines <- data
	}
	return lines
}

func open() *os.File {
	file, err := os.Open("data/valid.csv")
	if err != nil {
		panic(err)
	}
	return file
}

type order struct {
	id      string
	value   float32
	express bool
	date    time.Time
}

// func parse(line string) order {
// 	values := strings.Split(line, ",")
// 	var order order
// 	order.id = values[0]
// 	order.value = strconv.ParseFloat(values[1])
// 	order.express = values[2]
// 	order.date = values[3]
// }
