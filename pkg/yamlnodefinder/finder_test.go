package yamlnodefinder_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yuya-takeyama/yamod/pkg/jsonpointer"
	"github.com/yuya-takeyama/yamod/pkg/yamlnodefinder"
	"gopkg.in/yaml.v3"
)

const inputYAML string = `---
foo: FOO
bar: BAR
baz:
  one:
    uno: UNO
    dos: DOS
    tres: TRES
  two:
    - 1
    - 2
    - 3
  three: true
  four: null
`

func TestFindRoot(t *testing.T) {
	doc := fixture()
	ptr, _ := jsonpointer.Parse("/")
	node, err := yamlnodefinder.Find(doc, *ptr)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	if node != doc {
		t.Error("found Node must be the input Node")
	}
}

func TestFindSingleReference(t *testing.T) {
	doc := fixture()
	ptr, _ := jsonpointer.Parse("/foo")
	node, err := yamlnodefinder.Find(doc, *ptr)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	expected := "FOO"
	if node.Value != expected {
		t.Errorf("Wrong Node is returned\nExpected: %s, Actual: %s", expected, node.Value)
	}
}

func TestFindFromNestedMappings(t *testing.T) {
	doc := fixture()
	ptr, _ := jsonpointer.Parse("/baz/one/uno")
	node, err := yamlnodefinder.Find(doc, *ptr)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	expected := "UNO"
	if node.Value != expected {
		t.Errorf("Wrong Node is returned\nExpected: %s\nActual: %s", expected, inspectNode(node))
	}
}

func TestFindFromSequenceInsideNestedMappings(t *testing.T) {
	doc := fixture()
	ptr, _ := jsonpointer.Parse("/baz/two/1")
	node, err := yamlnodefinder.Find(doc, *ptr)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	expected := "2"
	if node.Value != expected {
		t.Errorf("Wrong Node is returned\nExpected: %s\nActual: %s", expected, inspectNode(node))
	}
}

func TestFindNonExisting(t *testing.T) {
	doc := fixture()
	ptr, _ := jsonpointer.Parse("/baz/42")
	node, err := yamlnodefinder.Find(doc, *ptr)
	if err == nil {
		t.Error("it must be error")
		return
	}

	if node != nil {
		t.Errorf("Node must be nil\nActual: %s", inspectNode(node))
	}
}

func fixture() *yaml.Node {
	buf := bytes.NewBufferString(inputYAML)
	d := yaml.NewDecoder(buf)
	var n yaml.Node
	d.Decode(&n)
	return &n
}

func inspectNode(n *yaml.Node) string {
	return fmt.Sprintf("%s [kind=%s, tag=%s, line=%d, column=%d]", n.Value, inspectKind(n.Kind), n.Tag, n.Line, n.Column)
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
