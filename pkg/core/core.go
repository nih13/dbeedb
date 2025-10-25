package core

import "fmt"

// Database is an interface that defines methods for getting and setting a value.

// type Database interface {
// 	Get(key string) string
// 	Set(key string, value string)
// }

type dbeedb struct {
	db       map[string]string
	loglevel string
}

func (d dbeedb) Get(key string) string {
	return d.db[key]
}

func (d dbeedb) Set(key string, value string) {
	d.db[key] = value
}

func New() dbeedb {
	abc := dbeedb{db: map[string]string{}, loglevel: "debug"}
	return abc
}

func (d dbeedb) Print() {

	fmt.Println(d.db)
}

func (d dbeedb) Delete(key string) {
	delete(d.db, key)
}

//["a":"apple", "b":"mango"]
