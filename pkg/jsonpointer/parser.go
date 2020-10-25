package jsonpointer

import (
	"errors"
	"strings"
)

var errMustStartWithSlash error = errors.New(`pointer must start with "/"`)

// Parse parses JSON Pointer string and returns JSONPointer struct.
func Parse(pointer string) (*JSONPointer, error) {
	tokenStrs := strings.Split(pointer, "/")
	tokens := make([]ReferenceToken, len(tokenStrs)-1)

	if tokenStrs[0] != "" {
		return nil, errMustStartWithSlash
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
