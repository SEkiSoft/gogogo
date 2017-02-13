// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"github.com/sekisoft/gogogo/model"
)

type StoreResult struct {
	Data interface{}
	Err  *model.AppError
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
	Token() SqlTokenStore
	Close()
	DropAllTables()
}

type SqlGameStore interface {
	Save(game *model.Game) StoreChannel
	Update(game *model.Game) StoreChannel
	Get(id string) StoreChannel
	GetAll() StoreChannel
	GetGamesByOnePlayerID(playerID string) StoreChannel
	GetGamesByTwoPlayerID(player1ID, player2ID string) StoreChannel
	GetTotalGamesCount() StoreChannel
	GetTotalFinishedGamesCount() StoreChannel
	Delete(gameID string) StoreChannel
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
	GetByGame(gameID string) StoreChannel
	GetAll() StoreChannel
	GetByPlayer(playerID string) StoreChannel
	GetTotalMovesCount() StoreChannel
	Delete(id string) StoreChannel
}

type SqlTokenStore interface {
	Save(token *model.Token) StoreChannel
	Get(id string) StoreChannel
	GetTokens(playerID string) StoreChannel
	Delete(id string) StoreChannel
	DeleteAll(playerID string) StoreChannel
	GetAll() StoreChannel
}
