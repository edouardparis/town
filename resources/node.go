package resources

import (
	"time"

	"git.iiens.net/edouardparis/town/models"
)

type Node struct {
	PubKey          string     `json:"pub_key"`
	AmountCollected int64      `json:"amount_collected"`
	AmountReceived  int64      `json:"amount_received"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

func NewNode(node *models.Node) *Node {
	if node == nil {
		return nil
	}

	resource := &Node{
		PubKey:          node.PubKey,
		AmountCollected: node.AmountCollected,
		AmountReceived:  node.AmountReceived,
		CreatedAt:       node.CreatedAt,
	}

	if node.UpdatedAt.Valid {
		resource.UpdatedAt = &node.UpdatedAt.Time
	}

	return resource
}
