package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

func readLine(f *os.File, rc chan string) {
	reader := bufio.NewReader(f)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			close(rc)
		}
		rc <- input
	}
}

func main() {
	rc := make(chan string, 1)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc)

	go readLine(os.Stdin, rc)

	for {
		select {
		case input, ok := <-rc:
			fmt.Println(input)
			if !ok {
				rc = nil
				break
			}
		case sig := <-sc:
			fmt.Println(sig)
			os.Exit(1)
		}
	}
}
