package helpers

import (
	"strings"

	"github.com/jinzhu/inflection"
)

type Name struct {
	Name string
}

func NewName(name string) *Name {
	return &Name{Name: name}
}

func (n *Name) ToString() string {
	return n.Name
}

func (n *Name) FirstLetter() *Name {
	return NewName(n.Name[0:1])
}

func (n *Name) Lower() *Name {
	return NewName(strings.ToLower(n.Name))
}

func (n *Name) Camel() *Name {
	return NewName(Camelize(n.Name))
}

func (n *Name) Plural() *Name {
	return NewName(inflection.Plural(n.Name))
}

func Camelize(snaked string) (camelized string) {
	for _, token := range strings.Split(snaked, "_") {
		camelized += strings.Title(token)
	}
	return
}
