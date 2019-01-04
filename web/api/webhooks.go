package api

import (
	"net/http"

	"github.com/LizardsTown/opennode"
	"github.com/go-chi/chi"
	"github.com/mholt/binding"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/payloads"
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
		charge := opennode.Charge{ID: payload.ID}
		err := opennode.NewClient(&a.Config.PaymentConfig).UpdateCharge(&charge)
		if err != nil {
			return err
		}

		return render(w, r, nil, http.StatusOK)
	}
}
