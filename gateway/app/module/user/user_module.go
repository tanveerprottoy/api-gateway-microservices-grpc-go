package user

type UserModule struct {
	UserHandler *UserHandler
	UserService *UserService
}

func (m *UserModule) InitComponents() {
	m.UserService = &UserService{}
	m.UserHandler = &UserHandler{
		m.UserService,
	}
}
