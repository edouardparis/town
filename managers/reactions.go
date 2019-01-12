package managers

import (
	"context"

	"github.com/EdouardParis/town/constants"
	"github.com/EdouardParis/town/failures"
	"github.com/EdouardParis/town/models"
	"github.com/EdouardParis/town/payloads"
	"github.com/EdouardParis/town/store"
)

func ReactionCreate(ctx context.Context, s store.Store, payload *payloads.Reaction, article *models.Article) (*models.Reaction, error) {
	reaction := &models.Reaction{
		Emoji:     payload.Emoji,
		ArticleID: article.ID,
	}

	order, err := store.NewOrders(s).GetByPublicID(ctx, payload.OrderID)
	if err != nil {
		if err == failures.ErrNotFound {
			return nil, failures.ErrBadRequest
		}
		return nil, err
	}

	if order.Status == constants.OrderStatusClaimed {
		return nil, failures.ErrBadRequest
	}

	reaction.OrderID = order.ID

	err = s.Transaction(ctx, func(tx store.Store) error {
		err := store.NewOrders(tx).MarkOrderAs(ctx, order.ID, constants.OrderStatusClaimed)
		if err != nil {
			return err
		}

		return store.NewReactions(tx).Create(ctx, reaction)
	})
	if err != nil {
		return nil, err
	}

	return reaction, nil
}
