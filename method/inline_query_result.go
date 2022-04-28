package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

type InlineQueryResultOLD struct {
	Type                string                      `json:"type"`
	ID                  string                      `json:"id"`
	Title               *string                     `json:"title,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
	ReplyMarkup         *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 *string                     `json:"url,omitempty"`
	HideUrl             *bool                       `json:"hide_url,omitempty"`
	Description         *string                     `json:"description,omitempty"`
	ThumbUrl            *string                     `json:"thumb_url,omitempty"`
	ThumbWidth          *int64                      `json:"thumb_width,omitempty"`
	ThumbHeight         *int64                      `json:"thumb_height,omitempty"`

	PhotoUrl        *string               `json:"photo_url,omitempty"`
	PhotoWidth      *int64                `json:"photo_width,omitempty"`
	PhotoHeight     *int64                `json:"photo_height,omitempty"`
	Caption         *string               `json:"caption,omitempty"`
	ParseMode       *string               `json:"parse_mode,omitempty"`
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`

	GifUrl        *string `json:"gif_url,omitempty"`
	GifWidth      *int64  `json:"gif_width,omitempty"`
	GifHeight     *int64  `json:"gif_height,omitempty"`
	GifDuration   *int64  `json:"gif_duration,omitempty"`
	ThumbMimeType *string `json:"thumb_mime_type,omitempty"`

	Mpeg4Url      *string `json:"mpeg4_url,omitempty"`
	Mpeg4Width    *int64  `json:"mpeg4_width,omitempty"`
	Mpeg4Height   *int64  `json:"mpeg4_height,omitempty"`
	Mpeg4Duration *int64  `json:"mpeg4_duration,omitempty"`

	VideoUrl      *string `json:"video_url,omitempty"`
	MimeType      *string `json:"mime_type,omitempty"`
	VideoWidth    *int64  `json:"video_width,omitempty"`
	VideoHeight   *int64  `json:"video_height,omitempty"`
	VideoDuration *int64  `json:"video_duration,omitempty"`

	AudioUrl      *string `json:"audio_url,omitempty"`
	Performer     *string `json:"performer,omitempty"`
	AudioDuration *int64  `json:"audio_duration,omitempty"`

	VoiceUrl      *string `json:"voice_url,omitempty"`
	VoiceDuration *int64  `json:"voice_duration,omitempty"`

	DocumentUrl *string `json:"document_url,omitempty"`

	Latitude             *float64 `json:"latitude,omitempty"`
	Longitude            *float64 `json:"longitude,omitempty"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int64   `json:"live_period,omitempty"`
	Heading              *int64   `json:"heading,omitempty"`
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"`

	Address         *string `json:"address,omitempty"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`

	PhoneNumber *string `json:"phone_number,omitempty"`
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
	VCard       *string `json:"vcard,omitempty"`

	GameShortName *string `json:"game_short_name,omitempty"`

	PhotoFileID    *string `json:"photo_file_id,omitempty"`
	GifFileID      *string `json:"gif_file_id,omitempty"`
	Mpeg4FileID    *string `json:"mpeg4_file_id,omitempty"`
	StickerFileID  *string `json:"sticker_file_id,omitempty"`
	DocumentFileID *string `json:"document_file_id,omitempty"`
	VideoFileID    *string `json:"video_file_id,omitempty"`
	VoiceFileID    *string `json:"voice_file_id,omitempty"`
	AudioFileID    *string `json:"audio_file_id,omitempty"`
}

