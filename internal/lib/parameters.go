package lib

import (
	"net/http"
)

// ValuesURL retorna los valores pasados por la url
// ejemplo server/products/?key=value&key2=value2
func ValuesURL(r *http.Request, key string) string {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return ""
	}

	return keys[0]
}
