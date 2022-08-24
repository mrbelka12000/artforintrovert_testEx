package models

type ProductStore interface {
	Update(product *Product) error
	GetAll() ([]Product, error)
	Delete(id string) error
	Insert() error
}
