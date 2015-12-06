package pickjson

import (
	"strings"
	"testing"

	"gopkg.in/ddo/request.v1"
)

const (
	JSON_SAMPLE = `
{   
    "menu": {
        "header": "SVG Viewer",
        "image": { 
            "src": "Images/Sun.png",
            "name": "sun1",
            "hOffset": 250,
            "vOffset": 250,
            "alignment": "center"
        },
        "items": [{
                "id": 1
            }, {
                "id": 2,
                "label": "Open New"
            },
            null, {
                "id": 3,
                "label": "Zoom In"
            }, {
                "id": 4,
                "label": "Zoom Out"
            }, {
                "id": 5,
                "label": "Original View"
            },
            null, {
                "id": 6
            }, {
                "id": 7
            }, {
                "id": 8
            },
            null, {
                "id": 9,
                "label": "Find..."
            }, {
                "id": 10,
                "label": "Find Again"
            }, {
                "id": 11
            }, {
                "id": "CopyAgain",
                "label": "Copy Again",
                "ok": false
            }, {
                "id": 12,
                "label": "Copy SVG",
                "ok": true
            }, {
                "id": 13,
                "label": "View SVG",
                "ok": false
            }, {
                "id": 14,
                "label": "View Source",
                "ok": false
            }, {
                "id": 15.00,
                "label": "Save As",
                "ok": false
            },
            null, {
                "id": 16
            }, {
                "id": "About",
                "label": "About Adobe CVG Viewer..."
            }, {
                "id": "About 2",
                "label": 1,
                "ok": "false"
            }
        ]
    },
    "image": { 
        "src": "Images/Sun2.png",
        "name": "sun2",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    }
}
`
	// still works
	JSON_SAMPLE_INCOMPLETE = `
{
    "label": "from incomplete sample"
`

	// still works
	JSON_SAMPLE_INCOMPLETE_OBJECT = `
{
    "image": { 
        "src": "Images/Sun2.png",
        "name": "sun2",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    },
    asdasd asd wqeqw eqe qsad
`

	JSON_SAMPLE_ARRAY = `
[
    {
        "id": 1
    }, {
        "id": 2,
        "label": "Open New"
    },
    null, {
        "id": 3,
        "label": "Zoom In"
    }, {
        "id": 4,
        "label": "Zoom Out"
    }, {
        "id": 5,
        "label": "Original View"
    },
    null, {
        "id": 6
    }, {
        "id": 7
    }, {
        "id": 8
    },
    null, {
        "id": 9,
        "label": "Find..."
    }, {
        "id": 10,
        "label": "Find Again"
    }, {
        "id": 11
    }, {
        "id": "CopyAgain",
        "label": "Copy Again",
        "ok": false
    }, {
        "id": 12,
        "label": "Copy SVG",
        "ok": true
    }, {
        "id": 13,
        "label": "View SVG",
        "ok": false
    }, {
        "id": 14,
        "label": "View Source",
        "ok": false
    }, {
        "id": 15.00,
        "label": "Save As",
        "ok": false
    },
    null, {
        "id": 16
    }, {
        "id": "About",
        "label": "About Adobe CVG Viewer..."
    }
]
`

	JSON_SAMPLE_INVALID_KEY = `
{
    label: "from incomplete sample"
}
`
)

type Image struct {
	Src       string `json:"src"`
	Name      string `json:"name"`
	HOffset   int    `json:"hOffset"`
	VOffset   int    `json:"vOffset"`
	Alignment string `json:"alignment"`
}

type Headers struct {
	Accept         string `json:"Accept"`
	AcceptEncoding string `json:"Accept-Encoding"`
	AcceptLanguage string `json:"Accept-Language"`
	Host           string `json:"Host"`
	UserAgent      string `json:"User-Agent"`
}

func TestPickString(t *testing.T) {
	res := PickString(strings.NewReader(JSON_SAMPLE), "label", 0)

	if len(res) != 12 {
		t.Error()
	}
}

