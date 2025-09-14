package utils

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			fmt.Println("Error parsing body:", err)
		}
	}
}
