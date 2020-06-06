package tinyauth

// Authenticate authenticates a user given the identifier and an optional password.
// Password check is only done if `require_password` is false for the user.
func (t *TinyAuth) Authenticate(identifier, password string) (User, error) {
	return User{}, nil
}

// Authenticate authenticates a user given the identifier and an optional password.
// Password check is only done if `require_password` is false for the user.
func (t *TinyAuth) CreateSession(identifier, password string) (User, error) {
	return User{}, nil
}
