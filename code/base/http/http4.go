package main
import (
	"net/http"
	"fmt"
)
type SingleMiddle struct {
	handler   http.Handler
	allowHost []string
}
type myHandler struct{}
func main() {
	var allowHosts []string
	allowHosts = append(allowHosts, "localhost:8080")
	MyHandler := &myHandler{}
	singleHost := NewSingle(MyHandler, allowHosts)
	http.ListenAndServe(":8080", singleHost)
}
// middleware
func NewSingle(handler http.Handler, allowHost []string) (*SingleMiddle) {
	return &SingleMiddle{handler, allowHost}
}
// middleware
func (s *SingleMiddle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	isIn := false
	fmt.Println(s.allowHost)
	fmt.Println(host)
	for _, allowH := range s.allowHost {
		if allowH == host {
			isIn = true
		}
	}
	if isIn {
		s.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}
// custom handler
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("helle golang"))
}