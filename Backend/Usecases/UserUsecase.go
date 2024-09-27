package usecases

import domain "unique-minds/Domain"

type UserUseCase struct {
	userRepo domain.UserRepoInterface
}


func NewUserUseCase(repo domain.UserRepoInterface) *UserUseCase {
	return &UserUseCase{
		userRepo : repo,
	}
}