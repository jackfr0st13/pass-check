package pass

import "errors"

type Origin string

const (
	USER      Origin = "User"
	GENERATED Origin = "Generated"
)

type PasswordData struct {
	password string
	username string
	origin   Origin
}

type PasswordDataOption func(*PasswordData)

func NewPasswordData(opts ...PasswordDataOption) *PasswordData {
	passData := PasswordData{
		origin: USER,
	}
	for _, o := range opts {
		o(&passData)
	}
	return &passData
}

func WithPassword(password string) PasswordDataOption {
	return func(c *PasswordData) {
		c.password = password
	}
}

func WithUsername(username string) PasswordDataOption {
	return func(c *PasswordData) {
		c.username = username
	}
}

func (pd *PasswordData) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password cannot be empty")
	}
	pd.password = password
	return nil
}

func (pd *PasswordData) GetPassword() string {
	return pd.password
}

func (pd *PasswordData) SetUsername(username string) error {
	if len(username) == 0 {
		return errors.New("username cannot be empty")
	}
	pd.username = username
	return nil
}

func (pd *PasswordData) GetUsername() string {
	return pd.username
}

func (pd *PasswordData) SetOrigin(origin Origin) error {
	if !(origin == USER || origin == GENERATED) {
		return errors.New("invalid value for origin")
	}
	pd.origin = origin
	return nil
}

func (pd *PasswordData) GetOrigin() Origin {
	return pd.origin
}
