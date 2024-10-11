package testMiddleware

import (
	"net/http"

	"github.com/gorilla/mux"

	"awesome/awesomeProject/testMiddleware/handler"
	"awesome/awesomeProject/testMiddleware/middleware"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()

	uh := handler.NewUploadHandler()

	publicRoutes := router.PathPrefix("").Subrouter()
	publicRoutes.Handle("/upload", uh).Methods(http.MethodGet)

	router = router.SkipClean(true)

	publicRoutes.Use(
		middleware.NewMetaLogger("log"),
		middleware.NewObserver("metrics"),
	)

	return router
}
