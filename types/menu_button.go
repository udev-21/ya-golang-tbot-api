package types

type MenuButton struct {
	Type   string      `json:"type"`
	Text   *string     `json:"text,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

const (
	MenuButtonCommandsType = "commands"
	MenuButtonWebAppType   = "web_app"
	MenuButtonDefaultType  = "default"
)

var MenuButtonCommands = MenuButton{
	Type: MenuButtonCommandsType,
}
var MenuButtonDefault = MenuButton{
	Type: MenuButtonDefaultType,
}

func NewMenuButtonWebApp(text string, webApp WebAppInfo) MenuButton {
	return MenuButton{
		Type:   MenuButtonWebAppType,
		Text:   &text,
		WebApp: &webApp,
	}
}
