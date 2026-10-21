package bot

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/bwmarrin/discordgo"
)

const morgulBladeCmd = "morgul-blade"

//go:embed assets/morgul-blade.png
var morgulBladePNG []byte

func morgulBladeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        morgulBladeCmd,
		Description: "Stabs Frodo with the Morgul Blade to silence him indefinitely",
		Contexts:    &[]discordgo.InteractionContextType{discordgo.InteractionContextGuild},
	}
}

func (b *Bot) onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if i.ApplicationCommandData().Name != morgulBladeCmd {
		return
	}

	if !b.wounds.Stab(i.GuildID) {
		respond(s, i, &discordgo.InteractionResponseData{
			Content: "You awful person, Frodo has already been stabbed.",
			Flags:   discordgo.MessageFlagsEphemeral,
		})
		return
	}

	respond(s, i, &discordgo.InteractionResponseData{
		Content: "DEUUEAUGH!!",
		Files: []*discordgo.File{
			{
				Name:        "morgul-blade.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(morgulBladePNG),
			},
		},
	})
}

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, data *discordgo.InteractionResponseData) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
	if err != nil {
		log.Printf("respond to /%s: %v", i.ApplicationCommandData().Name, err)
	}
}
