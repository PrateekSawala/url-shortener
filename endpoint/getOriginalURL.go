package endpoint

import (
	"net/http"

	"github.com/PrateekSawala/url-shortener/domain"
	"github.com/PrateekSawala/url-shortener/domain/logging"
	"github.com/PrateekSawala/url-shortener/service"

	"github.com/gorilla/mux"
)

func GetOriginalURLHandler(writer http.ResponseWriter, request *http.Request) {
	log := logging.Log("GetOriginalURLHandler")
	log.Tracef("Start")
	defer log.Tracef("End")

	routeParameter, ok := mux.Vars(request)["id"]
	if !ok {
		log.Debugf("id parameter is missing in the request")
		http.Error(writer, domain.EmptyInput, http.StatusBadRequest)
		return
	}

	if !IsUUIDValid(routeParameter) {
		log.Debugf("uuid %s is invalid", routeParameter)
		http.Error(writer, domain.InvalidInput, http.StatusBadRequest)
		return
	}

	getLongURLResp, err := service.GetLongURL(routeParameter)
	if err != nil {
		log.Debugf("service.GetLongURL error: %s", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(writer, request, getLongURLResp.Url, http.StatusSeeOther)
	return
}
