// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	"database/sql"
	"fmt"
	"github.com/davidlu1997/gogogo/model"
	"strconv"
	"strings"
)

type SqlPlayerStore struct {
	*SqlStore
}

func NewSqlPlayerStore(sqlStore *SqlStore) PlayerStore {
	ps := &SqlPlayerStore{sqlStore}

	for _, db := range sqlStore.GetAllConns() {
		table := db.AddTableWithName(model.Player{}, "Players").SetKeys(false, "Id")
		table.ColMap("Id").SetMaxSize(24)
		table.ColMap("Username").SetMaxSize(64).SetUnique(true)
		table.ColMap("Password").SetMaxSize(128)
		table.ColMap("Email").SetMaxSize(128).SetUnique(true)
		table.ColMap("AllowStats").SetMaxSize(1)
		table.ColMap("Locale").SetMaxSize(5)
	}

	return ps
}

func (ps SqlPlayerStore) CreateIndexesIfNotExists() {
	ps.CreateIndexesIfNotExists("idx_player_email", "Players", "Email")
}

func (ps SqlPlayerStore) Save(player *model.Player) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		player.PreSave()
		if result.Err = player.IsValid(); result.Err != nil {
			storeChannel <- result
			close(storeChannel)
			return
		}

		if err := ps.GetMaster().Insert(user); err != nil {
			if IsUniqueConstraintError(err.Error(), []string{"Email", "players_email_key", "idx_players_email_unique"}) {
				result.Err = model.NewLocError("SqlUserStore.Save", "Email already exists", nil, "player_id="+player.Id+", "+err.Error())
			} else if IsUniqueConstraintError(err.Error(), []string{"Username", "players_username_key", "idx_players_username_unique"}) {
				result.Err = model.NewLocError("SqlUserStore.Save", "Username already exists", nil, "player_id="+player.Id+", "+err.Error())
			} else {
				result.Err = model.NewLocError("SqlUserStore.Save", "Player saving error", nil, "user_id="+player.Id+", "+err.Error())
			}
		} else {
			result.Data = player
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
