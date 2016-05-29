// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	"database/sql"
	"fmt"
	"github.com/davidlu1997/gogogo/model"
)

type SqlMoveStore struct {
	*SqlStore
}

func NewSqlMoveStore(sqlStore *SqlStore) MoveStore {
	ms := &SqlMoveStore(sqlStore)

	for _, db := range sqlStore.GetAllConns() {
		table := db.AddTableWithName(model.Move{}, "Moves").Setkeys(false, "Id")
		table.ColMap("Id").SetMaxSize(24)
		table.ColMap("PlayerId").SetMaxSize(24)
		table.ColMap("GameId").SetMaxSize(24)
		table.ColMap("X").SetMaxSize(2)
		table.ColMap("Y").SetMaxSize(2)
		table.ColMap("CreateAt").SetMaxSize(20)
	}

	return ms
}

func (ms SqlMoveStore) Save(move *model.Move) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		move.PreSave()

		if err := ms.GetMaster().Insert(player); err != nil {
			result.Err = model.NewLocError("SqlMoveStore.Save", "Move saving error", nil, "move_id="+move.Id+", "+err.Error())
		} else {
			result.Data = move
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms SqlMoveStore) Get(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if obj, err := ms.GetMaster().Get(model.Move{}, id); err != nil {
			result.Err = model.NewLocError("SqlMoveStore.Get", "Get move by id error", nil, "move_id="+id+", "+err.Error())
		} else if obj == nil {
			result.Err = model.NewLocError("SqlMoveStore.Get", "Missing move error", nil, "move_id="+id)
		} else {
			result.Data = obj.(*model.Move)
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms SqlMoveStore) GetByGame(gameId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Move

		if err := ms.GetMaster().SelectOne(&move, "SELECT * FROM Moves WHERE GameId = :GameId", map[string]interface{}{"GameId": gameId}); err != nil {
			result.Err = model.NewLocError("SqlMoveStore.GetByGame", "Missing game error", nil, "game_id="+gameId+", "+err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms SqlMoveStore) GetByPlayer(playerId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Move

		if err := ms.GetMaster().SelectOne(&data, "SELECT * FROM Moves WHERE PlayerId = :PlayerId", map[string]interface{}{"PlayerId": playerId}); err != nil {
			result.Err = model.NewLocError("SqlMoveStore.GetByPlayer", "Missing player error", nil, "player_id="+playerId+", "+err.Error())
		}

		result.Data = &move

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms SqlMoveStore) GetTotalMovesCount() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if count, err := ms.GetMaster().SelectInt("SELECT COUNT(Id) FROM Moves"); err != nil {
			result.Err = model.NewLocError("SqlMoveStore.GetTotalMovesCount", "Get total moves count error", nil, err.Error())
		} else {
			result.Data = count
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms SqlMoveStore) PermanentDelete(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := ms.GetMaster().Exec("DELETE FROM Moves WHERE Id = :MoveId", map[string]interface{}{"MoveID": id}); err != nil {
			result.Err = model.NewLocError("SqlMoveStore.PermanentDelete", "Permanent delete move error", nil, "moveId="+id+", "+err.Error())
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
