package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/robherley/guesslang-go/pkg/guesser"
)

func getTextFromStdio() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""

	for {
		if scanner.Scan() {
			text = text + scanner.Text()
		} else {
			if err := scanner.Err(); err != nil {
				return "", err
			} else {
				break
			}
		}
	}

	return text, nil
}

func guess(snippet string) (string, error) {
	gsr, err := guesser.New()
	if err != nil {
		return "", err
	}

	answer, err := gsr.Guess(snippet)
	if err != nil {
		return "", err
	}

	if len(answer.Predictions) == 0 {
		return "", nil
	}

	return answer.Predictions[0].Language, nil
}

func main() {
	text, err := getTextFromStdio()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		return
	}

	language, err := guess(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		return
	}

	fmt.Fprintf(os.Stdout, "%s", language)
	os.Exit(0)
}
