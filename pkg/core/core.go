package core

// Database is an interface that defines methods for getting and setting a value.

type Database interface {
	Get(key string) string
	Set(key string, value string)
}

type dbeedb struct {
	db map[string]string
}

func (d dbeedb) Get(key string) string {
	return d.db[key]
}

func (d dbeedb) Set(key string, value string) {
	d.db[key] = value
}

func New() Database {
	abc := dbeedb{db: map[string]string{}}

	return abc
}
