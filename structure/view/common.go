package view

import "net/http"

func InternalServerError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "terdapat kesalahan pada server",
		"status":  false,
		"data":    nil,
	}
}
