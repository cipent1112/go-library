package category

import (
	"go-library/api/category/gateway/handler"
	"go-library/api/category/repository"
	"go-library/api/category/usecase"
	"go-library/config"
)

func ApiHandler() handler.HandlerContract {
	repo := repository.NewRepository()
	uc := usecase.NewUsecase(config.Conf.DBConfig.Orm, repo)
	return handler.NewHandler(uc)
}
