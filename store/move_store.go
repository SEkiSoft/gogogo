// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"github.com/SEkiSoft/gogogo/model"
)

type MoveStore struct {
	*SqlStore
}

func NewMoveStore(sqlStore *SqlStore) SqlMoveStore {
	ms := &MoveStore{sqlStore}

	db := sqlStore.GetMaster()
	table := db.AddTableWithName(model.Move{}, "Moves").SetKeys(false, "Id")
	table.ColMap("Id").SetMaxSize(24)
	table.ColMap("PlayerId").SetMaxSize(24)
	table.ColMap("GameId").SetMaxSize(24)
	table.ColMap("X").SetMaxSize(2)
	table.ColMap("Y").SetMaxSize(2)
	table.ColMap("CreateAt").SetMaxSize(20)

	return ms
}

func (ms MoveStore) Save(move *model.Move) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		move.PreSave()

		if err := ms.GetMaster().Insert(move); err != nil {
			result.Err = model.NewLocError("MoveStore.Save", "Move saving error", nil, "move_id="+move.Id+", "+err.Error())
		} else {
			result.Data = move
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms MoveStore) Get(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if obj, err := ms.GetMaster().Get(model.Move{}, id); err != nil {
			result.Err = model.NewLocError("MoveStore.Get", "Get move by id error", nil, "move_id="+id+", "+err.Error())
		} else if obj == nil {
			result.Err = model.NewLocError("MoveStore.Get", "Missing move error", nil, "move_id="+id)
		} else {
			result.Data = obj.(*model.Move)
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms MoveStore) GetByGame(gameId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Move

		if err := ms.GetMaster().SelectOne(&data, "SELECT * FROM Moves WHERE GameId = :GameId", map[string]interface{}{"GameId": gameId}); err != nil {
			result.Err = model.NewLocError("MoveStore.GetByGame", "Missing game error", nil, "game_id="+gameId+", "+err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms MoveStore) GetAll() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}
		var data []*model.Move

		if err := ms.GetMaster().SelectOne(&data, "SELECT * FROM Moves"); err != nil {
			result.Err = model.NewLocError("MoveStore.GetAll", "Couldn't retrieve moves", nil, err.Error())
		}
		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms MoveStore) GetByPlayer(playerId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Move

		if err := ms.GetMaster().SelectOne(&data, "SELECT * FROM Moves WHERE PlayerId = :PlayerId", map[string]interface{}{"PlayerId": playerId}); err != nil {
			result.Err = model.NewLocError("MoveStore.GetByPlayer", "Missing player error", nil, "player_id="+playerId+", "+err.Error())
		}

		result.Data = &data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms MoveStore) GetTotalMovesCount() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if count, err := ms.GetMaster().SelectInt("SELECT COUNT(Id) FROM Moves"); err != nil {
			result.Err = model.NewLocError("MoveStore.GetTotalMovesCount", "Get total moves count error", nil, err.Error())
		} else {
			result.Data = count
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ms MoveStore) Delete(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := ms.GetMaster().Exec("DELETE FROM Moves WHERE Id = :MoveId", map[string]interface{}{"MoveID": id}); err != nil {
			result.Err = model.NewLocError("MoveStore.PermanentDelete", "Permanent delete move error", nil, "moveId="+id+", "+err.Error())
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
