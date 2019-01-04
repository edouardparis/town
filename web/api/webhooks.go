package api

import (
	"net/http"

	"github.com/LizardsTown/opennode"
	"github.com/go-chi/chi"
	"github.com/mholt/binding"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/payloads"
	"github.com/EdouardParis/town/resources"
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
		errs := binding.Bind(r, &payload)
		if errs != nil {
			return errs
		}
		charge := opennode.Charge{}
		err := opennode.NewClient(&a.Config.PaymentConfig).UpdateCharge(&charge)
		if err != nil {
			return err
		}

		err = websockets.SendChargeAndCloseSession(resources.NewCharge(&charge))
		if err != nil {
			return err
		}

		return render(w, r, nil, http.StatusOK)
	}
}
