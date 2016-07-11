// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net"
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/davidlu1997/gogogo/model"
)

var allowedMethods []string = []string{
	"POST",
	"GET",
	"PUT",
}

type Session struct {
	RequestId string
	IpAddress string
	Path      string
	Err       *model.Error
	PlayerId  string
	RootUrl   string
}

type handler struct {
	handleFunc     func(*Session, http.ResponseWriter, *http.Request)
	requiredPlayer bool
	requiredGame   bool
}

func ApiHandler(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, false}
}

func ApiPlayerRequired(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, true, false}
}

func ApiGameRequired(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, true}
}

func GetProtocol(r *http.Request) string {
	if r.Header.Get(model.HEADER_FORWARDED_PROTO) == "https" {
		return "https"
	} else {
		return "http"
	}
}

func GetIpAddress(r *http.Request) string {
	address := r.Header.Get(model.HEADER_FORWARDED)

	if len(address) == 0 {
		address = r.Header.Get(model.HEADER_REAL_IP)
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return address
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func RenderWebError(err *model.Error, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/404.html", http.StatusTemporaryRedirect)
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	err := model.NewLocError("Handle404", "404 not found", nil, "")
	err.StatusCode = http.StatusNotFound
	l4g.Error("%v: code=404 ip=%v", r.URL.Path, GetIpAddress(r))

	RenderWebError(err, w, r)
}
