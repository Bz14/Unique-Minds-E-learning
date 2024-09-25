package usecases

type UserUseCase struct{}

type UserUseCaseInterface interface{}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}