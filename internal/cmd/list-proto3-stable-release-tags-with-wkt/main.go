// Package main implements a utility that lists all GitHub Releases
// tag names for stable releases of github.com/protocolbuffers/protobuf
// that contain the Well-Known Types.
//
// This is generally all versions >=3.0.0 except for 3.4.1.
// The printed list is in sorted SemVer order.
//
// If GITHUB_TOKEN is set, this is used for authentication.
package main

import (
	"context"
	"fmt"

	"github.com/bufbuild/wellknowntypes/internal/pkg/githubutil"
	"github.com/bufbuild/wellknowntypes/internal/pkg/mainutil"
	"github.com/bufbuild/wellknowntypes/internal/pkg/semverutil"
)

const (
	githubProtobufOwner          = "protocolbuffers"
	githubProtobufRepository     = "protobuf"
	protobufMinimumSemverTagName = "v3.0.0"
)

var (
	protobufSkipTagNames = map[string]struct{}{
		// v3.4.1 did not have protoc attached to it
		"v3.4.1": struct{}{},
	}
)

func main() {
	mainutil.Main(run)
}

func run(ctx context.Context) error {
	httpClient := githubutil.NewHTTPClient(ctx)
	githubClient := githubutil.NewGithubClient(httpClient)

	tagNames, err := githubutil.AllReleaseTagNames(
		ctx,
		githubClient,
		githubProtobufOwner,
		githubProtobufRepository,
	)
	if err != nil {
		return err
	}
	semverTagNames := semverutil.SemverTagNamesExcept(
		semverutil.SemverTagNamesAtLeast(
			semverutil.StableSemverTagNames(
				semverutil.SemverTagNames(
					tagNames,
				),
			),
			protobufMinimumSemverTagName,
		),
		protobufSkipTagNames,
	)
	semverutil.SortSemverTagNames(semverTagNames)

	for _, semverTagName := range semverTagNames {
		fmt.Println(semverTagName)
	}
	return nil
}
