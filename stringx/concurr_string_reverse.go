package stringx

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func StringReverse() {
	wg := sync.WaitGroup{}
	inCh := make(chan string, 128)
	outCh := make(chan string, 128)

	inFile, err := os.Open("input_file.txt")
	if err != nil {
		log.Fatal("Error opening input file")
	}

	// Start the second stage
	go reverseString(inCh, outCh)

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() { // splits on lines by default
		word := scanner.Text()
		//fmt.Println(word)

		wg.Add(1)
		go func(wrd string) {
			inCh <- wrd
			wg.Done()
		}(word)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error while scanning input file")
	}

	wg.Wait()
	close(inCh)

	for wrd := range outCh {
		fmt.Println("Word from output channel", wrd)
	}
}

func reverseString(inCh <-chan string, outCh chan<- string) {
	// Get Unicode code points.
	defer func() {
		close(outCh)
	}()
	for {
		word, ok := <-inCh
		if !ok || strings.TrimSpace(word) == "" {
			return
		}
		fmt.Printf("Word from input channel \"%s\"\n", word)
		rn := make([]rune, len(word))
		n := 0
		for _, r := range word {
			rn[n] = r
			n++
		}
		rn = rn[0:n]
		// Reverse
		for i := 0; i < n/2; i++ {
			rn[i], rn[n-1-i] = rn[n-1-i], rn[i]
		}
		// Convert back to UTF-8.
		out := string(rn)
		//fmt.Println(out)
		outCh <- out
	}
}
