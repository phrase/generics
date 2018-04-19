package generics

import (
	"bytes"
	"log"
	"testing"
)

func exampleFirst(l *log.Logger) {
	type Thing struct {
		Name string
	}

	list := []*Thing{
		{Name: "Parrot"},
		{Name: "PhraseApp"},
	}

	first := First(list).(*Thing)
	l.Printf("%s", first.Name)
} // end example

func TestRunExample(t *testing.T) {
	buf := &bytes.Buffer{}
	l := log.New(buf, "", 0)
	exampleFirst(l)
	name := buf.String()
	if name != "Parrot\n" {
		t.Errorf("name was %q", name)
	}
}
