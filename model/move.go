package model

type Move struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
	X        uint   `json:"move_x"`
	Y        uint   `json:"move_y"`
	CreateAt int64  `json:"create_at"`
}

func (m *Move) ToJson() string {
	s, err := json.Marshal(m)
	if err != nil {
		return ""
	} else {
		return string(s)
	}
}

func (m *Move) PreSave() string {
	m.CreateAt = GetMillis()
}

func MoveFromJson(data io.Reader) *Move {
	decoder := json.NewDecoder(data)
	var m Move
	err := decoder.Decode(&m)
	if err == nil {
		return &m
	} else {
		return nil
	}
}
