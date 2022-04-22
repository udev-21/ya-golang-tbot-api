package universal

import "github.com/udev21/golang-tbot-api/types"

type InlineQueryResult struct {
	Type                string                      `json:"type"`
	ID                  string                      `json:"id"`
	Title               *string                     `json:"title,omitempty"`
	InputMessageContent *types.InputMessageContent  `json:"input_message_content,omitempty"`
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
