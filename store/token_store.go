// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"net/http"

	"github.com/sekisoft/gogogo/model"
)

type TokenStore struct {
	*SqlStore
}

func NewTokenStore(sqlStore *SqlStore) SqlTokenStore {
	ts := &TokenStore{sqlStore}

	db := sqlStore.GetMaster()
	table := db.AddTableWithName(model.Token{}, "Tokens").SetKeys(false, "ID")
	table.ColMap("ID").SetMaxSize(model.ID_LENGTH)
	table.ColMap("PlayerID").SetMaxSize(model.ID_LENGTH)
	table.ColMap("DeviceID").SetMaxSize(model.ID_LENGTH)
	table.ColMap("Roles").SetMaxSize(64)

	return ts
}

func (ts TokenStore) Save(token *model.Token) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		token.PreSave()

		if err := ts.GetMaster().Insert(token); err != nil {
			result.Err = model.NewAppError("TokenStore.Save", err.Error(), http.StatusBadGateway)
		} else {
			result.Data = token
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ts TokenStore) Get(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if obj, err := ts.GetReplica().Get(model.Token{}, id); err != nil {
			result.Err = model.NewAppError("TokenStore.Get", err.Error(), http.StatusBadGateway)
		} else if obj == nil {
			result.Err = model.NewAppError("TokenStore.Get", err.Error(), http.StatusBadGateway)
		} else {
			result.Data = obj.(*model.Token)
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ts TokenStore) GetTokens(playerID string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Token
		if _, err := ts.GetReplica().Select(&data, "SELECT * FROM Tokens WHERE PlayerID = :PlayerID", map[string]interface{}{"PlayerID": playerID}); err != nil {
			result.Err = model.NewAppError("TokenStore.Get", err.Error(), http.StatusBadGateway)
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ts TokenStore) Delete(id string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := ts.GetMaster().Exec("DELETE FROM Tokens WHERE ID = :ID", map[string]interface{}{"ID": id}); err != nil {
			result.Err = model.NewAppError("TokenStore.Delete", err.Error(), http.StatusBadGateway)
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ts TokenStore) DeleteAll(playerID string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := ts.GetMaster().Exec("DELETE * FROM Tokens WHERE PlayerID = :PlayerID", map[string]interface{}{"PlayerID": playerID}); err != nil {
			result.Err = model.NewAppError("TokenStore.DeleteAll", err.Error(), http.StatusBadGateway)
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ts TokenStore) GetAll() StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Token
		if _, err := ts.GetReplica().Select(&data, "SELECT * FROM Tokens"); err != nil {
			result.Err = model.NewAppError("TokenStore.GetAll", err.Error(), http.StatusBadGateway)
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
