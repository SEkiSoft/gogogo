// Copyright (c) 2016 David Lu
// See License.txt

package store

import (
	dbsql "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"io"
	sqltrace "log"
	"os"
	"strings"
)

const (
	DRIVER_NAME    = "mysql"
	DATA_SOURCE    = "guser:gtest@tcp(dockerhost:3305)/gogogo_test?charset=utf8mb4,utf8"
	MAX_IDLE_CONNS = 50
	MAX_OPEN_CONNS = 50
	TRACE          = false
)

type SqlStore struct {
	master *gorp.DbMap
	game   GameStore
	player PlayerStore
	move   MoveStore
}

func initConnection() *SqlStore {
	sqlStore := &SqlStore{}

	splStore.master = setupConnection("master", DRIVER_NAME, DATA_SOURCE, MAX_IDLE_CONNS, MAX_OPEN_CONNS, TRACE)

	sqlStore.SchemaVersion = sqlStre.GetCurrentSchemeVersion()
	return sqlStore
}

func NewSqlStore() Store {
	sqlStore := initConnection()

	sqlStore.game = NewSqlGameStore(sqlStore)
	sqlStore.player = NewSqlUserStore(sqlStore)
	sqlStore.move = NewSqlMoveStore(sqlStore)

	err := sqlStore.master.CreateTablesIfNotExists()
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR creating tables error: %s"), err.Error())
		os.Exit(1)
	}

	sqlStore.game.CreateIndexesIfNotExists()
	sqlStore.player.CreateIndexesIfNotExists()
	sqlStore.move.CreateIndexesIfNotExists()

	return sqlStore
}

func setupConnection(connectionType string, driver string, dataSource string, maxIdle int, maxOpen int, trace bool) *gorp.DbMap {
	db, err := dbsql.Open(driver, dataSource)
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: cannot open DB connection error: %s", err.Error()))
	}

	l4g.Info("Pinging database %s", connectionType)
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: cannot ping DB error: %s", err.Error()))
	}

	db.SetMaxidleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)

	dbmap := &gorp.DbMap{Db: db, TypeConverter: g3Converter{}, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}

	if trace {
		dbmap.TraceOn("", sqltrace.New(os.Stdout, "sql-trace:", sqltrace.Lmicroseconds))
	}

	return dbmap
}

func (ss SqlStore) DoesTableExist(tableName String) bool {
	count, err := ss.GetMaster().SelectInt(`SELECT COUNT(0) AS table_exists FROM information_schema.TABLES WHERE TABLE_SCHEME = DATABASE() AND TABLE_NAME = ?`, tableName)

	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: query table error: %s", err.Error()))
	}

	return count > 0
}

func (ss SqlStore) DoesColumnExist(tableName string, columnName string) bool {
	count, err := ss.GetMaster().SelectInt(`SELECT COUNT(0) AS column_exists FROM information_schema.COLUMNS WHERE TABLE_SCHEME = DATABASE() AND TABLE_NAME = ? AND COLUMN_NAME = ?`, tableName, columnName)

	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: query column error: %s", err.Error()))
	}

	return count > 0
}

func (ss SqlStore) CreateColumnIfNotExists(tableName string, columnName string, colType string, defaultValue string) bool {
	if ss.DoesColumnExist(tableName, columnName) {
		return false
	}

	_, err := ss.GetMaster().Exec("ALTER TABLE " + tableName + " ADD " + columnName + " " + colType + " DEFAULT '" + defaultValue + "'")
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: create column error: %s", err.Error()))
	}

	return true
}

func (ss SqlStore) RemoveColumnIfExists(tableName string, columnName string) bool {
	if !ss.DoesColumnExist(tableName, columnName) {
		return false
	}

	_, err := ss.GetMaster().Exec("Alter TABLE " + tableName + " DROP COLUMN " + columnName)
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: dropping column error: %s", err.Error()))
	}

	return true
}

func (ss SqlStore) RenameColumnIfExists(tableName string, oldColumnName string, newColumnName string, colType string) bool {
	if !ss.DoesColumnExist(tableName, oldColumnName) {
		return false
	}

	_, err := ss.GetMaster().Exec("ALTER TABLE " + tableName + " CHANGE " + oldColumnName + " " + newColumnName + " " + colType)
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: rename column error: %s", err.Error()))
	}

	return true
}

