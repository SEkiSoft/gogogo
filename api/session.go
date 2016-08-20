// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net"
	"net/http"
	"strings"

	"github.com/SEkiSoft/gogogo/model"
	l4g "github.com/alecthomas/log4go"
)

var allowedMethods []string = []string{
	"POST",
	"GET",
}

type Session struct {
	RequestId string
	IpAddress string
	Path      string
	Err       *model.Error
	RootUrl   string
	Token     *model.Token
}

type handler struct {
	handleFunc     func(*Session, http.ResponseWriter, *http.Request)
	requiredPlayer bool
	requiredAdmin  bool
}

func ApiHandler(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, false}
}

func ApiPlayerRequired(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, true, false}
}

func ApiAdminRequired(h func(*Session, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h, false, true}
}

func GetProtocol(r *http.Request) string {
	if r.Header.Get(model.HEADER_FORWARDED_PROTO) == "https" {
		return "https"
	}

	return "http"
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
	s.Token = nil
	s.Path = r.URL.Path

	w.Header().Set(model.HEADER_REQUEST_ID, s.RequestId)
	w.Header().Set("Content-Type", "application/json")

	token := ""

	auth := r.Header.Get(model.HEADER_AUTH)
	if len(auth) > 4 && strings.ToUpper(auth[0:4]) == model.HEADER_BEAR {
		token = auth[5:]
	}

	if len(token) > 0 {
		if result := <-Srv.Store.Token().Get(token); result.Err != nil {
			s.Err = result.Err
		} else {
			s.Token = result.Data.(*model.Token)
		}
	}

	if s.Err == nil && h.requiredPlayer {
		s.CheckPlayerRequired()
	}

	if s.Err == nil && h.requiredAdmin {
		s.CheckAdminRequired()
	}
	if s.Err == nil {
		h.handleFunc(s, w, r)
	}

	if s.Err != nil {
		w.WriteHeader(s.Err.StatusCode)
		w.Write([]byte(s.Err.ToJson()))
	}
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	err := model.NewLocError("Handle404", "404 not found", nil, "")
	err.StatusCode = http.StatusNotFound
	l4g.Error("%v: code=404 ip=%v", r.URL.Path, GetIpAddress(r))

	w.WriteHeader(err.StatusCode)
	w.Write([]byte(err.ToJson()))
}

func (s *Session) SetInvalidParam(location string, name string) {
	s.Err = NewInvalidParamError(location, name)
}

func (s *Session) CheckPlayerRequired() {
	if len(s.Token.PlayerId) == 0 {
		s.Err = model.NewLocError("CheckPlayerRequired", "Player invalid", nil, "")
		s.Err.StatusCode = http.StatusUnauthorized
	}
}

func (s *Session) CheckAdminRequired() {
	if len(s.Token.PlayerId) == 0 {
		s.Err = model.NewLocError("CheckAdminRequired", "Player invalid", nil, "")
		s.Err.StatusCode = http.StatusUnauthorized
	} else if !s.IsAdmin() {
		s.Err = model.NewLocError("CheckAdminRequired", "Admin invalid", nil, "")
		s.Err.StatusCode = http.StatusUnauthorized
	}
}

func (s *Session) IsAdmin() bool {
	if player, err := GetPlayer(s.Token.PlayerId); err == nil {
		return player.IsAdmin
	}
	return false
}

func NewInvalidParamError(location string, name string) *model.Error {
	err := model.NewLocError(location, "Invalid parameters error", map[string]interface{}{"Name": name}, "")
	err.StatusCode = http.StatusBadRequest
	return err
}
