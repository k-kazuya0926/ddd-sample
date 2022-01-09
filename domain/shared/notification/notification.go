package notification

import "ddd-sample/domain/user"

type Notification struct {
	TargetUserID user.UserID
	Message      string
}

type NotificationClient interface {
	Notify(notification Notification) error
}
