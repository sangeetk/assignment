package handlers

import (
	"fmt"
	"net/http"

	"assignment/matrix"
)

func Echo(w http.ResponseWriter, r *http.Request) {

	m, err := matrix.New(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, m)
}
