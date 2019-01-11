package resources

import (
	"github.com/EdouardParis/town/constants"
	"github.com/EdouardParis/town/models"
)

type Header struct {
	URLs    URLs
	Pricing Pricing
}

type URLs struct {
	Website string
}

type Pricing struct {
	ArticlePublicationPrice int64
}

func NewHeader(i *models.Info) *Header {
	return &Header{
		URLs: URLs{
			Website: i.URLs.Website,
		},
		Pricing: Pricing{
			ArticlePublicationPrice: constants.ArticlePublicationPrice,
		},
	}
}
