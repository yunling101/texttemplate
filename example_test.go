package texttemplate

import (
	"fmt"
	"io"
	"log"
	"net/url"
)

func ExampleTemplate() {
	template := "http://{{host}}/?id={{id}}&q={{query}}&z={{z}}"
	t, err := New(template, "{{", "}}")
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"host": "github.com",
		"id":   []byte("example"),
		"query": TagFunc(func(w io.Writer, tag string) (int, error) {
			return w.Write([]byte(url.QueryEscape(tag + "=world")))
		}),
	}

	s, err := t.ExecuteString(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)

	// Output:
	// http://github.com/?id=example&q=query%3Dworld&z=
}

func ExampleTagFunc() {
	template := "foo[baz]bar"
	t, err := NewTemplate(template, "[", "]")
	if err != nil {
		log.Fatalf("unexpected error when parsing template: %s", err)
	}

	bazSlice := [][]byte{[]byte("123"), []byte("456"), []byte("789")}
	m := map[string]interface{}{
		"baz": TagFunc(func(w io.Writer, tag string) (int, error) {
			var nn int
			for _, x := range bazSlice {
				n, err := w.Write(x)
				if err != nil {
					return nn, err
				}
				nn += n
			}
			return nn, nil
		}),
	}

	s, err := t.ExecuteString(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)

	// Output:
	// foo123456789bar
}

func ExampleTemplate_ExecuteFuncString() {
	template := "Hello, [user]! You won [prize]!!! [foobar]"
	t, err := NewTemplate(template, "[", "]")
	if err != nil {
		log.Fatalf("unexpected error when parsing template: %s", err)
	}
	s, err := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "user":
			return w.Write([]byte("Jon"))
		case "prize":
			return w.Write([]byte("$100500"))
		default:
			return w.Write([]byte(fmt.Sprintf("[unknown tag %q]", tag)))
		}
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)

	// Output:
	// Hello, John! You won $100500!!! [unknown tag "foobar"]
}
