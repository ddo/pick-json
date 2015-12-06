package pickjson

import (
	"encoding/json"
	"io"
)

func processToken(reader io.Reader, key string, hook func(*json.Decoder) bool) {
	decoder := json.NewDecoder(reader)

	isKey := true

	for {
		token, err := decoder.Token()

		// return when done
		if err == io.EOF {
			break
		}

		// skip the error token
		if err != nil {
			break
		}

		// start process the tokens
		switch token.(type) {
		case json.Delim: // { } [ ]
			isKey = true

		case string: // key or value
			if isKey && key == token {
				stop := hook(decoder)

				if stop {
					return
				}
			}
			isKey = !isKey

		default:
			isKey = !isKey
		}
	}
	return
}

func PickString(reader io.Reader, key string, limit int) (res []string) {
	processToken(reader, key, func(decoder *json.Decoder) bool {
		// matched token
		token, err := decoder.Token()

		// skip the error token
		if err != nil {
			return false
		}

		if tokenStr, ok := token.(string); ok {
			res = append(res, tokenStr)

			if limit > 0 && len(res) >= limit {
				return true
			}
		}

		return false
	})
	return
}

func PickBool(reader io.Reader, key string, limit int) (res []bool) {
	processToken(reader, key, func(decoder *json.Decoder) bool {
		// matched token
		token, err := decoder.Token()

		// skip the error token
		if err != nil {
			return false
		}

		if tokenStr, ok := token.(bool); ok {
			res = append(res, tokenStr)

			if limit > 0 && len(res) >= limit {
				return true
			}
		}

		return false
	})
	return
}

func PickNumber(reader io.Reader, key string, limit int) (res []float64) {
	processToken(reader, key, func(decoder *json.Decoder) bool {
		// matched token
		token, err := decoder.Token()

		// skip the error token
		if err != nil {
			return false
		}

		if tokenStr, ok := token.(float64); ok {
			res = append(res, tokenStr)

			if limit > 0 && len(res) >= limit {
				return true
			}
		}

		return false
	})
	return
}

// there is no limit for PickObject yet :(
func PickObject(reader io.Reader, key string, object interface{}) (err error) {
	processToken(reader, key, func(decoder *json.Decoder) bool {
		err = decoder.Decode(object)
		return true
	})

	return
}
