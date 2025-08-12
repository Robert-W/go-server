package system

import "net/http"

func Healthcheck(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
}