const (
	InlineQueryResultArticleType  = "article"
	InlineQueryResultPhotoType    = "photo"
	InlineQueryResultGifType      = "gif"
	InlineQueryResultMpeg4GifType = "mpeg4_gif"
	InlineQueryResultVideoType    = "video"
	InlineQueryResultAudioType    = "audio"
	InlineQueryResultVoiceType    = "voice"
	InlineQueryResultDocumentType = "document"
	InlineQueryResultLocationType = "location"
	InlineQueryResultVenueType    = "venue"
	InlineQueryResultContactType  = "contact"
	InlineQueryResultGameType     = "game"

	InlineQueryResultCachedPhotoType    = InlineQueryResultPhotoType
	InlineQueryResultCachedGifType      = InlineQueryResultGifType
	InlineQueryResultCachedMpeg4GifType = InlineQueryResultMpeg4GifType
	InlineQueryResultCachedStickerType  = "sticker"
	InlineQueryResultCachedDocumentType = InlineQueryResultDocumentType
	InlineQueryResultCachedVideoType    = InlineQueryResultVideoType
	InlineQueryResultCachedVoiceType    = InlineQueryResultVoiceType
	InlineQueryResultCachedAudioType    = InlineQueryResultAudioType
)

type InlineQueryResult interface {
	IsInlineQueryResult()
}

type BaseInlineQueryResult struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (biqr *BaseInlineQueryResult) IsInlineQueryResult() {}

type ThumbUrler struct {
	ThumbUrl string `json:"thumb_url,omitempty"`
}

type ThumbWidther struct {
	ThumbWidth int64 `json:"thumb_width,omitempty"`
}

type ThumbHeighter struct {
	ThumbHeight int64 `json:"thumb_height,omitempty"`
}

func (tuwh *ThumbUrler) WithThumbUrl(thumbUrl string) {
	tuwh.ThumbUrl = thumbUrl
}

func (tuwh *ThumbHeighter) WithThumbHeight(height int64) {
	tuwh.ThumbHeight = height
}

func (tuwh *ThumbWidther) WithThumbWidth(width int64) {
	tuwh.ThumbWidth = width
}

type InputMessageContenter struct {
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (imcr *InputMessageContenter) WithInputMessageContent(imc InputMessageContent) {
	imcr.InputMessageContent = imc
}

func NewInlineQueryResultArticle(ID string, inputMessageContent InputMessageContent) *InlineQueryResultArticle {
	return &InlineQueryResultArticle{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultArticleType,
			ID:   ID,
		},
		InputMessageContenter: InputMessageContenter{
			InputMessageContent: inputMessageContent,
		},
	}
}

type InlineQueryResultArticle struct {
	BaseInlineQueryResult
	Title       string `json:"title"`
	Url         string `json:"url,omitempty"`
	HideUrl     bool   `json:"hide_url,omitempty"`
	Description string `json:"description,omitempty"`

	ThumbUrler
	ThumbWidther
	ThumbHeighter

	InputMessageContenter
	myTypes.ReplyMarkuper
}

func (iqra *InlineQueryResultArticle) WithUrl(url string) {
	iqra.Url = url
}

func (iqra *InlineQueryResultArticle) WithHideUrl() {
	iqra.HideUrl = true
}

func (iqra *InlineQueryResultArticle) WithDescription(description string) {
	iqra.Description = description
}

type InlineQueryResultPhotoBase struct {
	BaseInlineQueryResult
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	InputMessageContenter
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
}

func (iqrpb *InlineQueryResultPhotoBase) WithTitle(title string) {
	iqrpb.Title = title
}

func (iqrpb *InlineQueryResultPhotoBase) WithDescription(description string) {
	iqrpb.Description = description
}

func NewInlineQueryResultPhoto(ID, photoUrl, thumbUrl string) *InlineQueryResultPhoto {
	return &InlineQueryResultPhoto{
		InlineQueryResultPhotoBase: InlineQueryResultPhotoBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultPhotoType,
				ID:   ID,
			},
		},
		PhotoUrl: photoUrl,
		ThumbUrl: thumbUrl,
	}
}

type InlineQueryResultPhoto struct {
	InlineQueryResultPhotoBase

	PhotoUrl string `json:"photo_url"`
	ThumbUrl string `json:"thumb_url"`

	PhotoWidth  int64 `json:"photo_width,omitempty"`
	PhotoHeight int64 `json:"photo_height,omitempty"`
}

func (iqrp *InlineQueryResultPhoto) WithPhotoWidth(width int64) {
	iqrp.PhotoWidth = width
}

