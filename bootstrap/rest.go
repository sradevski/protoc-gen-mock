package bootstrap

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sradevski/protoc-gen-mock/grpchandler"
	"github.com/sradevski/protoc-gen-mock/restcontrollers"
	"github.com/sradevski/protoc-gen-mock/stub"
	"net/http"
)

func StartRESTServer(port uint, controllers []restcontrollers.RESTController) {
	log.Infof("REST Server listening on port: %d", port)

	r := mux.NewRouter()
	for _, controller := range controllers {
		api := r.PathPrefix(controller.GetPath()).Subrouter()
		for _, handler := range controller.GetHandlers() {
			api.HandleFunc(handler.Path, handler.Handler).Methods(handler.Methods...)
		}
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func CreateRESTControllers(
	stubExamples []stub.Stub,
	stubsStore stub.StubsStore,
	recordingsStore stub.RecordingsStore,
	service grpchandler.MockService) []restcontrollers.RESTController {
	return []restcontrollers.RESTController{
		restcontrollers.ExamplesController{StubExamples: stubExamples},
		restcontrollers.StubsController{
			StubsStore:   stubsStore,
			StubExamples: stubExamples,
			Service:      service,
		},
		restcontrollers.RecordingsController{
			RecordingsStore: recordingsStore,
		},
	}
}
