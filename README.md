# generics
The missing pieces...

## usage

### First

[embedmd]:# (example_test.go /func example.*/ / end example/)
```go
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
```
