package pickjson

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antonholmquist/jason"
)

const (
	jsonSampleBenchmark = `
{
    "benchmark": "benchmark text 1",
    "nested": {
        "nested_benchmark": "benchmark text 2"
    },
    "data": [{
        "ID": "X999_Y999",
        "from": {
            "name": "Tom Brady",
            "id": "X12"
        },
        "message": "Looking forward to 2010!",
        "actions": [{
            "name": "Comment",
            "link": "http://www.facebook.com/X999/posts/1"
        }, {
            "name": "Like",
            "link": "http://www.facebook.com/X999/posts/2"
        }],
        "type": "status",
        "created_time": "2010-08-02T21:27:44+0000",
        "updated_time": "2010-08-02T21:27:44+0000"
    }, {
        "ID": "X998_Y998",
        "from": {
            "name": "Peyton Manning",
            "id": "X18"
        },
        "message": "Where's my contract?",
        "actions": [{
            "name": "Comment",
            "link": "http://www.facebook.com/X998/posts/3"
        }, {
            "name": "Like",
            "link": "http://www.facebook.com/X998/posts/4"
        }],
        "type": "status",
        "created_time": "2010-08-02T21:27:44+0000",
        "updated_time": "2010-08-02T21:27:44+0000"
    }]
}
`
)

func BenchmarkPickSimple(B *testing.B) {
	for n := 0; n < B.N; n++ {
		benchmark := PickString(strings.NewReader(jsonSampleBenchmark), "benchmark", 1)
		if len(benchmark) == 0 {
			fmt.Println(benchmark)
		}
	}
}

// jason get from the the bottom up
func BenchmarkJasonSimple(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(jsonSampleBenchmark))

		benchmark, err := j.GetString("benchmark")
		if err != nil {
			fmt.Println(err, benchmark)
		}
	}
}

func BenchmarkPickNested(B *testing.B) {
	for n := 0; n < B.N; n++ {
		nestedBenchmark := PickString(strings.NewReader(jsonSampleBenchmark), "nested_benchmark", 1)
		if len(nestedBenchmark) == 0 {
			fmt.Println(nestedBenchmark)
		}
	}
}

func BenchmarkJasonNested(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(jsonSampleBenchmark))

		nestedBenchmark, err := j.GetString("nested", "nested_benchmark")
		if err != nil {
			fmt.Println(err, nestedBenchmark)
		}
	}
}

func BenchmarkPickArray(B *testing.B) {
	for n := 0; n < B.N; n++ {
		id := PickString(strings.NewReader(jsonSampleBenchmark), "ID", 0)
		if len(id) == 0 {
			fmt.Println(id)
		}
	}
}

func BenchmarkJasonArray(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(jsonSampleBenchmark))

		dataArray, err := j.GetObjectArray("data")
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, dataItem := range dataArray {
			id, err := dataItem.GetString("ID")
			if err != nil {
				fmt.Println(err, id)
				return
			}
		}
	}
}

func BenchmarkPickNestedArray(B *testing.B) {
	for n := 0; n < B.N; n++ {
		id := PickString(strings.NewReader(jsonSampleBenchmark), "link", 0)
		if len(id) == 0 {
			fmt.Println(id)
		}
	}
}

func BenchmarkJasonNestedArray(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(jsonSampleBenchmark))

		dataArray, err := j.GetObjectArray("data")
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, dataItem := range dataArray {
			actionArray, err := dataItem.GetObjectArray("actions")
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, actionItem := range actionArray {
				link, err := actionItem.GetString("link")
				if err != nil {
					fmt.Println(err, link)
					return
				}
			}
		}
	}
}
