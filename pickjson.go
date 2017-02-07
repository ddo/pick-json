package pickjson

import (
	"encoding/json"
	"io"
)

func processToken(reader io.Reader, key string, hook func(*json.Decoder) bool) {
	decoder := json.NewDecoder(reader)

	// if empty key -> hook the whole decoder
	// only for #PickObject root object
	if key == "" {
		hook(decoder)
		return
	}

	isKey := true

	for {
		token, err := decoder.Token()

		// return when done
		if err == io.EOF {
			break
		}

		// stop on error
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

// PickString pick String by key, return empty array if there is no matching key or error
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

// PickBool pick Bool by key, return empty array if there is no matching key or error
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

// PickNumber pick Float64 by key, return empty array if there is no matching key or error
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

// PickObject pick struct by key
// there is no limit for PickObject yet :(
// TODO: add limit
func PickObject(reader io.Reader, key string, object interface{}) (err error) {
	processToken(reader, key, func(decoder *json.Decoder) bool {
		err = decoder.Decode(object)
		return true
	})

	return
}
