package domain

type User struct {
	Id   UserId
	Name UserName
}

type UserName string

func (n UserName) String() string {
	return string(n)
}

type UserId string

func (i UserId) String() string {
	return string(i)
}

type UserNotFound error
