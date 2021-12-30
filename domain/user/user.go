package user

type User struct {
	id   string
	name string
}

// UseCase用
func NewUser(name string) (User, error) {
	// TODO validation
	user := User{
		id: "", // TODO ULID
	}
	err := user.SetName(name)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// DBからの再構築用
func ReconstructUser(id string, name string) User {
	// ここはバリデーション不要
	return User{
		id:   id,
		name: name,
	}
}

func (u *User) SetName(name string) error {
	// TODO validation
	u.name = name
	return nil
}
