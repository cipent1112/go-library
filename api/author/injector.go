package author

import (
	"go-library/api/author/gateway/handler"
	"go-library/api/author/repository"
	"go-library/api/author/usecase"
	"go-library/config"
)

func ApiHandler() handler.HandlerContract {
	repo := repository.NewRepository()
	uc := usecase.NewUsecase(config.Conf.DBConfig.Orm, repo)
	return handler.NewHandler(uc)
}
