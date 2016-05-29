// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	"github.com/davidlu1997/gogogo/model"
)

type StoreResult struct {
	Data interface{}
	Err  *model.Error
}

type StoreChannel chan StoreResult

func Must(sc StoreChannel) interface{} {
	r := <-sc
	if r.Err != nil {
		panic(r.Err)
	}

	return r.Data
}

type Store interface {
	Game() GameStore
	Player() PlayerStore
	Move() MoveStore
	Close()
	DropAllTables()
}

type GameStore interface {
	Save(game *model.Game) StoreChannel
	Update(game *model.Game) StoreChannel
	Get(id string) StoreChannel
	GetAll() StoreChannel
	GetGamesByOnePlayerId(playerId string) StoreChannel
	GetGamesByTwoPlayerId(player1Id, player2Id string) StoreChannel
	GetTotalGamesCount() StoreChannel
	GetTotalFinishedGamesCount() StoreChannel
	PermanentDelete(gameId String) StoreChannel
}

type PlayerStore interface {
	Save(player *model.Player) StoreChannel
	Update(player *model.Player) StoreChannel
	UpdatePassword(playerId, newPassword string) StoreChannel
	Get(id string) StoreChannel
	GetAll() StoreChannel
	GetByEmail(email string) StoreChannel
	GetByUsername(username string) StoreChannel
	GetTotalPlayersCount() StoreChannel
	PermanentDelete(playerId string) StoreChannel
}

type MoveStore interface {
	Save(move *model.Move) StoreChannel
	Get(id string) StoreChannel
	GetByGame(gameId string) StoreChannel
	GetByPlayer(playerId string) StoreChannel
	GetTotalMovesCount() StoreChannel
	PermanentDelete(id string) StoreChannel
}
