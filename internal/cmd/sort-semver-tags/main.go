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
	"github.com/bufbuild/wellknowntypes/internal/pkg/semverutil"
)

func main() {
	mainutil.Main(run)
}

func run(ctx context.Context) error {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	var semverTagNames []string
	for _, field := range strings.Fields(string(data)) {
		semverTagNames = append(semverTagNames, field)
	}
	if err := semverutil.ValidateSemverTagNames(semverTagNames); err != nil {
		return err
	}
	semverutil.SortSemverTagNames(semverTagNames)
	for _, semverTagName := range semverTagNames {
		fmt.Println(semverTagName)
	}
	return nil
}
