package conn

import "github.com/jmcarbo/surrealdb.go/pkg/model"

type Connection interface {
	Connect(url string) (Connection, error)
	Send(method string, params []interface{}) (interface{}, error)
	Close() error
	LiveNotifications(id string) (chan model.Notification, error)
}
