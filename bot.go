package golangtbotapi

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/udev21/golang-tbot-api/types"
)

type BotAPI struct {
	httpClient *http.Client
	token      string
	debug      bool
	logger     log.Logger
	privateKey rsa.PrivateKey
}

func NewBotAPI(token string) *BotAPI {
	return &BotAPI{
		httpClient: &http.Client{Timeout: 3 * time.Second},
		token:      token,
	}
}

func (ba *BotAPI) WithPrivateKey(privKey rsa.PrivateKey) *BotAPI {
	ba.privateKey = privKey
	return ba
}

func (ba *BotAPI) Send(reciever types.Chat, payload MessagePayload) (interface{}, error) {
	if _, ok := payload.(UploadFiler); !ok {
		p, err := payload.GetFieldsJsonPayload(reciever.ID)
		if err != nil {
			panic(err)
		}
		log.Println(string(p))
		body := bytes.NewBuffer(p)

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/%s", ba.token, payload.GetEndpoint()), body)

		if err != nil {
			panic(err)
		}

		request.Header.Set("Content-Type", "application/json")

		var res map[string]interface{}

		response, err := ba.httpClient.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		json.NewDecoder(response.Body).Decode(&res)
		return res, nil
	}
	return nil, nil
}
