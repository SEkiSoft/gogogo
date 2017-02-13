// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net"
	"net/http"
	"strings"

	l4g "github.com/alecthomas/log4go"
	"github.com/sekisoft/gogogo/model"
)

var allowedMethods []string = []string{
	"POST",
	"GET",
}

type Session struct {
	RequestID string
	IpAddress string
	Path      string
	Err       *model.AppError
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

func GetToken(id string) *model.Token {
	var token *model.Token

	if tokenResult := <-Srv.Store.Token().Get(id); tokenResult.Err != nil {
		l4g.Error("Invalid token id: %s, %s", id, tokenResult.Err.Message)
	} else {
		token = tokenResult.Data.(*model.Token)
	}

	return token
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
	s.RequestID = model.NewID()
	s.IpAddress = GetIpAddress(r)
	s.RootUrl = GetProtocol(r) + "://" + r.Host
	s.Err = nil
	s.Token = nil
	s.Path = r.URL.Path

	w.Header().Set(model.HEADER_REQUEST_ID, s.RequestID)
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
		http.Error(w, s.Err.Error(), s.Err.StatusCode)
	}
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	err := model.NewAppError("Handle404", "404 not found", http.StatusBadGateway)
	err.StatusCode = http.StatusNotFound
	l4g.Error("%v: code=404 ip=%v", r.URL.Path, GetIpAddress(r))

	http.Error(w, err.Error(), err.StatusCode)
}

func (s *Session) SetInvalidParam(location string, name string) {
	s.Err = NewInvalidParamError(location, name)
}

func (s *Session) CheckPlayerRequired() {
	if len(s.Token.PlayerID) == 0 {
		s.Err = model.NewAppError("CheckPlayerRequired", "Player invalid", http.StatusBadRequest)
		s.Err.StatusCode = http.StatusUnauthorized
	}
}

func (s *Session) CheckAdminRequired() {
	if len(s.Token.PlayerID) == 0 {
		s.Err = model.NewAppError("CheckAdminRequired", "Player invalid", http.StatusBadRequest)
		s.Err.StatusCode = http.StatusUnauthorized
	} else if !s.IsAdmin() {
		s.Err = model.NewAppError("CheckAdminRequired", "Admin invalid", http.StatusBadRequest)
		s.Err.StatusCode = http.StatusUnauthorized
	}
}

func (s *Session) IsAdmin() bool {
	if player, err := GetPlayer(s.Token.PlayerID); err == nil {
		return player.IsAdmin
	}
	return false
}

func NewInvalidParamError(location string, name string) *model.AppError {
	err := model.NewAppError(location, "Invalid parameters error", http.StatusBadRequest)
	err.StatusCode = http.StatusBadRequest
	return err
}
