package user

import (
	domain "ddd-sample/domain/user"
)

type UserDataModelBuilder struct {
	id   domain.UserID
	name domain.UserName
}

func NewUserDataModelBuilder() UserDataModelBuilder {
	return UserDataModelBuilder{}
}

func (b *UserDataModelBuilder) ID(id domain.UserID) {
	b.id = id
}

func (b *UserDataModelBuilder) Name(name domain.UserName) {
	b.name = name
}

// 通知されたデータからデータモデルを生成するメソッド
func (b *UserDataModelBuilder) Build() UserDataModel {
	return UserDataModel{
		ID:   b.id,
		Name: b.name,
	}
}
