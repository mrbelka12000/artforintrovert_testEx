package models

type Store interface {
	Update(product *Product) error
	GetAll() ([]Product, error)
	Delete(id string) error
	Insert() error
}
