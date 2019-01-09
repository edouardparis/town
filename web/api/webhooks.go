package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/LizardsTown/opennode"
	"github.com/go-chi/chi"
	"github.com/mholt/binding"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/constants"
	"github.com/EdouardParis/town/logging"
	"github.com/EdouardParis/town/models"
	"github.com/EdouardParis/town/payloads"
	"github.com/EdouardParis/town/resources"
	"github.com/EdouardParis/town/store"
	"github.com/EdouardParis/town/web/websockets"
)

func webhooksRoutes(a *app.App) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/checkout", handle(a, CheckoutWebhook))
	}
}

func CheckoutWebhook(a *app.App) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		payload := payloads.Charge{}
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			return err
		}
		a.Logger.Info("new webhook", logging.String("request", fmt.Sprintf("%q", dump)))

		errs := binding.Form(r, &payload)
		if errs != nil {
			return errs
		}
		a.Logger.Info("binding success", logging.String("payload", fmt.Sprintf("%d", payload.ID)))
		charge := opennode.Charge{ID: payload.ID}
		err = opennode.NewClient(&a.Config.PaymentConfig).UpdateCharge(&charge)
		if err != nil {
			return err
		}

		if charge.Status != constants.OrderStatusPaidStr {
			return render(w, r, nil, http.StatusOK)
		}

		order := models.NewOrder(&charge)
		err = store.NewOrders(a.Store).Create(r.Context(), order)
		if err != nil {
			return err
		}

		a.Logger.Info("new order paid",
			logging.Int64("amount", order.Amount),
			logging.String("notes", order.Notes),
		)

		err = websockets.SendChargeAndCloseSession(resources.NewCharge(&charge))
		if err != nil {
			return err
		}

		return render(w, r, nil, http.StatusOK)
	}
}
