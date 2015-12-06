package pickjson

import (
	"fmt"
	"strings"
)

const (
	JSON_EXAMPLE = `{   
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
)

// pick string
func ExamplePickString() {
	benchmarks := PickString(strings.NewReader(JSON_EXAMPLE), "benchmark", 0)
	fmt.Println(benchmarks)
	// Output: [benchmark text 1 benchmark text 2]
}

// pick string just the 1st one
func ExamplePickStringLimit() {
	benchmarks := PickString(strings.NewReader(JSON_EXAMPLE), "benchmark", 1)
	fmt.Println(benchmarks)
	// Output: [benchmark text 1]
}

// pick bool
func ExamplePickBool() {
	hidden := PickBool(strings.NewReader(JSON_EXAMPLE), "hidden", 0)
	fmt.Println(hidden)
	// Output: [true]
}

// pick object
func ExamplePickObject() {
	var image Image

	PickObject(strings.NewReader(JSON_EXAMPLE), "image", &image)

	fmt.Println(image)
	// Output: {Images/Sun.png sun1 250 250 center}
}
