package conn

import (
	"errors"
	"github.com/jmcarbo/surrealdb.go/pkg/model"
)

var (
	ErrTimeout           = errors.New("timeout")
	ErrInvalidResponseID = errors.New("invalid response id")
)

type Connection interface {
	Connect(url string) (Connection, error)
	Send(method string, params []interface{}) (interface{}, error)
	Close() error
	LiveNotifications(id string) (chan model.Notification, error)
}
