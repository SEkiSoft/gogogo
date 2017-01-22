// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"fmt"

	"github.com/sekisoft/gogogo/model"
)

type GameStore struct {
	*SqlStore
}

func NewGameStore(sqlStore *SqlStore) SqlGameStore {
	gs := &GameStore{sqlStore}

	db := sqlStore.GetMaster()
	table := db.AddTableWithName(model.Game{}, "Games").SetKeys(false, "ID")
	table.ColMap("ID").SetMaxSize(model.ID_LENGTH)
	table.ColMap("IDBlack").SetMaxSize(model.ID_LENGTH)
	table.ColMap("IDWhite").SetMaxSize(model.ID_LENGTH)
	table.ColMap("Board").SetMaxSize(400)
	table.ColMap("NumLines").SetMaxSize(2)
	table.ColMap("Turn").SetMaxSize(1)
	table.ColMap("CreateAt").SetMaxSize(20)
	table.ColMap("UpdateAt").SetMaxSize(20)
	table.ColMap("DeleteAt").SetMaxSize(20)
	table.ColMap("Finished").SetMaxSize(1)

	return gs
}

func (gs GameStore) Save(game *model.Game) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		game.PreSave()
		if result.Err = game.IsValid(); result.Err != nil {
			storeChannel <- result
			close(storeChannel)
			return
		}

		if err := gs.GetMaster().Insert(game); err != nil {
			result.Err = model.NewLocError("GameStore.Save", "Game saving error", nil, "game_id="+game.ID+", "+err.Error())
		} else {
			result.Data = game
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) Update(game *model.Game) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		game.PreUpdate()

		if result.Err = game.IsValid(); result.Err != nil {
			storeChannel <- result
			close(storeChannel)
			return
		}

		if oldGameResult, err := gs.GetMaster().Get(model.Game{}, game.ID); err != nil {
			result.Err = model.NewLocError("GameStore.Update", "", nil, "game_id="+game.ID+", "+err.Error())
		} else if oldGameResult == nil {
			result.Err = model.NewLocError("GameStore.Update", "Cannot find game to update", nil, "game_id="+game.ID+", "+err.Error())
		} else {
			oldGame := oldGameResult.(*model.Game)
			game.CreateAt = oldGame.CreateAt

			if count, err := gs.GetMaster().Update(game); err != nil {
				result.Err = model.NewLocError("GameStore.Update", "Game updating error", nil, "game_id="+game.ID+", "+err.Error())
			} else if count != 1 {
				result.Err = model.NewLocError("GameStore.Update", "Game not unique", nil, fmt.Sprintf("game_id=%v, count=%v", game.ID, count))
			} else {
				result.Data = [2]*model.Game{game, oldGame}
			}
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) Get(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if obj, err := gs.GetReplica().Get(model.Game{}, id); err != nil {
			result.Err = model.NewLocError("GameStore.Get", "Get game by id error", nil, "game_id="+id+", "+err.Error())
		} else if obj == nil {
			result.Err = model.NewLocError("GameStore.Get", "Missing game error", nil, "player_id="+id)
		} else {
			result.Data = obj.(*model.Game)
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) GetAll() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Game
		if _, err := gs.GetReplica().Select(&data, "SELECT * FROM Games"); err != nil {
			result.Err = model.NewLocError("GameStore.GetAll", "Get all games error", nil, err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)

	}()

	return storeChannel
}

func (gs GameStore) GetGamesByOnePlayerID(playerID string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Game

		if _, err := gs.GetReplica().Select(&data, "SELECT * FROM Games WHERE IDBlack = :PlayerID OR IDWhite = :PlayerID", map[string]interface{}{"PlayerID": playerID}); err != nil {
			result.Err = model.NewLocError("GameStore.GetGamesByOnePlayerID", "Missing game error", nil, "player_id="+playerID+", "+err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) GetGamesByTwoPlayerID(player1ID, player2ID string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Game

		if _, err := gs.GetReplica().Select(&data, "SELECT * FROM Games WHERE (IDBlack = :Player1ID AND IDWhite = :Player2ID) OR (IDBlack = :Player2ID AND IDWhite = :Player1ID)", map[string]interface{}{"Player1ID": player1ID, "Player2ID": player2ID}); err != nil {
			result.Err = model.NewLocError("GameStore.GetGamesByTwoPlayerID", "Missing game error", nil, "player1_id="+player1ID+", player2_id="+player2ID+", "+err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) GetTotalGamesCount() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if count, err := gs.GetReplica().SelectInt("SELECT COUNT(ID) FROM Games"); err != nil {
			result.Err = model.NewLocError("GameStore.GetTotalPlayersCount", "Get total games count error", nil, err.Error())
		} else {
			result.Data = count
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) GetTotalFinishedGamesCount() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if count, err := gs.GetReplica().SelectInt("SELECT COUNT(ID) FROM Games WHERE Finished = 1"); err != nil {
			result.Err = model.NewLocError("GameStore.GetTotalFinishedGamesCount", "Get total finished games count error", nil, err.Error())
		} else {
			result.Data = count
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) Delete(gameID string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := gs.GetMaster().Exec("DELETE FROM Games WHERE ID = :GameID", map[string]interface{}{"GameID": gameID}); err != nil {
			result.Err = model.NewLocError("GameStore.PermanentDelete", "Permanent delete game error", nil, "game_id="+gameID+", "+err.Error())
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
