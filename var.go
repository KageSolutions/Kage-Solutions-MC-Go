package kagesolutionsmcgo

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	/* Frontend websocket connection */
	ConnSocketMu sync.RWMutex
	ConnSocket   *websocket.Conn
)

var TokensMu sync.Mutex
var Tokens = make(map[string]string)
