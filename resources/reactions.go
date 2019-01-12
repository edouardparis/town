package resources

import (
	"github.com/EdouardParis/town/models"
	"github.com/hackebrot/turtle"
)

type Reactions map[string]int

func NewReactions(reactions []models.Reaction) Reactions {
	resource := make(Reactions)

	for i := range reactions {
		emoji := turtle.Emojis[reactions[i].Emoji]
		count := resource[emoji.Char]
		count++
		resource[emoji.Char] = count
	}

	return resource
}

type Reaction struct {
	Emoji string `json:"emoji"`
}

func NewReaction(reaction *models.Reaction) Reaction {
	emoji := turtle.Emojis[reaction.Emoji]
	return Reaction{Emoji: emoji.Char}
}
