package server

import "github.com/shenzh1990/file-sync/model"

type NetConn interface {
	// HandlerLoop 不能阻塞
	HandlerLoop()
	GetMsg() (*model.Message, bool)
	SendMsg(m *model.Message)
	Close() error
	IsClose() bool
}
