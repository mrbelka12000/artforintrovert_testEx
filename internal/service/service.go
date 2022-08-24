package service

type Service struct {
	*productRepo
}

func NewService(repo ProductStoreRepo) *Service {
	return &Service{
		productRepo: newProduct(repo),
	}
}
