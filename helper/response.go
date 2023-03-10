package helper

import (
	"net/http"
	"strings"
)

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}

	if message != "" {
		resp["message"] = message
	}

	switch len(data) {
	case 1:
		resp["data"] = data[0]
	case 2:
		resp["token"] = data[1].(string)
		resp["data"] = data[0]
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	// if strings.Contains(msg, "server") {
	// 	code = http.StatusInternalServerError
	// } else if strings.Contains(msg, "format") {
	// 	code = http.StatusBadRequest
	// } else if strings.Contains(msg, "not found") {
	// 	code = http.StatusNotFound
	// } else if strings.Contains(msg, "validation") {
	// 	code = http.StatusBadRequest
	// }

	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "not found"):
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "duplicate"):
		code = http.StatusConflict
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "validation"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "unmarshal"):
		resp["message"] = "bad request"
		code = http.StatusBadRequest
	}

	return code, resp
}
