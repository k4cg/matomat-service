package api

import (
	"net/http"

	"github.com/k4cg/matomat-service/maas-server/auth"
	"github.com/k4cg/matomat-service/maas-server/matomat"
)

type ServiceApiHandler struct {
	auth    auth.AuthInterface
	matomat *matomat.Matomat
}

func NewServiceApiHandler(auth auth.AuthInterface, matomat *matomat.Matomat) *ServiceApiHandler {
	return &ServiceApiHandler{auth: auth, matomat: matomat}
}

func (sah *ServiceApiHandler) ServiceStatsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	w.Header().Set("Cache-Control", "max-age=30") //allow caching by the client for 30 seconds
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil && sah.matomat.IsAllowed(loginUserID, matomat.ACTION_SERVICE_STATS) {
		stats, err := sah.matomat.ServiceStatsGet()
		if err == nil {
			status, response = getResponse(http.StatusOK, newServiceStats(stats))
		} else {
			status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}
