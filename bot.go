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
	telegramAPIUrl string
	poller         Poller

	httpClient *http.Client
	token      string
	debug      bool
	logger     log.Logger
	privateKey rsa.PrivateKey

	updates chan types.Update

	handlers []HandlerFunc
	stop     chan chan struct{}
}

func NewBotAPI(token string) *BotAPI {

	res := BotAPI{
		httpClient: &http.Client{Timeout: 3 * time.Second},
		token:      token,
		poller:     NewLongPoller(),
		updates:    make(chan types.Update, 200),
	}
	res.telegramAPIUrl = fmt.Sprintf("https://api.telegram.org/bot%s/", token)
	return &res
}

func (ba *BotAPI) WithPrivateKey(privKey rsa.PrivateKey) *BotAPI {
	ba.privateKey = privKey
	return ba
}

func (ba *BotAPI) GetUpdates(payload MessagePayload) ([]types.Update, error) {

	bodyMap, err := payload.RawJsonPayload()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, ba.telegramAPIUrl+payload.GetEndpoint(), bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("prepare request error")
	}

	request.Header.Set("Content-Type", "application/json")

	var res struct {
		OK      bool           `json:"ok"`
		Updates []types.Update `json:"result"`
	}

	response, err := ba.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request error")
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("parse response error")
	}
	if !res.OK {
		return nil, fmt.Errorf("telegram api response error")
	}
	return res.Updates, nil
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
		body := bytes.NewBuffer(p)
		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf(ba.telegramAPIUrl+"%s", payload.GetEndpoint()), body)

		if err != nil {
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

func (b *BotAPI) Start() {
	if b.poller == nil {
		panic("telebot: can't start without a poller")
	}
	log.Println("starting bot long polling")
	stop := make(chan struct{})
	stopConfirm := make(chan struct{})

	go func() {
		log.Println("poller gone")
		b.poller.Poll(b, b.updates, stop)
		close(stopConfirm)
	}()

	log.Println("after poller")
	// time.Sleep(2 * time.Second)

	for {
		select {
		// handle incoming updates
		case upd := <-b.updates:
			b.ProcessUpdate(upd)
			// call to stop polling
		case confirm := <-b.stop:
			close(stop)
			<-stopConfirm
			close(confirm)
			return
		}
	}
}
