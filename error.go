package imgi

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound = errors.New("Not found")
)

func replyError(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"msg": "%s"}`, err.Error())
}