func (iqrp *InlineQueryResultPhoto) WithPhotoHeight(height int64) {
	iqrp.PhotoHeight = height
}

type InlineQueryResultGifBase struct {
	BaseInlineQueryResult
	Title string `json:"title,omitempty"`

	InputMessageContenter
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
}

func (iqrg *InlineQueryResultGifBase) WithTitle(title string) {
	iqrg.Title = title
}

func NewInlineQueryResultGif(ID, gifUrl, thumbUrl string) *InlineQueryResultGif {
	return &InlineQueryResultGif{
		InlineQueryResultGifBase: InlineQueryResultGifBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultGifType,
				ID:   ID,
			},
		},
		GifUrl:   gifUrl,
		ThumbUrl: thumbUrl,
	}
}

type InlineQueryResultGif struct {
	InlineQueryResultGifBase
	GifUrl        string `json:"gif_url"`
	ThumbUrl      string `json:"thumb_url"`
	GifWidth      int64  `json:"gif_width,omitempty"`
	GifHeight     int64  `json:"gif_height,omitempty"`
	GifDuration   int64  `json:"gif_duration,omitempty"`
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`
}

func (iqrg *InlineQueryResultGif) WithGifWidth(width int64) {
	iqrg.GifWidth = width
}

func (iqrg *InlineQueryResultGif) WithGifHeight(height int64) {
	iqrg.GifHeight = height
}

func (iqrg *InlineQueryResultGif) WithGifDuration(duration int64) {
	iqrg.GifDuration = duration
}

func (iqrg *InlineQueryResultGif) WithThumbMimeType(mimeType string) {
	iqrg.ThumbMimeType = mimeType
}

type InlineQueryResultMpeg4GifBase struct {
	BaseInlineQueryResult
	Title string `json:"title,omitempty"`

	InputMessageContenter
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
}

func (iqrmg *InlineQueryResultMpeg4GifBase) WithTitle(title string) {
	iqrmg.Title = title
}

func NewInlineQueryResultMpeg4Gif(ID, mpeg4Url, thumbUrl string) *InlineQueryResultMpeg4Gif {
	return &InlineQueryResultMpeg4Gif{
		InlineQueryResultMpeg4GifBase: InlineQueryResultMpeg4GifBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultGifType,
				ID:   ID,
			},
		},
		Mpeg4Url: mpeg4Url,
		ThumbUrl: thumbUrl,
	}
}

type InlineQueryResultMpeg4Gif struct {
	InlineQueryResultMpeg4GifBase
	Mpeg4Url      string `json:"mpeg4_url"`
	ThumbUrl      string `json:"thumb_url"`
	Mpeg4Width    int64  `json:"mpeg4_width,omitempty"`
	Mpeg4Height   int64  `json:"mpeg4_height,omitempty"`
	Mpeg4Duration int64  `json:"mpeg4_duration,omitempty"`
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`
}

func (iqrmg *InlineQueryResultMpeg4Gif) WithMpeg4Width(width int64) {
	iqrmg.Mpeg4Width = width
}

func (iqrmg *InlineQueryResultMpeg4Gif) WithMpeg4Height(height int64) {
	iqrmg.Mpeg4Height = height
}

func (iqrmg *InlineQueryResultMpeg4Gif) WithMpeg4Duration(duration int64) {
	iqrmg.Mpeg4Duration = duration
}

func (iqrmg *InlineQueryResultMpeg4Gif) WithThumbMimeType(mimeType string) {
	iqrmg.ThumbMimeType = mimeType
}

type InlineQueryResultVideoBase struct {
	BaseInlineQueryResult
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	InputMessageContenter
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
}

func (iqrv *InlineQueryResultVideoBase) WithDescription(description string) {
	iqrv.Description = description
}

