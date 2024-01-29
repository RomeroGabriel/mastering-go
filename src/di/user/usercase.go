package user

type UserUseCase struct {
	repository *UserRepository
}

func NewUserUseCase(r *UserRepository) *UserUseCase {
	return &UserUseCase{r}
}

func (u *UserUseCase) GetUser(id int) (*User, error) {
	return u.repository.GetUser(id)
}
