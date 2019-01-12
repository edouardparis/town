package resources

import (
	"github.com/EdouardParis/town/models"
	"github.com/hackebrot/turtle"
)

type Reactions map[string]int

func NewReactions(reactions []models.ArticleReaction) Reactions {
	resource := make(Reactions)

	for i := range reactions {
		emoji := turtle.Emojis[reactions[i].Emoji]
		count := resource[emoji.Char]
		count++
		resource[emoji.Char] = count
	}

	return resource
}
