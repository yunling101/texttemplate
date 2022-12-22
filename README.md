texttemplate
============

A simple and efficient Go text template engine

Usage
=====

```go
    go get github.com/yunling101/texttemplate
```

Example
=====

```go
    template := "http://{{host}}/?q={{query}}&z={{z}}"
    
    t := texttemplate.New(template, "{{", "}}")
    s := t.ExecuteString(map[string]interface{}{
        "host":  "github.com",
        "query": url.QueryEscape("texttemplate"),
        "z":   "00111",
    })
    fmt.Printf("%s", s)

    // Output:
    // http://github.com/?q=texttemplate&z=00111
```