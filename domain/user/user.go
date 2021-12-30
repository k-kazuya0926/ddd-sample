package user

// エンティティ
type User struct {
	id   string
	name UserName
}

// UseCase用
func NewUser(name UserName) User {
	return User{
		id:   "", // TODO ULID
		name: name,
	}
}

// DBからの再構築用
func ReconstructUser(id string, name UserName) User {
	return User{
		id:   id,
		name: name,
	}
}

func (u *User) SetName(name UserName) {
	u.name = name
}
