// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	"database/sql"
	"fmt"
	"github.com/davidlu1997/gogogo/model"
)

type GameStore struct {
	*SqlStore
}

func NewGameStore(sqlStore *SqlStore) GameStore {
	gs := &GameStore{sqlStore}

	for _, db := range sqlStore.GetAllConns() {
		table := db.AddTableWithName(model.Game{}, "Games").SetKeys(false, "Id")
		table.ColMap("Id").SetMaxSize(24)
		table.ColMap("IdBlack").SetMaxSize(24)
		table.ColMap("IdWhite").SetMaxSize(24)
		table.ColMap("Board").SetMaxSize(400)
		table.ColMap("NumLines").SetMaxSize(2)
		table.ColMap("Turn").SetMaxSize(1)
		table.ColMap("CreateAt").SetMaxSize(20)
		table.ColMap("UpdateAt").SetMaxSize(20)
		table.ColMap("DeleteAt").SetMaxSize(20)
		table.ColMap("Finished").SetMaxSize(1)
	}

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

		if err := gs.GetMaster.Insert(game); err != nil {
			result.Err = model.NewLocError("GameStore.Save", "Game saving error", nil, "game_id="+game.Id+", "+err.Error())
		} else {
			result.Data = player
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

		if oldGameResult, err := gs.GetMaster().Get(model.Game{}, game.Id); err != nil {
			result.Err = model.NewLocError("GameStore.Update", "", nil, "game_id="+game.Id+", "+err.Error())
		} else if oldPlayerResult == nil {
			result.Err = model.NewLocError("GameStore.Update", "Cannot find game to update", nil, "game_id="+game.Id+", "+err.Error())
		} else {
			oldGame := oldGameResult.(*model.Game)
			game.CreateAt = oldGame.CreateAt

			if count, err := gs.GetMaster().Update(game); err != nil {
				result.Err = model.NewLocError("GameStore.Update", "Game updating error", nil, "game_id="+game.Id+", "+err.Error())
			} else if count != 1 {
				result.Err = model.NewLocError("GameStore.Update", "Game not unique", nil, fmt.Sprintf("game_id=%v, count=%v", game.Id, count))
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

		if obj, err := gs.GetMaster().Get(model.Game{}, id); err != nil {
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
		if _, err := gs.GetMaster().Select(&data, "SELECT * FROM Games"); err != nil {
			result.Err = model.NewLocError("GameStore.GetAll", "Get all games error", nil, err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)

	}()

	return storeChannel
}

func (gs GameStore) GetGamesByOnePlayerId(playerId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Game

		if _, err := gs.GetMaster().Select(&data, "SELECT * FROM Games WHERE IdBlack = :PlayerId OR IdWhite = :PlayerId", map[string]interface{}{"PlayerId": playerId}); err != nil {
			result.Err = model.NewLocError("GameStore.GetGamesByOnePlayerId", "Missing game error", nil, "player_id="+playerId+", "+err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) GetGamesByTwoPlayerId(player1Id, player2Id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Game

		if _, err := gs.GetMaster().Select(&data, "SELECT * FROM Games WHERE (IdBlack = :Player1Id AND IdWhite = :Player2Id) OR (IdBlack = :Player2Id AND IdWhite = :Player1Id)", map[string]interface{}{"Player1Id": player1Id, "Player2Id": player2Id}); err != nil {
			result.Err = model.NewLocError("GameStore.GetGamesByTwoPlayerId", "Missing game error", nil, "player1_id="+player1Id+", player2_id="+player2Id+", "+err.Error())
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

		if count, err := gs.GetMaster().SelectInt("SELECT COUNT(Id) FROM Games"); err != nil {
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

		if count, err := gs.GetMaster().SelectInt("SELECT COUNT(Id) FROM Games WHERE Finished = 1"); err != nil {
			result.Err = model.NewLocError("GameStore.GetTotalFinishedGamesCount", "Get total finished games count error", nil, err.Error())
		} else {
			result.Data = count
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (gs GameStore) PermanentDelete(gameId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := gs.GetMaster().Exec("DELETE FROM Games WHERE Id = :GameId", map[string]interface{}{"GameId": gameId}); err != nil {
			result.Err = model.NewLocError("GameStore.PermanentDelete", "Permanent delete game error", nil, "game_id="+gameId+", "+err.Error())
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
