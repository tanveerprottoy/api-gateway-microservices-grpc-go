package user

type UserModule struct {
	UserHandler *UserHandler
	UserService *UserService
}

func (m *UserModule) InitComponents() {
	m.UserService = new(UserService)
	m.UserHandler = NewUserHandler(
		m.UserService,
	)
}
