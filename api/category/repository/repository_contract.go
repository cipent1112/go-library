package repository

type RepositoryContract interface {
	CategoryContract
}

type Repository struct{}

func NewRepository() RepositoryContract {
	return &Repository{}
}
