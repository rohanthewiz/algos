package stringx

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MakeUpper() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()
		out := ""

		makeUpper := true

		for i := 0; i < len(line); i++ {
			ch := string(line[i])
			if ch == " " {
				out += ch
				continue
			}

			if makeUpper {
				out += strings.ToUpper(ch)
			} else {
				out += strings.ToLower(ch)
			}
			makeUpper = !makeUpper
		}

		fmt.Println(out)
	}
}