func NewInlineQueryResultVideo(ID, videoUrl, mimeType, thumbUrl, title string) *InlineQueryResultVideo {
	return &InlineQueryResultVideo{
		InlineQueryResultVideoBase: InlineQueryResultVideoBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultVideoType,
				ID:   ID,
			},
			Title: title,
		},
		VideoUrl: videoUrl,
		MimeType: mimeType,
		ThumbUrl: thumbUrl,
	}
}

type InlineQueryResultVideo struct {
	InlineQueryResultVideoBase
	VideoUrl string `json:"video_url"`
	MimeType string `json:"mime_type"`
	ThumbUrl string `json:"thumb_url"`

	VideoWidth    int64 `json:"video_width,omitempty"`
	VideoHeight   int64 `json:"video_height,omitempty"`
	VideoDuration int64 `json:"video_duration,omitempty"`
}

func (iqrv *InlineQueryResultVideo) WithVideoWidth(width int64) {
	iqrv.VideoWidth = width
}

func (iqrv *InlineQueryResultVideo) WithVideoHeight(height int64) {
	iqrv.VideoHeight = height
}

func (iqrv *InlineQueryResultVideo) WithVideoDuration(duration int64) {
	iqrv.VideoDuration = duration
}
func NewInlineQueryResultAudio(ID, audioUrl, title string) *InlineQueryResultAudio {
	return &InlineQueryResultAudio{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultAudioType,
			ID:   ID,
		},
		AudioUrl: audioUrl,
		Title:    title,
	}
}

type InlineQueryResultAudio struct {
	BaseInlineQueryResult
	AudioUrl string `json:"audio_url"`
	Title    string `json:"title"`

	Performer     string `json:"performer,omitempty"`
	AudioDuration int64  `json:"audio_duration,omitempty"`

	myTypes.Captioner
	myTypes.ParseModer
}

func (iqra *InlineQueryResultAudio) WithPerformer(performer string) {
	iqra.Performer = performer
}

func (iqra *InlineQueryResultAudio) WithAudioDuration(duration int64) {
	iqra.AudioDuration = duration
}

type InlineQueryResultVoiceBase struct {
	BaseInlineQueryResult

	Title string `json:"title"`
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
	InputMessageContenter
}

func NewInlineQueryResultVoice(ID, voiceUrl, title string) *InlineQueryResultVoice {
	return &InlineQueryResultVoice{
		InlineQueryResultVoiceBase: InlineQueryResultVoiceBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultVoiceType,
				ID:   ID,
			},
			Title: title,
		},
		VoiceUrl: voiceUrl,
	}
}

type InlineQueryResultVoice struct {
	InlineQueryResultVoiceBase
	VoiceUrl      string `json:"voice_url"`
	VoiceDuration int64  `json:"voice_duration,omitempty"`
}

func (iqrv *InlineQueryResultVoice) WithVoiceDuration(duration int64) {
	iqrv.VoiceDuration = duration
}

type InlineQueryResultDocumentBase struct {
	BaseInlineQueryResult
	Title string `json:"title"`

	Description string `json:"description,omitempty"`

	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
	InputMessageContenter
}

func (iqrd *InlineQueryResultDocumentBase) WithDescription(description string) {
	iqrd.Description = description
}

func NewInlineQueryResultDocument(ID, title, docUrl, mimeType string) *InlineQueryResultDocument {
	return &InlineQueryResultDocument{
		InlineQueryResultDocumentBase: InlineQueryResultDocumentBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultDocumentType,
				ID:   ID,
			},
			Title: title,
		},
		DocumentUrl: docUrl,
		MimeType:    mimeType,
	}
}

type InlineQueryResultDocument struct {
	InlineQueryResultDocumentBase
	DocumentUrl string `json:"document_url"`
	MimeType    string `json:"mime_type"`
	ThumbUrler
	ThumbHeighter
	ThumbWidther
}

type InlineQueryResultLocationBase struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64   `json:"live_period,omitempty"`
	Heading              int64   `json:"heading,omitempty"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`
}

func (iqrl *InlineQueryResultLocationBase) WithHorizontalAccuracy(ha float64) {
	iqrl.HorizontalAccuracy = ha
}

