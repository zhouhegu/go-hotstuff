package hotstuff

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

func OpenDB(path string) (*leveldb.DB, error) {
	return leveldb.OpenFile(path, nil)
}

// 连接leveldb数据库
func NewMemDB() *leveldb.DB {
	db, err := leveldb.Open(storage.NewMemStorage(), nil)
	if err != nil {
		panic(err.Error())
	}
	return db
}
