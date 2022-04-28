package types

type IMenuButton interface {
	IsMenuButton()
}

type MenuButton struct {
	BaseMenuButton
	Text   *string     `json:"text,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

const (
	MenuButtonCommandsType = "commands"
	MenuButtonWebAppType   = "web_app"
	MenuButtonDefaultType  = "default"
)

type BaseMenuButton struct {
	Type string `json:"type"`
}

type MenuButtonCommands struct {
	BaseMenuButton
}

func NewMenuButtonCommands() *MenuButtonCommands {
	return &MenuButtonCommands{
		BaseMenuButton: BaseMenuButton{
			Type: MenuButtonCommandsType,
		},
	}
}

type MenuButtonWebApp struct {
	BaseMenuButton
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}

func NewMenuButtonWebApp(text string, webApp WebAppInfo) *MenuButtonWebApp {
	return &MenuButtonWebApp{
		BaseMenuButton: BaseMenuButton{
			Type: MenuButtonWebAppType,
		},
		Text:   text,
		WebApp: webApp,
	}
}

type MenuButtonDefault struct {
	BaseMenuButton
}

func NewMenuButtonDefault() *MenuButtonDefault {
	return &MenuButtonDefault{
		BaseMenuButton: BaseMenuButton{
			Type: BotCommandScopeDefaultType,
		},
	}
}
