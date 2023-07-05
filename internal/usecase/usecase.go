package usecase

import "github.com/Rosto4eks/eclipse/internal/database"

type Iusecase interface {
}

type usecase struct {
	database database.Idatabase
}

func New(database database.Idatabase) *usecase {
	return &usecase{
		database,
	}
}
