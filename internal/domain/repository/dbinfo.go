package repository

type DBinfo interface {
	Ping() error
}
