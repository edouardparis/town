package websockets

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	chirender "github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"gopkg.in/olahol/melody.v1"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/logging"
	"git.iiens.net/edouardparis/town/opennode"
)

var sessions = struct {
	counter int
	objects map[string]*melody.Session
	sync.RWMutex
}{
	objects: make(map[string]*melody.Session),
}

func NewRouter(a *app.App) http.Handler {
	mrouter := melody.New()
	r := chi.NewRouter()
	r.Get("/checkout", handleError(a.Logger, mrouter.HandleRequest))

	mrouter.HandleConnect(func(s *melody.Session) {
		sessions.Lock()
		defer sessions.Unlock()

		sessions.objects[funk.RandomString(10)] = s
		sessions.counter += 1
		a.Logger.Info("New websocket connection", logging.Int("total_connected", sessions.counter))

		charge, err := opennode.NewClient(a.PaymentConfig).CreateCharge(&opennode.ChargePayload{
			Amount: int64(1000),
		})
		if err != nil {
			s.Close()
			a.Logger.Error("Error during charge creation", logging.Error(err))
			return
		}
		resource := struct {
			PayReq string `json:"payreq"`
			Amount int64  `json:"amount"`
		}{
			PayReq: charge.LightningInvoice.PayReq,
			Amount: charge.Amount,
		}

		rsc, err := json.Marshal(resource)
		if err != nil {
			s.Close()
			a.Logger.Error("Error during charge creation", logging.Error(err))
			return
		}

		s.Write(rsc)
	})
	return r
}

func handleError(logger logging.Logger, fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err == nil {
			return
		}

		var status int
		switch cerr := errors.Cause(err).(type) {
		case failures.Error:
			status = cerr.Code
			err = cerr
		default:
			logger.Error(cerr.Error())
			status = http.StatusInternalServerError
		}

		chirender.Status(r, status)
		chirender.JSON(w, r, err)
	}
}
