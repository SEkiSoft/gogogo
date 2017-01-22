// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"fmt"
	"strings"

	"github.com/sekisoft/gogogo/model"
)

type PlayerStore struct {
	*SqlStore
}

func NewPlayerStore(sqlStore *SqlStore) SqlPlayerStore {
	ps := &PlayerStore{sqlStore}

	db := sqlStore.GetMaster()
	table := db.AddTableWithName(model.Player{}, "Players").SetKeys(false, "ID")
	table.ColMap("ID").SetMaxSize(model.ID_LENGTH)
	table.ColMap("Username").SetMaxSize(64).SetUnique(true)
	table.ColMap("Password").SetMaxSize(128)
	table.ColMap("Email").SetMaxSize(128).SetUnique(true)
	table.ColMap("AllowStats").SetMaxSize(1)
	table.ColMap("Locale").SetMaxSize(5)

	return ps
}

func (ps PlayerStore) Save(player *model.Player) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		player.PreSave()
		if result.Err = player.IsValid(); result.Err != nil {
			storeChannel <- result
			close(storeChannel)
			return
		}

		if err := ps.GetMaster().Insert(player); err != nil {
			if IsUniqueConstraintError(err.Error(), []string{"Email", "players_email_key", "idx_players_email_unique"}) {
				result.Err = model.NewLocError("PlayerStore.Save", "Email already exists", nil, "player_id="+player.ID+", "+err.Error())
			} else if IsUniqueConstraintError(err.Error(), []string{"Playername", "players_username_key", "idx_players_username_unique"}) {
				result.Err = model.NewLocError("PlayerStore.Save", "Username already exists", nil, "player_id="+player.ID+", "+err.Error())
			} else {
				result.Err = model.NewLocError("PlayerStore.Save", "Player saving error", nil, "player_id="+player.ID+", "+err.Error())
			}
		} else {
			result.Data = player
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ps PlayerStore) Update(player *model.Player) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		player.PreUpdate()

		if result.Err = player.IsValid(); result.Err != nil {
			storeChannel <- result
			close(storeChannel)
			return
		}

		if oldPlayerResult, err := ps.GetMaster().Get(model.Player{}, player.ID); err != nil {
			result.Err = model.NewLocError("PlayerStore.Update", "", nil, "player_id="+player.ID+", "+err.Error())
		} else if oldPlayerResult == nil {
			result.Err = model.NewLocError("PlayerStore.Update", "Cannot find player to update", nil, "player_id="+player.ID)
		} else {
			oldPlayer := oldPlayerResult.(*model.Player)
			player.CreateAt = oldPlayer.CreateAt
			player.Password = oldPlayer.Password

			if count, err := ps.GetMaster().Update(player); err != nil {
				if IsUniqueConstraintError(err.Error(), []string{"Email", "players_email_key", "idx_players_email_unique"}) {
					result.Err = model.NewLocError("PlayerStore.Update", "Email already exists", nil, "player_id="+player.ID+", "+err.Error())
				} else if IsUniqueConstraintError(err.Error(), []string{"Username", "players_username_key", "idx_players_username_unique"}) {
					result.Err = model.NewLocError("PlayerStore.Update", "Username already exists", nil, "player_id="+player.ID+", "+err.Error())
				} else {
					result.Err = model.NewLocError("PlayerStore.Update", "Player updating error", nil, "player_id="+player.ID+", "+err.Error())
				}
			} else if count != 1 {
				result.Err = model.NewLocError("PlayerStore.Update", "Player not unique", nil, fmt.Sprintf("player_id=%v, count=%v", player.ID, count))
			} else {
				result.Data = [2]*model.Player{player, oldPlayer}
			}
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ps PlayerStore) UpdatePassword(playerID string, newPassword string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := ps.GetMaster().Exec("UPDATE Players SET Password = :Password WHERE ID = :PlayerID", map[string]interface{}{"Password": newPassword, "PlayerID": playerID}); err != nil {
			result.Err = model.NewLocError("PlayerStore.UpdatePassword", "Player update password error", nil, "player_id="+playerID)
		} else {
			result.Data = playerID
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ps PlayerStore) Get(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if obj, err := ps.GetReplica().Get(model.Player{}, id); err != nil {
			result.Err = model.NewLocError("PlayerStore.Get", "Get player by id error", nil, "player_id="+id+", "+err.Error())
		} else if obj == nil {
			result.Err = model.NewLocError("PlayerStore.Get", "Missing player error", nil, "player_id="+id)
		} else {
			result.Data = obj.(*model.Player)
		}

		storeChannel <- result
		close(storeChannel)

	}()

	return storeChannel
}

func (ps PlayerStore) GetAll() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Player
		if _, err := ps.GetReplica().Select(&data, "SELECT * FROM Players"); err != nil {
			result.Err = model.NewLocError("PlayerStore.GetAll", "Get all players error", nil, err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)

	}()

	return storeChannel
}

func (ps PlayerStore) GetPlayerGames(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Game
		if _, err := ps.GetReplica().Select(&data, "SELECT * FROM Games WHERE IDBlack = :ID OR IDWhite = :ID", map[string]interface{}{"ID": id}); err != nil {
			result.Err = model.NewLocError("PlayerStore.GetPlayerGames", "Get player games error", nil, err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ps PlayerStore) GetByEmail(email string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		email = strings.ToLower(email)

		player := model.Player{}

		if err := ps.GetReplica().SelectOne(&player, "SELECT * FROM Players WHERE Email = :Email", map[string]interface{}{"Email": email}); err != nil {
			result.Err = model.NewLocError("PlayerStore.GetByEmail", "Missing player error", nil, "email="+email+", "+err.Error())
		}

		result.Data = &player

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ps PlayerStore) GetByUsername(username string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		username = strings.ToLower(username)

		player := model.Player{}

		if err := ps.GetReplica().SelectOne(&player, "SELECT * FROM Players WHERE Username = :Username", map[string]interface{}{"Username": username}); err != nil {
			result.Err = model.NewLocError("PlayerStore.GetByUsername", "Missing player error", nil, "username="+username+", "+err.Error())
		}

		result.Data = &player

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ps PlayerStore) GetTotalPlayersCount() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if count, err := ps.GetReplica().SelectInt("SELECT COUNT(ID) FROM Players"); err != nil {
			result.Err = model.NewLocError("PlayerStore.GetTotalPlayersCount", "Get total players count error", nil, err.Error())
		} else {
			result.Data = count
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ps PlayerStore) Delete(playerID string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := ps.GetMaster().Exec("DELETE FROM Players WHERE ID = :PlayerID", map[string]interface{}{"PlayerID": playerID}); err != nil {
			result.Err = model.NewLocError("PlayerStore.PermanentDelete", "Permanent delete player error", nil, "player_id="+playerID+", "+err.Error())
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