func (ss SqlStore) GetMaxLengthOfColumnIfExists(tableName string, columnName string) string {
	if !ss.DoesColumnExist(tableName, columnName) {
		return ""
	}

	result, err := ss.GetMaster().SelectStr("SELECT CHARACTER_MAXIMUM_LENGTH FROM information_schema.columns WHERE table_name = '" + tableName + "' AND COLUMN_NAME = '" + columnName + "'")

	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: get max column length error: %s", err.Error()))
	}

	return result
}

func (ss SqlStore) AlterColumnTypeIfExists(tableName string, columnName string, mySqlColType string, postgresColType string) bool {
	if !ss.DoesColumnExist(tableName, columnName) {
		return false
	}

	_, err := ss.GetMaster().Exec("ALTER TABLE " + tableName + " MODIFY " + columnName + " " + mySqlColType)

	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: alter column type error: %s", err.Error()))
	}

	return true
}

func (ss SqlStore) CreateUniqueIndexIfNotExists(indexName string, tableName string, columnName string) {
	ss.createIndexIfNotExists(indexName, tableName, columnName, INDEX_TYPE_DEFAULT, true)
}

func (ss SqlStore) CreateIndexIfNotExists(indexName string, tableName string, columnName string) {
	ss.createIndexIfNotExists(indexName, tableName, columnName, INDEX_TYPE_DEFAULT, false)
}

func (ss SqlStore) CreateFullTextIndexIfNotExists(indexName string, tableName string, columnName string) {
	ss.createIndexIfNotExists(indexName, tableName, columnName, INDEX_TYPE_FULL_TEXT, false)
}

func (ss SqlStore) createIndexIfNotExists(indexName string, tableName string, columnName string, indexType string, unique bool) {

	uniqueStr := ""
	if unique {
		uniqueStr = "UNIQUE "
	}

	count, err := ss.GetMaster().SelectInt("SELECT COUNT(0) AS index_exists FROM information_schema.statistics WHERE TABLE_SCHEMA = DATABASE() and table_name = ? AND index_name = ?", tableName, indexName)
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: check index exists error: %s", err.Error()))
	}

	if count > 0 {
		return
	}

	fullTextIndex := ""
	if indexType == INDEX_TYPE_FULL_TEXT {
		fullTextIndex = " FULLTEXT "
	}

	_, err = ss.GetMaster().Exec("CREATE  " + uniqueStr + fullTextIndex + " INDEX " + indexName + " ON " + tableName + " (" + columnName + ")")
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: create index error: %s", err.Error()))
	}
}

func (ss SqlStore) RemoveIndexIfExists(indexName string, tableName string) {
	count, err := ss.GetMaster().SelectInt("SELECT COUNT(0) AS index_exists FROM information_schema.statistics WHERE TABLE_SCHEMA = DATABASE() and table_name = ? AND index_name = ?", tableName, indexName)
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: check index exists error: %s", err.Error()))
	}

	if count > 0 {
		return
	}

	_, err = ss.GetMaster().Exec("DROP INDEX " + indexName + " ON " + tableName)
	if err != nil {
		panic(fmt.Sprintf("CRITICAL ERROR: remove index error: %s", err.Error()))
	}
}

func IsUniqueConstraintError(err string, indexName []string) bool {
	unique := strings.Contains(err, "unique constraint") || strings.Contains(err, "Duplicate entry")
	field := false
	for _, contain := range indexName {
		if strings.Contains(err, contain) {
			field = true
			break
		}
	}

	return unique && field
}

func (ss SqlStore) GetMaster() *gorp.DbMap {
	return ss.master
}

func (ss SqlStore) GetAllConns() []*gorp.DbMap {
	all := make([]*gorp.DbMap, len(ss.replicas)+1)
	copy(all, ss.replicas)
	all[len(ss.replicas)] = ss.master
	return all
}

func (ss SqlStore) Close() {
	l4g.Info(utils.T("store.sql.closing.info"))
	ss.master.Db.Close()
	for _, replica := range ss.replicas {
		replica.Db.Close()
	}
}

func (ss SqlStore) Game() GameStore {
	return ss.game
}

func (ss SqlStore) Player() PlayerStore {
	return ss.player
}

func (ss SqlStore) Move() MoveStore {
	return ss.move
}

type g3Converter struct{}

func (g g3Converter) ToDb(val interface{}) (interface{}, error) {
	return val, nil
}

func (g g3Converter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	return gorp.CustomScanner{}, false
}
