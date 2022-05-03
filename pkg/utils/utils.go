package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
 * Convert body from JSON
 *
 */
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// handle error
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
