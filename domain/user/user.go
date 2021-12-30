package user

// エンティティ
type User struct {
	// 属性はパッケージプライベート
	id   string
	name UserName
}

// リポジトリ経由の再構築およびテストコード用コンストラクタ
func ReconstructUser(id string, name UserName) User {
	return User{
		id:   id,
		name: name,
	}
}

func (u *User) SetName(name UserName) {
	u.name = name
}
