package user

type UserModule struct {
	UserHandler *UserHandler
	Uservice    *UserService
}

func (m *UserModule) InitComponents() {
	m.Uservice = &UserService{}
	m.UserHandler = &UserHandler{
		m.Uservice,
	}
}
