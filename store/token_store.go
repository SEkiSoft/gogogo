// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"github.com/SEkiSoft/gogogo/model"
)

type TokenStore struct {
	*SqlStore
}

func NewTokenStore(sqlStore *SqlStore) SqlTokenStore {
	ts := &TokenStore{sqlStore}

	db := sqlStore.GetMaster()
	table := db.AddTableWithName(model.Token{}, "Tokens").SetKeys(false, "Id")
	table.ColMap("Id").SetMaxSize(model.ID_LENGTH)
	table.ColMap("PlayerId").SetMaxSize(model.ID_LENGTH)
	table.ColMap("DeviceId").SetMaxSize(model.ID_LENGTH)
	table.ColMap("Roles").SetMaxSize(64)

	return ts
}

func (ts TokenStore) Save(token *model.Token) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		token.PreSave()

		if err := ts.GetMaster().Insert(token); err != nil {
			result.Err = model.NewLocError("TokenStore.Save", "Token saving error", nil, "token_id="+token.Id+", "+err.Error())
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
			result.Err = model.NewLocError("TokenStore.Get", "Get token by id error", nil, "token_id="+id+", "+err.Error())
		} else if obj == nil {
			result.Err = model.NewLocError("TokenStore.Get", "Missing token error", nil, "token_id="+id)
		} else {
			result.Data = obj.(*model.Token)
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ts TokenStore) GetTokens(playerId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		var data []*model.Token
		if _, err := ts.GetReplica().Select(&data, "SELECT * FROM Tokens WHERE PlayerId = :PlayerId", map[string]interface{}{"PlayerId": playerId}); err != nil {
			result.Err = model.NewLocError("TokenStore.Get", "Get tokens by player id error", nil, err.Error())
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

		if _, err := ts.GetMaster().Exec("DELETE FROM Tokens WHERE Id = :Id", map[string]interface{}{"Id": id}); err != nil {
			result.Err = model.NewLocError("TokenStore.Delete", "Delete token error", nil, "token_id="+id+", "+err.Error())
		}

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (ts TokenStore) DeleteAll(playerId string) StoreChannel {
	storeChannel := make(StoreChannel)

	go func() {
		result := StoreResult{}

		if _, err := ts.GetMaster().Exec("DELETE * FROM Tokens WHERE PlayerId = :PlayerId", map[string]interface{}{"PlayerId": playerId}); err != nil {
			result.Err = model.NewLocError("TokenStore.DeleteAll", "Delete all tokens by player id error", nil, err.Error())
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
			result.Err = model.NewLocError("TokenStore.GetAll", "Get all tokens error", nil, err.Error())
		}

		result.Data = data

		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}
