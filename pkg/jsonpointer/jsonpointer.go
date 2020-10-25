package jsonpointer

// JSONPointer contains refeernces represent JSON Pointer.
type JSONPointer struct {
	ReferenceTokens []ReferenceToken
}

// ReferenceToken represents a token splitted by slashes.
type ReferenceToken struct {
	Reference string
}
