package semverutil

import (
	"fmt"
	"sort"

	"golang.org/x/mod/semver"
)

// SemVerTagNames gets the valid semver tag names from the list.
func SemverTagNames(tagNames []string) []string {
	semverTagNames := make([]string, 0, len(tagNames))
	for _, tagName := range tagNames {
		if semver.IsValid(tagName) {
			semverTagNames = append(semverTagNames, tagName)
		}
	}
	return semverTagNames
}

// ValidateSemverTagNames validates that all tag names in the list are valid semver tag names.
func ValidateSemverTagNames(tagNames []string) error {
	for _, tagName := range tagNames {
		if !semver.IsValid(tagName) {
			return fmt.Errorf("%q is not a valid SemVer tag name", tagName)
		}
	}
	return nil
}

// StableSemverTagNames gets the stable semver tag names from the list.
func StableSemverTagNames(semverTagNames []string) []string {
	stableTagNames := make([]string, 0, len(semverTagNames))
	for _, semverTagName := range semverTagNames {
		if semver.Prerelease(semverTagName) == "" && semver.Build(semverTagName) == "" {
			stableTagNames = append(stableTagNames, semverTagName)
		}
	}
	return stableTagNames
}

// SemverTagNamesAtLeast gets the semver tag names that are at least the minimum tag name.
func SemverTagNamesAtLeast(semverTagNames []string, minimumSemverTagName string) []string {
	atLeastTagNames := make([]string, 0, len(semverTagNames))
	for _, semverTagName := range semverTagNames {
		if semver.Compare(semverTagName, minimumSemverTagName) >= 0 {
			atLeastTagNames = append(atLeastTagNames, semverTagName)
		}
	}
	return atLeastTagNames
}

// SemverTagNamesExcept gets the semver tag names minus the skip tag names.
func SemverTagNamesExcept(semverTagNames []string, skipTagNames map[string]struct{}) []string {
	exceptTagNames := make([]string, 0, len(semverTagNames))
	for _, semverTagName := range semverTagNames {
		if _, ok := skipTagNames[semverTagName]; !ok {
			exceptTagNames = append(exceptTagNames, semverTagName)
		}
	}
	return exceptTagNames
}

// SortSemverTagNames sorts the semver tag names.
func SortSemverTagNames(semverTagNames []string) {
	sort.Slice(
		semverTagNames,
		func(i int, j int) bool {
			return semver.Compare(
				semverTagNames[i],
				semverTagNames[j],
			) < 0
		},
	)
}
