package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := scanner.Text()

		switch cmd {
		case "exit":
			os.Exit(0)
		default:
			output := fmt.Sprintf("%s: command not found", cmd)
			fmt.Println(output)
		}
		fmt.Fprint(os.Stdout, "$ ")
	}

}
