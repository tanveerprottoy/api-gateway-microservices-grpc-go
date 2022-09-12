package user

type UserModule struct {
	UserRepository *UserRepository
	UserService    *UserService
}

func (m *UserModule) InitComponents() {
	m.UserRepository = &UserRepository{}
	m.UserService = &UserService{
		m.UserRepository,
	}
}
