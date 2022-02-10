package kagesolutionsmcgo

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// Launch client to connect to ws
func LaunchClient(host, key string) {
	for {
		u := url.URL{Scheme: "wss", Host: host, Path: "/"}
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Println("[ERROR CONNECTING TO TOKENS SOCKET]", err.Error())
		} else {
			ConnSocket = conn
			ConnSocketMu.Lock()
			ConnSocket.WriteJSON(RequestToken{
				Message: "activate",
				Key:     key,
			})
			ConnSocketMu.Unlock()

			for {
				// receive message
				_, message, err := conn.ReadMessage()

				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Println("[WS] Closing Tokens Connection." + err.Error())
					}

					ConnSocket = nil
					break
				}

				TokenMessage := TokenResponse{}
				json.Unmarshal(message, &TokenMessage)
				switch TokenMessage.MessageType {
				case "success":
					log.Println("[WEBSOCKET SERVER]", TokenMessage.Message)
				case "token-fail":
					TokensMu.Lock()
					Tokens[TokenMessage.Captcha.Task] = "dev"
					TokensMu.Unlock()
				case "token-success":
					TokensMu.Lock()
					Tokens[TokenMessage.Captcha.Task] = TokenMessage.Captcha.Token
					TokensMu.Unlock()
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

type TokenResponse struct {
	MessageType string `json:"messagetype"`
	Message     string `json:"message"`
	Captcha     struct {
		Token string `json:"token"`
		Task  string `json:"task"`
		Site  string `json:"site"`
	}
}

type RequestToken struct {
	Message string `json:"message"`
	Task    string `json:"task"`
	Site    string `json:"site"`
	Key     string `json:"key"`
}

func RemoveToken(key string) {
	TokensMu.Lock()
	defer TokensMu.Unlock()
	delete(Tokens, key)
}

func CheckToken(taskid string) (string, bool) {
	TokensMu.Lock()
	defer TokensMu.Unlock()
	if _, ok := Tokens[taskid]; ok {
		token := Tokens[taskid]
		delete(Tokens, taskid)
		return token, true
	}
	return "", false
}

func TokenRequest(request RequestToken) error {
	ConnSocketMu.Lock()
	defer ConnSocketMu.Unlock()
	if ConnSocket == nil {
		return errors.New("socket connection is not configured correctly")
	}
	ConnSocket.WriteJSON(request)
	return nil
}
