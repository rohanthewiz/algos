package stringx

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func ReverseStrings() (rvsStrs []string) {
	rvsStrs = make([]string, 0, 8)

	inFile, err := os.Open("input_file.txt")
	if err != nil {
		log.Fatal("Error opening input file")
		return
	}

	wg := sync.WaitGroup{}
	inCh := make(chan string, 128)
	outCh := make(chan string, 128)

	// Start the second stage
	// Receive input from the inCh
	// Return output over the outCh
	go processStringReverse(inCh, outCh)

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() { // splits on lines by default
		word := scanner.Text()

		wg.Add(1)
		go func(wrd string) {
			inCh <- wrd
			wg.Done()
		}(word) // pass the current value in
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error while scanning input file")
	}

	wg.Wait()
	close(inCh)

	for wrd := range outCh {
		fmt.Println("Word from output channel", wrd)
		rvsStrs = append(rvsStrs, wrd)
	}

	return
}

func processStringReverse(inCh <-chan string, outCh chan<- string) {
	// Get Unicode code points.
	defer close(outCh)
	for {
		word, ok := <-inCh
		fmt.Printf("Word from input channel \"%s\"\n", word)
		if !ok { // channel is closed
			return
		}

		outCh <- reverseString(word)
	}
}

func reverseString(inStr string) (outStr string) {
	if inStr = strings.TrimSpace(inStr); inStr == "" {
		return
	}

	rn := make([]rune, len(inStr))
	n := 0
	for _, r := range inStr {
		rn[n] = r
		n++
	}
	rn = rn[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rn[i], rn[n-1-i] = rn[n-1-i], rn[i]
	}
	// Convert back to UTF-8.
	outStr = string(rn)

	return
}
