package domain

type UserClient struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  []byte
}

type UserDriver struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	CarModel  string
}