func TestPickStringLimit(t *testing.T) {
	res := PickString(strings.NewReader(JSON_SAMPLE), "label", 5)

	if len(res) != 5 {
		t.Error()
	}
}

func TestPickStringIncomplete(t *testing.T) {
	res := PickString(strings.NewReader(JSON_SAMPLE_INCOMPLETE), "label", 0)

	if len(res) != 1 {
		t.Error()
	}
}

func TestPickStringArray(t *testing.T) {
	res := PickString(strings.NewReader(JSON_SAMPLE_ARRAY), "label", 0)

	if len(res) != 12 {
		t.Error()
	}
}

func TestPickBool(t *testing.T) {
	res := PickBool(strings.NewReader(JSON_SAMPLE), "ok", -1)

	if len(res) != 5 {
		t.Error()
	}
}

func TestPickBoolLimit(t *testing.T) {
	res := PickBool(strings.NewReader(JSON_SAMPLE), "ok", 2)

	if len(res) != 2 {
		t.Error()
	}
}

func TestPickNumber(t *testing.T) {
	res := PickNumber(strings.NewReader(JSON_SAMPLE), "id", 0)

	if len(res) != 16 {
		t.Error()
	}
}

func TestPickNumberLimit(t *testing.T) {
	res := PickNumber(strings.NewReader(JSON_SAMPLE), "id", 10)

	if len(res) != 10 {
		t.Error()
	}
}

func TestPickObject(t *testing.T) {
	var image Image

	err := PickObject(strings.NewReader(JSON_SAMPLE), "image", &image)

	if err != nil {
		t.Error()
	}

	// empty struct
	if image == (Image{}) {
		t.Error()
	}
}

func TestPickObjectIncomplete(t *testing.T) {
	var image Image

	err := PickObject(strings.NewReader(JSON_SAMPLE_INCOMPLETE_OBJECT), "image", &image)

	if err != nil {
		t.Error()
	}

	// empty struct
	if image == (Image{}) {
		t.Error()
	}
}

/////////////////////////////// FAIL TEST ////////////////////////////////
func TestEmptyPickString(t *testing.T) {
	res := PickString(strings.NewReader(""), "label", 0)

	if len(res) != 0 {
		t.Error()
	}
}

func TestEmptyPickObject(t *testing.T) {
	var image Image

	err := PickObject(strings.NewReader(""), "image", &image)

	if err != nil {
		t.Error()
	}

	// empty struct
	if image != (Image{}) {
		t.Error()
	}
}

func TestInvalidKeyPickString(t *testing.T) {
	res := PickString(strings.NewReader(JSON_SAMPLE_INVALID_KEY), "label", 0)

	if len(res) != 0 {
		t.Error()
	}
}

/////////////////////////////// FAIL TEST ////////////////////////////////

/////////////////////////////// HTTP ////////////////////////////////
func TestPickHttpString(t *testing.T) {
	resp, err := request.New().Request(&request.Option{
		Url: "https://httpbin.org/get",
	})

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	res := PickString(resp.Body, "Host", 0)

	if len(res) == 0 {
		t.Error()
	}

	if res[0] != "httpbin.org" {
		t.Error()
	}
}

func TestPickHttpBool(t *testing.T) {
	resp, err := request.New().Request(&request.Option{
		Url: "http://httpbin.org/gzip",
	})

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	res := PickBool(resp.Body, "gzipped", 0)

	if len(res) == 0 {
		t.Error()
	}

	if !res[0] {
		t.Error()
	}
}

func TestPickHttpObject(t *testing.T) {
	resp, err := request.New().Request(&request.Option{
		Url: "http://httpbin.org/get",
	})

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var headers Headers

	err = PickObject(resp.Body, "headers", &headers)

	if err != nil {
		t.Error()
	}

	// empty struct
	if headers == (Headers{}) {
		t.Error()
	}
}

/////////////////////////////// HTTP ////////////////////////////////
