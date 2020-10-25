package yamlnodefinder

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/yuya-takeyama/yamod/pkg/jsonpointer"
	"gopkg.in/yaml.v3"
)

// Find returns YAML Node using JSON Pointer
func Find(doc *yaml.Node, ptr jsonpointer.JSONPointer) (*yaml.Node, error) {
	if len(ptr.ReferenceTokens) == 0 {
		return doc, nil
	}

	node := doc.Content[0]
	return find(node, ptr.ReferenceTokens)
}

func find(node *yaml.Node, refTokens []jsonpointer.ReferenceToken) (*yaml.Node, error) {
	if node.Kind == yaml.MappingNode {
		return findFromMapping(node, refTokens)
	} else if node.Kind == yaml.SequenceNode {
		return findFromSequence(node, refTokens)
	}

	return nil, fmt.Errorf("Unsupported kind: %s", inspectKind(node.Kind))
}

func findFromMapping(doc *yaml.Node, refTokens []jsonpointer.ReferenceToken) (*yaml.Node, error) {
	token := refTokens[0]
	isKey := true
	returnNext := false
	for _, c := range doc.Content {
		if isKey && c.Value == token.Reference {
			returnNext = true
		} else if !isKey && returnNext {
			if len(refTokens) > 1 {
				return find(c, refTokens[1:])
			}

			return c, nil
		}

		isKey = !isKey
	}

	return nil, errors.New("Node not found")
}

func findFromSequence(doc *yaml.Node, refTokens []jsonpointer.ReferenceToken) (*yaml.Node, error) {
	token := refTokens[0]
	for i, c := range doc.Content {
		if strconv.Itoa(i) == token.Reference {
			if len(refTokens) > 1 {
				return find(c, refTokens[1:])
			}

			return c, nil
		}
	}

	return nil, errors.New("Node not found")
}

func inspectKind(kind yaml.Kind) string {
	switch kind {
	case yaml.DocumentNode:
		return "Document"
	case yaml.SequenceNode:
		return "Sequence"
	case yaml.MappingNode:
		return "Mapping"
	case yaml.ScalarNode:
		return "Scalar"
	case yaml.AliasNode:
		return "Alias"

	}

	return "Unknown"
}
