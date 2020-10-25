package jsonpointer_test

import (
	"testing"

	"github.com/yuya-takeyama/yamod/pkg/jsonpointer"
)

func TestRoot(t *testing.T) {
	ptr, err := jsonpointer.Parse("/")
	if err != nil {
		t.Errorf("error: %w", err)
		return
	}

	tokensLen := len(ptr.ReferenceTokens)
	if tokensLen != 0 {
		t.Errorf("Length of tokens must be 0, Actual: %d", tokensLen)
		return
	}
}

func TestSingleUnescapedReference(t *testing.T) {
	ptr, err := jsonpointer.Parse("/foo")
	if err != nil {
		t.Errorf("error: %w", err)
		return
	}

	tokensLen := len(ptr.ReferenceTokens)
	if tokensLen != 1 {
		t.Errorf("Length of tokens must be 1, Actual: %d", tokensLen)
		return
	}

	ref0 := ptr.ReferenceTokens[0].Reference
	if ptr.ReferenceTokens[0].Reference != "foo" {
		t.Errorf("First token must be foo, Actual: %s", ref0)
	}
}

func TestSingleReferenceContainsTilde(t *testing.T) {
	ptr, err := jsonpointer.Parse("/foo~0bar")
	if err != nil {
		t.Errorf("error: %w", err)
		return
	}

	tokensLen := len(ptr.ReferenceTokens)
	if tokensLen != 1 {
		t.Errorf("Length of tokens must be 1, Actual: %d", tokensLen)
		return
	}

	ref0 := ptr.ReferenceTokens[0].Reference
	if ptr.ReferenceTokens[0].Reference != "foo~bar" {
		t.Errorf("First token must be foo~bar, Actual: %s", ref0)
	}
}

func TestSingleReferenceContainsSlash(t *testing.T) {
	ptr, err := jsonpointer.Parse("/foo~1bar")
	if err != nil {
		t.Errorf("error: %w", err)
		return
	}

	tokensLen := len(ptr.ReferenceTokens)
	if tokensLen != 1 {
		t.Errorf("Length of tokens must be 1, Actual: %d", tokensLen)
		return
	}

	ref0 := ptr.ReferenceTokens[0].Reference
	if ptr.ReferenceTokens[0].Reference != "foo/bar" {
		t.Errorf("First token must be foo/bar, Actual: %s", ref0)
	}
}

func TestNestedReference(t *testing.T) {
	ptr, err := jsonpointer.Parse("/foo/bar/baz")
	if err != nil {
		t.Errorf("error: %w", err)
		return
	}

	tokensLen := len(ptr.ReferenceTokens)
	if tokensLen != 3 {
		t.Errorf("Length of tokens must be 3, Actual: %d", tokensLen)
		return
	}

	for i, expectedRef := range []string{"foo", "bar", "baz"} {
		actualRef := ptr.ReferenceTokens[i].Reference
		if actualRef != expectedRef {
			t.Errorf("ptr.ReferenceTokens[%d].Refeernce must be %s, Actual: %s", i, expectedRef, actualRef)
		}
	}
}

func TestPointerNotStartsWithSlash(t *testing.T) {
	ptr, err := jsonpointer.Parse("foo")
	if err == nil {
		t.Error("It must return err")
		return
	}

	if ptr != nil {
		t.Error("JSONPointer must be nil")
	}
}
