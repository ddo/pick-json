package pickjson

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antonholmquist/jason"
)

const (
	JSON_SAMPLE_BENCHMARK = `
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
		benchmark := PickString(strings.NewReader(JSON_SAMPLE_BENCHMARK), "benchmark", 1)

		if len(benchmark) == 0 {
			fmt.Println(benchmark)
		}
	}
}

// jason get from the the bottom up
func BenchmarkJasonSimple(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(JSON_SAMPLE_BENCHMARK))

		benchmark, err := j.GetString("benchmark")

		if err != nil {
			fmt.Println(err, benchmark)
		}
	}
}

func BenchmarkNestedPickString(B *testing.B) {
	for n := 0; n < B.N; n++ {
		nestedBenchmark := PickString(strings.NewReader(JSON_SAMPLE_BENCHMARK), "nested_benchmark", 1)

		if len(nestedBenchmark) == 0 {
			fmt.Println(nestedBenchmark)
		}
	}
}

func BenchmarkNestedJasonString(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(JSON_SAMPLE_BENCHMARK))

		nestedBenchmark, err := j.GetString("nested", "nested_benchmark")

		if err != nil {
			fmt.Println(err, nestedBenchmark)
		}
	}
}

func BenchmarkArrayPickString(B *testing.B) {
	for n := 0; n < B.N; n++ {
		id := PickString(strings.NewReader(JSON_SAMPLE_BENCHMARK), "ID", 0)

		if len(id) == 0 {
			fmt.Println(id)
		}
	}
}

func BenchmarkArrayJasonString(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(JSON_SAMPLE_BENCHMARK))

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

func BenchmarkNestedArrayPickString(B *testing.B) {
	for n := 0; n < B.N; n++ {
		id := PickString(strings.NewReader(JSON_SAMPLE_BENCHMARK), "link", 0)

		if len(id) == 0 {
			fmt.Println(id)
		}
	}
}

func BenchmarkNestedArrayJasonString(B *testing.B) {
	for n := 0; n < B.N; n++ {
		j, _ := jason.NewObjectFromReader(strings.NewReader(JSON_SAMPLE_BENCHMARK))

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
