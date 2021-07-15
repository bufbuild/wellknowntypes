// Package main implements a utility that lists pairs of an input list from stdin.
//
// Example:
//
// Input:
//
//   foo bar baz
//
// Output:
//
//   foo bar
//   bar baz
//   baz bat
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bufbuild/wellknowntypes/internal/pkg/mainutil"
)

func main() {
	mainutil.Main(run)
}

func run(ctx context.Context) error {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	prev := ""
	next := ""
	for _, field := range strings.Fields(string(data)) {
		if err != nil {
			return err
		}
		if prev != "" {
			fmt.Println(prev, next)
		}
		prev = next
		next = field
	}
	if prev != "" {
		fmt.Println(prev, next)
	}
	return nil
}
