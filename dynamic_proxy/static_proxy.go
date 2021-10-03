package dynamic_proxy

type IUser interface {
	Login(username, password string) error
}

type User struct {
	
}

func (u *User) Login(username, password string) error {
	return nil
}

// UserProxy 代理类
type UserProxy struct {
	user *User
}

// NewUserProxy New UserProxy
func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

func (p *UserProxy) Login(username, password string) error {
	if err := p.user.Login(username, password); err != nil {
		return err
	}
	return nil
}
