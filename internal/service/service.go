package service

type Service struct {
	Product
}

func NewService(repo ProductStoreRepo) *Service {
	return &Service{
		Product: *NewProduct(repo),
	}
}
