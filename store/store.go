// Copyright (c) 2016 SEkiSoft
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
	Game() SqlGameStore
	Player() SqlPlayerStore
	Move() SqlMoveStore
	Close()
	DropAllTables()
}

type SqlGameStore interface {
	Save(game *model.Game) StoreChannel
	Update(game *model.Game) StoreChannel
	Get(id string) StoreChannel
	GetAll() StoreChannel
	GetGamesByOnePlayerId(playerId string) StoreChannel
	GetGamesByTwoPlayerId(player1Id, player2Id string) StoreChannel
	GetTotalGamesCount() StoreChannel
	GetTotalFinishedGamesCount() StoreChannel
	Delete(gameId string) StoreChannel
}

type SqlPlayerStore interface {
	Save(player *model.Player) StoreChannel
	Update(player *model.Player) StoreChannel
	UpdatePassword(id, newPassword string) StoreChannel
	Get(id string) StoreChannel
	GetAll() StoreChannel
	GetPlayerGames(id string) StoreChannel
	GetByEmail(email string) StoreChannel
	GetByUsername(username string) StoreChannel
	GetTotalPlayersCount() StoreChannel
	Delete(id string) StoreChannel
}

type SqlMoveStore interface {
	Save(move *model.Move) StoreChannel
	Get(id string) StoreChannel
	GetByGame(gameId string) StoreChannel
	GetAll() StoreChannel
	GetByPlayer(playerId string) StoreChannel
	GetTotalMovesCount() StoreChannel
	Delete(id string) StoreChannel
}
