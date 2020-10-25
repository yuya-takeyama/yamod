package jsonpointer

import (
	"errors"
	"strings"
)

// Parse parses JSON Pointer string and returns JSONPointer struct
func Parse(pointer string) (*JSONPointer, error) {
	tokenStrs := strings.Split(pointer, "/")
	var tokens []ReferenceToken

	if tokenStrs[0] != "" {
		return nil, errors.New(`pointer must start with "/"`)
	}

	if len(tokenStrs) == 2 && tokenStrs[1] == "" {
		return &JSONPointer{
			ReferenceTokens: tokens,
		}, nil
	}

	for _, tokenStr := range tokenStrs[1:] {
		tokens = append(tokens, ReferenceToken{
			Reference: unescape(tokenStr),
		})
	}

	return &JSONPointer{
		ReferenceTokens: tokens,
	}, nil
}

func unescape(reference string) string {
	return strings.ReplaceAll(strings.ReplaceAll(reference, "~0", "~"), "~1", "/")
}






