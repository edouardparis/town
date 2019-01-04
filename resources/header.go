package resources

import (
	"github.com/EdouardParis/town/models"
)

type Header struct {
	URLs URLs
}

type URLs struct {
	Website string
}

func NewHeader(i *models.Info) *Header {
	return &Header{URLs: URLs{
		Website: i.URLs.Website,
	}}
}
