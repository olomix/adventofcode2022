package internal

import (
	"bufio"
	"os"
)

func Perr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLines(path string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		f, err := os.Open(path)
		Perr(err)
		defer func() { Perr(f.Close()) }()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		Perr(scanner.Err())
	}()
	return ch
}