func (iqrl *InlineQueryResultLocationBase) WithLivePeriod(lp int64) {
	iqrl.LivePeriod = lp
}

func (iqrl *InlineQueryResultLocationBase) WithHeading(h int64) {
	iqrl.Heading = h
}

func (iqrl *InlineQueryResultLocationBase) WithProximityAlertRadius(r int64) {
	iqrl.ProximityAlertRadius = r
}

func NewInlineQueryResultLocation(ID, title string, latitude, longitude float64) *InlineQueryResultLocation {
	return &InlineQueryResultLocation{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultLocationType,
			ID:   ID,
		},
		InlineQueryResultLocationBase: InlineQueryResultLocationBase{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Title: title,
	}
}

type InlineQueryResultLocation struct {
	BaseInlineQueryResult
	InlineQueryResultLocationBase
	Title string `json:"title"`

	myTypes.ReplyMarkuper
	InputMessageContenter

	ThumbUrler
	ThumbHeighter
	ThumbWidther
}

type InlineQueryResultVenueBase struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Address   string  `json:"address"`

	FoursquareID    string `json:"foursquare_id,omitempty"`
	FoursquareType  string `json:"foursquare_type,omitempty"`
	GooglePlaceID   string `json:"google_place_id,omitempty"`
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

func (v *InlineQueryResultVenueBase) WithFoursquareID(foursquareID string) {
	v.FoursquareID = foursquareID
}

func (v *InlineQueryResultVenueBase) WithFoursquareType(foursquareType string) {
	v.FoursquareType = foursquareType
}

func (v *InlineQueryResultVenueBase) WithGooglePlaceID(googlePlaceID string) {
	v.GooglePlaceID = googlePlaceID
}

func (v *InlineQueryResultVenueBase) WithGooglePlaceType(googlePlaceType string) {
	v.GooglePlaceType = googlePlaceType
}

func NewInlineQueryResultVenue(ID, title, address string, latitude, longitude float64) *InlineQueryResultVenue {
	return &InlineQueryResultVenue{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultVenueType,
			ID:   ID,
		},
		InlineQueryResultVenueBase: InlineQueryResultVenueBase{
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title,
			Address:   address,
		},
	}
}

type InlineQueryResultVenue struct {
	BaseInlineQueryResult
	InlineQueryResultVenueBase

	myTypes.ReplyMarkuper
	InputMessageContenter
	ThumbUrler
	ThumbHeighter
	ThumbWidther
}

