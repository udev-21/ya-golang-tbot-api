package yagolangtbotapi

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

type BotAPI struct {
	telegramAPIUrl string
	poller         Poller

	httpClient      *http.Client
	token           string
	debug           bool
	logger          *log.Logger
	privateKey      rsa.PrivateKey
	testEnvironment bool

	updates chan types.Update

	handlers []HandlerFunc
	stop     chan chan struct{}
}

func NewBotAPI(token string) *BotAPI {

	res := BotAPI{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		token:      token,
		poller:     NewLongPoller(),
		updates:    make(chan types.Update, 200),
	}
	res.telegramAPIUrl = "https://api.telegram.org/"
	return &res
}

func (ba *BotAPI) TestEnvironment() {
	ba.testEnvironment = true
}

func (ba *BotAPI) WithPrivateKey(privKey rsa.PrivateKey) *BotAPI {
	ba.privateKey = privKey
	return ba
}
func (ba *BotAPI) WithLogger(logger *log.Logger) *BotAPI {
	ba.logger = logger
	return ba
}

func (ba *BotAPI) GetUpdates(payload myTypes.Sendable) ([]types.Update, error) {

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, ba.getPath(payload.Endpoint()), bytes.NewBuffer(body))

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
		return nil, fmt.Errorf("parse response error: " + err.Error())
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

func request(endpoint string, params myTypes.Params, client *http.Client) ([]byte, error) {
	p, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(p)

	request, err := http.NewRequest(http.MethodPost, endpoint, body)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ba *BotAPI) request(endpoint string, params myTypes.Params) (*types.ApiResponse, error) {
	body, err := request(ba.getPath(endpoint), params, ba.httpClient)
	if err != nil {
		writeLog(LogLevelError, ba.logger, "error while sending request: %v", err)
		return nil, err
	}
	if ba.debug {
		writeLog(LogLevelInfo, ba.logger, "recieved response from api: %q", string(body))
	}

	var res types.ApiResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		writeLog(LogLevelError, ba.logger, "error while unmarshal response: %v", err)
		return nil, err
	}

	return &res, nil
}

func hasFilesNeedingUpload(files []myTypes.InputFile) bool {
	for _, file := range files {
		if _, ok := file.(myTypes.Uploadable); ok {
			return true
		}
	}
	return false
}

func uploadFile(r *io.PipeReader, w *io.PipeWriter, uploadable myTypes.Uploadable, m *multipart.Writer) {
	reader, err := uploadable.UploadData()
	if err != nil {
		w.CloseWithError(err)
		return
	}
	part, err := m.CreateFormFile(uploadable.Field(), uploadable.FileName())
	if err != nil {
		w.CloseWithError(err)
		return
	}

	if _, err := io.Copy(part, reader); err != nil {
		w.CloseWithError(err)
		return
	}

	if closer, ok := reader.(io.ReadCloser); ok {
		if err = closer.Close(); err != nil {
			w.CloseWithError(err)
			return
		}
	}
}

func (ba *BotAPI) requestWithFiles(reciever interface{}, endpoint string, sendable myTypes.UploadWithFiles) (*types.ApiResponse, error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	if ba.debug {
		writeLog(LogLevelInfo, ba.logger, "uploading with files")
	}

	params, err := sendable.Params()
	if err != nil {
		return nil, err
	}

	var chatID interface{}

	if chat, ok := reciever.(types.Chat); ok {
		if chat.ID != 0 {
			chatID = chat.ID
		} else {
			chatID = chat.Username
		}
	} else if chat, ok := reciever.(*types.Chat); ok {
		if chat.ID != 0 {
			chatID = chat.ID
		} else {
			chatID = chat.Username
		}
	}
	if chatID != nil {
		params["chat_id"] = chatID
	}

	files := sendable.Files()

	go func() {
		defer w.Close()
		defer m.Close()

		for field, value := range params {

			if _, ok := value.(myTypes.Uploadable); ok {
				continue
			}

			if strval, ok := value.(string); ok {
				if err := m.WriteField(field, strval); err != nil {
					w.CloseWithError(err)
					return
				}
			} else if strval, ok := value.(*string); ok {
				if err := m.WriteField(field, *strval); err != nil {
					w.CloseWithError(err)
					return
				}
			} else {
				val, err := json.Marshal(value)
				log.Printf("field(%q) value(%s)", field, val)
				if err != nil {
					w.CloseWithError(err)
					return
				}
				if err := m.WriteField(field, string(val)); err != nil {
					w.CloseWithError(err)
					return
				}
			}
		}

		for _, file := range files {
			if f, ok := file.(myTypes.Uploadable); ok {
				uploadFile(r, w, f, m)
			}
		}
	}()

	req, err := http.NewRequest("POST", ba.getPath(endpoint), r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", m.FormDataContentType())

	response, err := ba.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	defer response.Body.Close()
	var res types.ApiResponse
	json.NewDecoder(response.Body).Decode(&res)
	return &res, nil
}

func (ba *BotAPI) Send(reciever interface{}, payload myTypes.Sendable) (*types.ApiResponse, error) {
	data, err := payload.Params()

	if err != nil {
		return nil, err
	}

	if ba.debug {
		writeLog(LogLevelInfo, ba.logger, "new request to api with endpoint: %q", payload.Endpoint())
	}

	var chatID interface{}

	if chat, ok := reciever.(types.Chat); ok {
		if chat.ID != 0 {
			chatID = chat.ID
		} else {
			chatID = chat.Username
		}
	} else if chat, ok := reciever.(*types.Chat); ok {
		if chat.ID != 0 {
			chatID = chat.ID
		} else {
			chatID = chat.Username
		}
	}
	if chatID != nil {
		if ba.debug {
			writeLog(LogLevelInfo, ba.logger, "chat_id set: %v", chatID)
		}
		data["chat_id"] = chatID
	}

	if payloadWithFiles, ok := payload.(myTypes.UploadWithFiles); ok {
		if hasFilesNeedingUpload(payloadWithFiles.Files()) {
			return ba.requestWithFiles(reciever, payload.Endpoint(), payloadWithFiles)
		}
	}

	return ba.request(payload.Endpoint(), data)

}

func (b *BotAPI) getPath(endpoint string) string {
	middle := "/"
	if b.testEnvironment {
		middle = "/test/"
	}
	log.Println(b.telegramAPIUrl + "bot" + b.token + middle + endpoint)
	return b.telegramAPIUrl + "bot" + b.token + middle + endpoint
}

func (b *BotAPI) getFilePath(filePath string) string {
	middle := "/"
	if b.testEnvironment {
		middle = "/test/"
	}
	return b.telegramAPIUrl + "file/bot" + b.token + middle + filePath
}

func (b *BotAPI) Start() {
	if b.poller == nil {
		panic("golangtbotapi: can't start without a poller")
	}

	writeLog(LogLevelInfo, b.logger, "starting bot long polling")
	stop := make(chan struct{})
	stopConfirm := make(chan struct{})

	go func() {
		log.Println("poller gone")
		b.poller.Poll(b, b.updates, stop)
		close(stopConfirm)
	}()

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
