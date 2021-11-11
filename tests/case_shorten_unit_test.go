package tests

import (
	"shorten-link/pkg/app/logic"

	"testing"
	"fmt"
)

func TestShorten(t *testing.T) {
	var tests = []struct{
		longLink string
		resultLink string
	}{
		{"https://gobyexample.com", "2l1V"},
		{"https://gobyexample.com", "l1V2"},
	}

	for _, testrun := range tests {
		testName := fmt.Sprintf("URL: %s", testrun.longLink)
		t.Run(testName, func(t *testing.T) {
			result := logic.ShortenLink(testrun.longLink).Hash
			if result != testrun.resultLink {
                t.Errorf("got %s, want %s", result, testrun.resultLink)
            }
			fmt.Println(result)
		})
	}
}