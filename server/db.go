package server

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type MyDB struct {
	KeyValue map[string]string
	mu       sync.RWMutex
}

func NewDB() MyDB {

	f, err := os.Open("backup.json")
	if err != nil {
		return MyDB{KeyValue: map[string]string{}}
	}
	items := make(map[string]string)
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		fmt.Println(err)
		return MyDB{KeyValue: map[string]string{}}

	}

	return MyDB{KeyValue: items}
}

func (db *MyDB) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.KeyValue[key] = value
}
func (db *MyDB) Get(key string) (string, bool) {
	value, ok := db.KeyValue[key]
	return value, ok
}
func (db *MyDB) Del(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.KeyValue, key)
}

func (db *MyDB) Save() {
	f, err := os.Create("backup.json")
	if err != nil {
		fmt.Println("Error in saving backup. :", err)
	} else {
		json.NewEncoder(f).Encode(&db.KeyValue)
		f.Close()
	}
}
