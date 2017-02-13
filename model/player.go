// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
	"regexp"
	"strings"

	"net/http"

	"golang.org/x/crypto/bcrypt"
)

const (
	DEFAULT_LOCALE      = "en"
	MIN_PASSWORD_LENGTH = 5
	MAX_PASSWORD_LENGTH = 64
	MIN_USERNAME_LENGTH = 4
	MAX_USERNAME_LENGTH = 24
)

type Player struct {
	ID         string `json:"id"`
	CreateAt   int64  `json:"create_at"`
	UpdateAt   int64  `json:"update_at"`
	DeleteAt   int64  `json:"delete_at"`
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email"`
	AllowStats bool   `json:"allow_stats"`
	Locale     string `json:"locale"`
	IsAdmin    bool   `json:"is_admin,omitempty"`
}

func (p *Player) IsValid() *AppError {
	if len(p.ID) != ID_LENGTH {
		return NewAppError("Player.IsValid", "Player ID is invalid", http.StatusBadRequest)
	}

	if p.CreateAt <= 0 {
		return NewAppError("Player.IsValid", "Created at is 0", http.StatusUnprocessableEntity)
	}

	if p.UpdateAt <= 0 {
		return NewAppError("Player.IsValid", "Updated at is 0", http.StatusUnprocessableEntity)
	}

	if !IsValidUsername(p.Username) {
		return NewAppError("Player.IsValid", "Username is invalid", http.StatusUnprocessableEntity)
	}

	if len(p.Email) > 128 || len(p.Email) == 0 || !strings.Contains(p.Email, "@") {
		return NewAppError("Player.IsValid", "Email is invalid", http.StatusBadRequest)
	}

	if len(p.Password) < MIN_PASSWORD_LENGTH || len(p.Password) > MAX_PASSWORD_LENGTH {
		return NewAppError("Player.IsValid", "Password is invalid", http.StatusBadRequest)
	}

	return nil
}

func (p *Player) PreSave() {
	if p.ID == "" {
		p.ID = NewID()
	}

	if p.Username == "" {
		p.Username = NewID()
	}

	p.Username = strings.ToLower(p.Username)
	p.Email = strings.ToLower(p.Email)
	p.Locale = strings.ToLower(p.Locale)
	p.Password = HashPassword(p.Password)

	p.CreateAt = GetMillis()
	p.UpdateAt = p.CreateAt

	if p.Locale == "" {
		p.Locale = DEFAULT_LOCALE
	}
}

func (p *Player) PreUpdate() {
	p.Username = strings.ToLower(p.Username)
	p.Email = strings.ToLower(p.Email)
	p.Locale = strings.ToLower(p.Locale)
	p.UpdateAt = GetMillis()
}

func (p *Player) ToJson() string {
	json, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(json)
}

func PlayerFromJson(data io.Reader) *Player {
	decoder := json.NewDecoder(data)
	var p Player
	err := decoder.Decode(&p)
	if err == nil {
		return &p
	}

	return nil
}

func PlayersToJson(p []*Player) string {
	json, err := json.Marshal(p)
	if err == nil {
		return string(json)
	}

	return "[]"
}

func PlayerToJson(p []*Player) string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(b)
}

func ComparePassword(hash string, password string) bool {
	if len(password) == 0 || len(hash) == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

var validUsernameChars = regexp.MustCompile(`^[a-z0-9\.\-_]+$`)

var restrictedUsernames = []string{
	"test",
	"admin",
}

func IsValidUsername(s string) bool {
	if len(s) < MIN_USERNAME_LENGTH || len(s) > MAX_USERNAME_LENGTH {
		return false
	}

	if !validUsernameChars.MatchString(s) {
		return false
	}

	for _, restrictedUsername := range restrictedUsernames {
		if s == restrictedUsername {
			return false
		}
	}

	return true
}

func (p *Player) Sanitize() {
	p.Password = ""
}
