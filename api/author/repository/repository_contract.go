package repository

type RepositoryContract interface {
	AuthorContract
}

type Repository struct{}

func NewRepository() RepositoryContract {
	return &Repository{}
}