type InlineQueryResultContactBase struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`

	LastName string `json:"last_name,omitempty"`
	Vcard    string `json:"vcard,omitempty"`
}

func (iqrc *InlineQueryResultContactBase) WithLastName(ln string) {
	iqrc.LastName = ln
}

func (iqrc *InlineQueryResultContactBase) WithVcard(vcard string) {
	iqrc.Vcard = vcard
}

func NewInlineQueryResultContact(ID, phoneNumber, firstName string) *InlineQueryResultContact {
	return &InlineQueryResultContact{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultContactType,
			ID:   ID,
		},
		InlineQueryResultContactBase: InlineQueryResultContactBase{
			PhoneNumber: phoneNumber,
			FirstName:   firstName,
		},
	}
}

type InlineQueryResultContact struct {
	BaseInlineQueryResult
	InlineQueryResultContactBase
	myTypes.ReplyMarkuper
	InputMessageContenter
	ThumbUrler
	ThumbHeighter
	ThumbWidther
}

func NewInlineQueryResultGame(ID, gameShortName string) *InlineQueryResultGame {
	return &InlineQueryResultGame{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultGameType,
			ID:   ID,
		},
		GameShortName: gameShortName,
	}
}

type InlineQueryResultGame struct {
	BaseInlineQueryResult
	GameShortName string `json:"game_short_name"`
	myTypes.ReplyMarkuper
}

func NewInlineQueryResultCachedPhoto(ID, photoFileID string) *InlineQueryResultCachedPhoto {
	return &InlineQueryResultCachedPhoto{
		InlineQueryResultPhotoBase: InlineQueryResultPhotoBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultCachedPhotoType,
				ID:   ID,
			},
		},
		PhotoFileID: photoFileID,
	}
}

type InlineQueryResultCachedPhoto struct {
	InlineQueryResultPhotoBase
	PhotoFileID string `json:"photo_file_id"`
}

func NewInlineQueryResultCachedGif(ID, gifFileID string) *InlineQueryResultCachedGif {
	return &InlineQueryResultCachedGif{
		InlineQueryResultGifBase: InlineQueryResultGifBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultCachedGifType,
				ID:   ID,
			},
		},
		GifFileID: gifFileID,
	}
}

type InlineQueryResultCachedGif struct {
	InlineQueryResultGifBase
	GifFileID string `json:"gif_file_id"`
}

func NewInlineQueryResultCachedMpeg4Gif(ID, mpeg4FileID string) *InlineQueryResultCachedMpeg4Gif {
	return &InlineQueryResultCachedMpeg4Gif{
		InlineQueryResultMpeg4GifBase: InlineQueryResultMpeg4GifBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultCachedMpeg4GifType,
				ID:   ID,
			},
		},
		Mpeg4FileID: mpeg4FileID,
	}
}

type InlineQueryResultCachedMpeg4Gif struct {
	InlineQueryResultMpeg4GifBase
	Mpeg4FileID string `json:"mpeg4_file_id"`
}

func NewInlineQueryResultCachedSticker(ID, stickerFileID string) *InlineQueryResultCachedSticker {
	return &InlineQueryResultCachedSticker{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultCachedStickerType,
			ID:   ID,
		},
		StickerFileID: stickerFileID,
	}
}

type InlineQueryResultCachedSticker struct {
	BaseInlineQueryResult
	StickerFileID string `json:"sticker_file_id"`

	myTypes.ReplyMarkuper
	InputMessageContenter
}

func NewInlineQueryResultCachedDocument(ID, title, docFileID string) *InlineQueryResultCachedDocument {
	return &InlineQueryResultCachedDocument{
		InlineQueryResultDocumentBase: InlineQueryResultDocumentBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultCachedDocumentType,
				ID:   ID,
			},
			Title: title,
		},
	}
}

type InlineQueryResultCachedDocument struct {
	InlineQueryResultDocumentBase
	DocumentFileID string `json:"document_file_id"`
}

func NewInlineQueryResultCachedVideo(ID, videoFileID, title string) *InlineQueryResultCachedVideo {
	return &InlineQueryResultCachedVideo{
		InlineQueryResultVideoBase: InlineQueryResultVideoBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultCachedVideoType,
				ID:   ID,
			},
			Title: title,
		},
		VideoFileID: videoFileID,
	}
}

type InlineQueryResultCachedVideo struct {
	InlineQueryResultVideoBase
	VideoFileID string `json:"video_file_id"`
}

func NewInlineQueryResultCachedVoice(ID, voiceFileID, title string) *InlineQueryResultCachedVoice {
	return &InlineQueryResultCachedVoice{
		InlineQueryResultVoiceBase: InlineQueryResultVoiceBase{
			BaseInlineQueryResult: BaseInlineQueryResult{
				Type: InlineQueryResultCachedVoiceType,
			},
			Title: title,
		},
	}
}

type InlineQueryResultCachedVoice struct {
	InlineQueryResultVoiceBase
	VoiceFileID string `json:"voice_file_id"`
}

func NewInlineQueryResultCachedAudio(ID, audiFileID string) *InlineQueryResultCachedAudio {
	return &InlineQueryResultCachedAudio{
		BaseInlineQueryResult: BaseInlineQueryResult{
			Type: InlineQueryResultCachedAudioType,
			ID:   ID,
		},
		AudioFileID: audiFileID,
	}
}

type InlineQueryResultCachedAudio struct {
	BaseInlineQueryResult
	AudioFileID string `json:"audio_file_id"`
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
	InputMessageContenter
}
