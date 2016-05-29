// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/davidlu1997/gogogo/model"
	"time"
)

type StoreResult struct {
	Data interface{}
	Err  *model.Error
}

type StoreChannel chan StoreResult

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
	GetGamesByOnePlayerId(playerId string) StoreChannel
	GetGamesByTwoPlayerId(player1Id, player2Id string) StoreChannel
	GetGamesByOnePlayerName(playerName string) StoreChannel
	GetGamesByTwoPlayerName(player1Name, player2Name string) StoreChannel
	GetAll() StoreChannel
	GetTotalGamesCount() StoreChannel
	GetTotalActiveGamesCount() StoreChannel
	PermanentDelete(gameId String) StoreChannel
}

type PlayerStore interface {
	Save(player *model.Player) StoreChannel
	Update(player *model.Player) StoreChannel
	UpdateUpdateAt(playerId string) StoreChannel
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
	GetByUser(userId string) StoreChannel
	GetTotalMovesCount() StoreChannel
	GetTotalRecentMovesCount() StoreChannel
	PermanentDelete(id string) StoreChannel
}
