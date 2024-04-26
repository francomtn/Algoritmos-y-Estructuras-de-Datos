package guia0

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	cont := 0
	for s.Scan() {
		if s.Text() == "s" {
			break
		}
		cont++
	}
	fmt.Println(cont)
}
