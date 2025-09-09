package user

import "net/http"

// Usage:
// subrouter.Handle("/users", testMiddleware(userHandler.list)).Methods("GET")
func testMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Testing is working")

		next(w, r)
	})
}
