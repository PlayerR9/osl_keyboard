package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/PlayerR9/osl_keyword/pkg"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./osl_keyboard <input_file> <output_file>")
		os.Exit(1)
	}

	err := command(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

func command(args []string) error {
	input_file := args[0]

	data, err := os.ReadFile(input_file)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	lines := bytes.Split(data, []byte("\n"))

	_, err = fmt.Println("Converting to OSL...")
	if err != nil {
		return fmt.Errorf("error printing: %w", err)
	}

	var res []string

	for i, line := range lines {
		words, err := pkg.Tokenize(line)

		if err != nil {
			return fmt.Errorf("error tokenizing line %d: %w", i+1, err)
		}

		err = pkg.CheckValidity(words)
		if err != nil {
			return fmt.Errorf("check validity error on line %d: %w", i+1, err)
		}

		pkg.FinalTweaks(words)

		str := pkg.SentenceString(words)

		res = append(res, str)
	}

	_, err = fmt.Println("Writing to file...")
	if err != nil {
		return fmt.Errorf("error printing: %w", err)
	}

	err = os.WriteFile(args[1], []byte(strings.Join(res, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	_, err = fmt.Println("Done!")
	if err != nil {
		return fmt.Errorf("error printing: %w", err)
	}

	return nil
}
