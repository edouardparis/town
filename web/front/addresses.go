package front

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/failures"
	"github.com/EdouardParis/town/resources"
	"github.com/EdouardParis/town/store"
	"github.com/EdouardParis/town/web/middlewares"
)

func addressesRoutes(a *app.App) func(r chi.Router) {
	handle := newHandle(a)
	addressCtx := middlewares.AddressCtx(a, handleError(a.Logger))
	return func(r chi.Router) {
		r.Get("/", handle(AddressList))
		r.Route("/{value}", func(r chi.Router) {
			r.With(addressCtx).Get("/", handle(AddressDetail))
		})
	}
}

func AddressDetail(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	data := struct {
		Header    *resources.Header
		Addresses []resources.Address
	}{}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		address, ok := middlewares.AddressFromCtx(ctx)
		if !ok || address == nil {
			handle(w, r, failures.ErrNotFound)
			return
		}
		data.Addresses = []resources.Address{*resources.NewAddress(address)}
		data.Header = resources.NewHeader(a.Info)
		err := render(w, r, "addresses.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}

func AddressList(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	data := struct {
		Addresses []resources.Address
		Header    *resources.Header
	}{Header: resources.NewHeader(a.Info)}
	s := store.NewAddresses(a.Store)
	return func(w http.ResponseWriter, r *http.Request) {
		addresses, err := s.List(r.Context())
		if err != nil {
			handle(w, r, err)
			return
		}

		data.Addresses = resources.NewAddressList(addresses)

		err = render(w, r, "addresses.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}
