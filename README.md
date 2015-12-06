pick-json [![Build Status][semaphoreci-img]][semaphoreci-url] [![Doc][godoc-img]][godoc-url]
==========
> pick stuff from json - FAST

[godoc-img]: https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat-square
[godoc-url]: https://godoc.org/github.com/ddo/pick-json

[semaphoreci-img]: https://semaphoreci.com/api/v1/projects/4d90ed7b-d2e8-45ed-bbb1-a077b1fefeeb/625358/badge.svg
[semaphoreci-url]: https://semaphoreci.com/ddo/pick-json

## Feature

* fast.
* simple.
* lightweight. just stdlib ``"encoding/json"`` and ~2.0 KB
* incomplete json still work well.
* no need to define parent key for nested target.

## Fast? How?

* process json as stream.
* stop when found.
* and you can also set the limit.

## Use cases

* we just need a part of the json.
* in case you need to parse all the json, please use the ``"encoding/json"`` ``Unmarshal`` or ``Decoder``

## Examples

```json
JSON_EXAMPLE := `{   
    "benchmark": "benchmark text 1",
    "menu": {
        "header": "SVG Viewer",
        "image": { 
            "src": "Images/Sun.png",
            "name": "sun1",
            "hOffset": 250,
            "vOffset": 250,
            "alignment": "center",
            "hidden": true
        }
    },
    "benchmark": "benchmark text 2",
}`
```

* pick string

```go
benchmarks := PickString(strings.NewReader(JSON_EXAMPLE), "benchmark", 0)
// [benchmark text 1 benchmark text 2]
```

* pick string just the 1st one

```go
benchmarks := PickString(strings.NewReader(JSON_EXAMPLE), "benchmark", 1)
// [benchmark text 1]
```

* pick bool

```go
hidden := PickBool(strings.NewReader(JSON_EXAMPLE), "hidden", 0)
// [true]
```

* pick object

```go
type Image struct {
    Src       string `json:"src"`
    Name      string `json:"name"`
    HOffset   int    `json:"hOffset"`
    VOffset   int    `json:"vOffset"`
    Alignment string `json:"alignment"`
}

var image Image

PickObject(strings.NewReader(JSON_EXAMPLE), "image", &image)
// {Images/Sun.png sun1 250 250 center}
```
