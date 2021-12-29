package user

type User struct {
	id   string
	name string
}

func NewUser(name string) User {
	return User{
		id:   "", // TODO ULID
		name: name,
	}
}

// TODO Reconstruct
