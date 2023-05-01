package helpers

import (
	"testing"

	"github.com/jinzhu/inflection"
)

func TestNewName(t *testing.T) {
	name := NewName("Scaffold")
	if name.Name != "Scaffold" {
		t.Errorf("expected Name to be 'Scaffold', got %q", name.Name)
	}
}

func TestToString(t *testing.T) {
	name := NewName("Scaffold")
	if name.ToString() != "Scaffold" {
		t.Errorf("expected ToString() to return 'Scaffold', got %q", name.ToString())
	}
}

func TestFirstLetter(t *testing.T) {
	name := NewName("Scaffold")
	if name.FirstLetter().ToString() != "S" {
		t.Errorf("expected FirstLetter() to return 'S', got %q", name.FirstLetter().ToString())
	}
}

func TestLower(t *testing.T) {
	name := NewName("Scaffold")
	if name.Lower().ToString() != "scaffold" {
		t.Errorf("expected Lower() to return 'scaffold', got %q", name.Lower().ToString())
	}
}

func TestCamel(t *testing.T) {
	name := NewName("scaffold_framework")
	if name.Camel().ToString() != "ScaffoldFramework" {
		t.Errorf("expected Camel() to return 'ScaffoldFramework', got %q", name.Camel().ToString())
	}
}

func TestPlural(t *testing.T) {
	name := NewName("scaffold")
	if name.Plural().ToString() != inflection.Plural("scaffolds") {
		t.Errorf("expected Plural() to return %q, got %q", inflection.Plural("scaffold"), name.Plural().ToString())
	}
}
