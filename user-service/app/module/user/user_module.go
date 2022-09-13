package user

type UserModule struct {
	UserRPC        *UserRPC
	UserService    *UserService
	UserRepository *UserRepository
}

func (m *UserModule) InitComponents() {
	m.UserRepository = new(UserRepository)
	m.UserService = NewUserService(
		m.UserRepository,
	)
	m.UserRPC = NewUserRPC(
		m.UserService,
	)
}
