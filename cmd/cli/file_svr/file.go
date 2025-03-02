package file_svr

import (
	"net/http"
)

type file struct {
	dir string
}

func (f file) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	// TODO implement me
	panic("implement me")
}
