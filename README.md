[![GoDoc](https://godoc.org/github.com/bserdar/gojsonpath?status.svg)](https://godoc.org/github.com/bserdar/gojsonpath)
[![Go Report Card](https://goreportcard.com/badge/github.com/bserdar/gojsonpath)](https://goreportcard.com/report/github.com/bserdar/gojsonpath)
[![Build Status](https://github.com/bserdar/gojsonpath/actions/workflows/CI.yml/badge.svg?branch=main)](https://github.com/bserdar/gojsonpath/actions/workflows/CI.yml)

# JSONPath

This is a Go implementation of JSONPath introduced in:

https://goessner.net/articles/JsonPath/

with the following exceptions:

  * If a path component is a it should be a valid identifier or
     number. If that is not the case, it should be a string literal,
     that is, quoted.

     For example: `$.dash-key` is invalid. You must quote it:
     `$."dash-key"` or `$.'dash-key'`.

   * Function and method calls must have a possibly empty argument
     list.

     For example: `length(@)` and `@.length()` are valid, but
     `@.length` will look for a `length` key under the current node.

   * Decimals cannot start with a period, that is `.5` is
     invalid. Use `0.5` instead. Decimal starting with a period
     confuses the lexer.

## Using JSONPath

First, compile the path:

``` 
path, err:= gojsonpath.Parse("$.store.book[*].author")
```

Then you can use:

```
results, err:= gojsonpath.Find(doc,path)
```

or you can capture the full-path for the nodes that match:

``` go
result:=make([]any,0)
err := gojsonpath.Search(doc,path,func(elem []gojsonpath.Element) {
  result=append(result,elem[len(elem)-1].Node)
})
```

## Document model

This implementation of JSONPath works with a flexible document
model. To use the familiar `map[string]any` implementation of JSON,
use:

``` go
var jsonDoc any
json.Unmarshal(jsonBytes,&jsonDoc)
doc := gojsonpath.MapModel{Doc:jsonDoc}
```

The `DocModel` interface provides a view of the underlying model, so
the JSON document does not have to be a `map[string]any`. Any
hierarchical document model supporting key-value, array, and value
nodes can be used.

