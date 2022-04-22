package golangtbotapi

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/udev-21/golang-tbot-api/types"
)

type BotAPI struct {
	httpClient *http.Client
	token      string
	debug      bool
	logger     log.Logger
	privateKey rsa.PrivateKey
	handlers   []HandlerFunc
}

func NewBotAPI(token string) *BotAPI {
	res := BotAPI{
		httpClient: &http.Client{Timeout: 3 * time.Second},
		token:      token,
	}
	return &res
}

func (ba *BotAPI) WithPrivateKey(privKey rsa.PrivateKey) *BotAPI {
	ba.privateKey = privKey
	return ba
}

func (ba *BotAPI) ProcessUpdate(update types.Update) {
	for _, handle := range ba.handlers {
		handle(&nativeContext{
			update: update,
			bot:    ba,
		})
	}
}

func (ba *BotAPI) Handle(condition Middleware, handler HandlerFunc, additionalMiddlewares ...Middleware) {
	handler = applyMiddlewares(handler, additionalMiddlewares...)
	handler = condition(handler)
	ba.handlers = append(ba.handlers, handler)
}

func (ba *BotAPI) Send(reciever types.Chat, payload MessagePayload) (*types.ApiResponse, error) {
	if _, ok := payload.(UploadWithFiles); !ok {
		data, err := payload.RawJsonPayload()
		if err != nil {
			return nil, err
		}
		if reciever.ID != 0 {
			data["chat_id"] = reciever.ID
		} else {
			data["chat_id"] = reciever.Username
		}

		p, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		log.Println(string(p))
		body := bytes.NewBuffer(p)
		log.Println(ba.token, payload)
		log.Println(fmt.Sprintf("https://api.telegram.org/bot%s/%s", ba.token, payload.GetEndpoint()))
		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/%s", ba.token, payload.GetEndpoint()), body)

		if err != nil {
			log.Println("papalsyas")
			panic(err)
		}

		request.Header.Set("Content-Type", "application/json")

		var res types.ApiResponse

		response, err := ba.httpClient.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		json.NewDecoder(response.Body).Decode(&res)
		return &res, nil
	}
	return nil, nil
}
