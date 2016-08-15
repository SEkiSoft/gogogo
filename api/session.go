// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net"
	"net/http"
	"strings"

	l4g "github.com/alecthomas/log4go"
	"github.com/SEkiSoft/gogogo/model"
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
	GameId    string
	RootUrl   string
}

type handler struct {
	handleFunc     func(*Session, http.ResponseWriter, *http.Request)
	requiredPlayer bool
	requiredGame   bool
	requiredAdmin  bool
	isApi          bool
}

func ApiHandler(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, false, false, true}
}

func ApiPlayerRequired(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, true, false, false, true}
}

func ApiGameRequired(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, true, false, true}
}

func ApiAdminRequired(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, false, true, true}
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

	s := &Session{}
	s.RequestId = model.NewId()
	s.IpAddress = GetIpAddress(r)
	s.RootUrl = GetProtocol(r) + "://" + r.Host
	s.Err = nil

	if h.isApi {
		w.Header().Set(model.HEADER_REQUEST_ID, s.RequestId)
		w.Header().Set("Content-Type", "application/json")
	}

	// TODO Get token from header, cookie, and query string
	// Authenicate user based on token
	// Using PlayerRequired, GameRequired, AdminRequired, etc.
	token := ""

	auth := r.Header.Get(model.HEADER_AUTH)
	if len(auth) > 6 && strings.ToUpper(auth[0:6]) == model.HEADER_BEARER {
		token = auth[7:]
	} else {
		return
	}

	if h.isApi {
		s.Path = r.URL.Path
	} else {
		splitURL := strings.Split(r.URL.Path, "/")
		s.Path = "/" + strings.Join(splitURL[2:], "/")
	}

	if len(token) > 0 {
		// TODO session store
	}

	if s.Err == nil && h.requiredPlayer {
		s.CheckPlayerRequired()
	}
	if s.Err == nil && h.requiredGame {
		s.CheckGameRequired()
	}
	if s.Err == nil && h.requiredAdmin {
		s.CheckAdminRequired()
	}
	if s.Err == nil {
		h.handleFunc(s, w, r)
	}

	if s.Err != nil {
		if h.isApi {
			w.WriteHeader(s.Err.StatusCode)
			w.Write([]byte(s.Err.ToJson()))
		} else {
			RenderWebError(s.Err, w, r)
		}
	}
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

func (s *Session) SetInvalidParam(location string, name string) {
	s.Err = NewInvalidParamError(location, name)
}

func (s *Session) CheckPlayerRequired() {
	if len(s.PlayerId) == 0 {
		s.Err = model.NewLocError("CheckPlayerRequired", "Player invalid", nil, "")
		s.Err.StatusCode = http.StatusUnauthorized
	}
}

func (s *Session) CheckGameRequired() {
	if len(s.GameId) == 0 {
		s.Err = model.NewLocError("CheckGameRequired", "Game invalid", nil, "")
		s.Err.StatusCode = http.StatusUnauthorized
	}
}

func (s *Session) CheckAdminRequired() {
	if len(s.PlayerId) == 0 {
		s.Err = model.NewLocError("CheckAdminRequired", "Player invalid", nil, "")
		s.Err.StatusCode = http.StatusUnauthorized
	} else if !s.IsAdmin() {
		s.Err = model.NewLocError("CheckAdminRequired", "Admin invalid", nil, "")
		s.Err.StatusCode = http.StatusUnauthorized
	}
}

func (s *Session) IsAdmin() bool {
	if player, err := GetPlayer(s.PlayerId); err == nil {
		return player.IsAdmin
	}
	return false
}

func NewInvalidParamError(location string, name string) *model.Error {
	err := model.NewLocError(location, "Invalid parameters error", map[string]interface{}{"Name": name}, "")
	err.StatusCode = http.StatusBadRequest
	return err
}
