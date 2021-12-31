package user

// エンティティ
type User struct {
	// 属性はパッケージプライベート
	id   UserID
	name UserName
}

// リポジトリ経由の再構築およびテストコード用コンストラクタ
func ReconstructUser(id UserID, name UserName) User {
	return User{
		id:   id,
		name: name,
	}
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) SetName(name UserName) {
	u.name = name
}

func (u *User) Notify(note UserNotification) {
	// 内部データを通知
	note.ID(u.id)
	note.Name(u.name)
}
