package stringx
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

// TODO this needs to be finished
// and tests added
func TestStringReverse() {
	wg := sync.WaitGroup{}
	inCh := make(chan string)
	outCh := make(chan string)

	inFile, err := os.Open("input_file.txt")
	if err != nil {
		log.Fatal("Error opening input file")
	}
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() { // splits on lines by default
		word := scanner.Text()
		fmt.Println(word)
		inCh <- word

		wg.Add(1)
		go func() {
			ReverseString(inCh, outCh)
			wg.Done()
		}()
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error while scanning input file")
	}
	wg.Wait()
	close(inCh)
}

func ReverseString(inCh <-chan string, outCh chan<- string) {
	// Get Unicode code points.
	n := 0
	for {
		word := <- inCh
		rune := make([]rune, len(word))
		for _, r := range word {
			rune[n] = r
			n++
		}
		rune = rune[0:n]
		// Reverse
		for i := 0; i < n/2; i++ {
			rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
		}
		// Convert back to UTF-8.
		out := string(rune)
		fmt.Println(out)
		outCh <- out
	}
}
