package token

type IAccessToken interface {
	GetAccessToken() (token string, err error)
}

type DefaultTokenManager struct {
}

func (d *DefaultTokenManager) GetAccessToken() (token string, err error) {
	//TODO implement me
	panic("implement me")
}
