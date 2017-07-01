package domain

// User at domain level, he only has an address and a way to check it
type User struct {
	Address string
	UserAddressChecker
}

// NewUser : domain.User constructor
func NewUser(a string, c UserAddressChecker) User {
	return User{a, c}
}

// UserAddressChecker is defined at interfaces layer
// check interfaces/dummys/DummyAddressChecker
// or interfaces/addressChecker for a "real" checker
type UserAddressChecker interface {
	CheckAddress(User) bool
}
