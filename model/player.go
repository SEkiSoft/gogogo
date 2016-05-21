package model

import (
	"strings"
	"encoding/json"
	"io"
	"regexp"
)

const (
	DEFAULT_LOCALE = "en"
	MIN_PASSWORD_LENGTH = 5
	MIN_USERNAME_LENGTH = 4
	MAX_USERNAME_LENGTH = 24
)

type Player struct {
	Id 			string 	`json:"id"`
	CreateAt 	int64 	`json:"create_at"`
	UpdateAt 	int64 	`json:"update_at"`
	DeleteAt 	int64 	`json:"delete_at"`
	Username 	string 	`json:"username"`
	Password	string	`json:"password,omitempty"`
	Email		string	`json:"email"`
	AllowStats	bool	`json:"allow_stats"`
	Locale		string	`json:"locale"`
}

func (p *Player) IsValid() *Error {
	if len(p.Id) != 26 {
		return NewLocError("Player.IsValid", "model.player.is_valid.id.error", nil, "")
	}

	if p.CreateAt == 0 {
		return NewLocError("Player.IsValid", "model.player.is_valid.create_at.error", nil, "player_id="+p.Id)
	}

	if p.CreateAt == 0 {
		return NewLocError("Player.IsValid", "model.player.is_valid.update_at.error", nil, "player_id="+p.Id)
	}

	if !IsValidUsername(p.Username) {
		return NewLocError("Player.IsValid", "model.player.is_valid.username.error", nil, "player_id="+p.Id)
	}

	if len(p.Email) > 128 || len(p.Email) == 0 || !strings.Contains(p.Email, "@") {
		return NewLocError("Player.IsValid", "model.player.is_valid.email.error", nil, "player_id="+p.Id)
	}

	return nil
}

func (p *Player) PreSave() {
	if p.Id == "" {
		p.Id = NewId()
	}

	if p.Username == "" {
		p.Username = NewId()
	}

	p.Username = strings.ToLower(p.Username)
	p.Email = strings.ToLower(p.Email)
	p.Locale = strings.ToLower(p.Locale)

	p.CreateAt = GetMillis()
	p.UpdateAt = p.CreateAt

	if p.Locale == "" {
		p.Locale = DEFAULT_LOCALE
	}

	if len(p.Password) > 0 {
		p.Password = HashPassword(p.Password)
	}
}

func (p *Player) PreUpdate() {
	p.Username = strings.ToLower(p.Username)
	p.Email = strings.ToLower(p.Email)
	p.Locale = strings.ToLower(p.Locale)
	p.UpdateAt = GetMillis()
}

func (p *Player) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func UserFromJson(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	var user User 
	err := decoder.Decode(&user)
	if err == nil {
		return user
	} else {
		return nil
	}
}

func ComparePassword(hash string, password string) bool {
	if len(password) == 0 || len(has) == 0 {
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

func (p *Player) Etag() string {
	return Etag(p.Id, p.UpdateAt)
}