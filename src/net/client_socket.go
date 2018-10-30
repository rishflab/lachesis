package net

// A listener 
type MessageListener interface {
	OnMessage(msg []byte, err error)
}

// ClientSocket represents a client socket connection
type ClientSocket interface {
	Connect(addr string, path string) error
	AddMessageListener(listener MessageListener)
	SendMessage(msg []byte) error
	Stop() error
}
