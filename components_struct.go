package discord

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

// TODO: Renames.

// Type = 1.
type ComponentSend struct {
	Components []Comp `json:"components"`
	// Type = 1. Other options not supported.
	//
	// 1: ActionsRowComponent
	// 2: ButtonComponent
	// 3: SelectMenuComponent
	// 4: TextInputComponent
	// 5: UserSelectMenuComponent
	// 6: RoleSelectMenuComponent
	// 7: MentionableSelectMenuComponent
	// 8: ChannelSelectMenuComponent
	Type int `json:"type"`
}
type Comp struct {
	Label string `json:"label"`
	// 1: PrimaryButton is a button with blurple color.
	// 2: SecondaryButton is a button with grey color.
	// 3: SuccessButton is a button with green color.
	// 4: DangerButton is a button with red color.
	// 5: LinkButton is a special type of button which navigates to a URL. Has grey color.
	Style    int  `json:"style"`
	Disabled bool `json:"disabled"`
	// Should not use name and id field in it...
	// Have no idea how to solve.
	Emoji    struct{} `json:"emoji,omitempty"`
	URL      string   `json:"url,omitempty"`
	CustomID string   `json:"custom_id"`
	// 1: TextInputShort
	// 2: TextInputParagraph
	Type int `json:"type"`
}

type MsgComponents struct {
	CompSend ComponentSend
}

func (com MsgComponents) Type() discordgo.ComponentType { return discordgo.ButtonComponent }
func (com MsgComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(com.CompSend)
}
