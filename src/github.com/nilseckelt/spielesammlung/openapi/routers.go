package openapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Index",
		"GET",
		"/gc/",
		GetAllGames,
	},

	Route{
		"CreateGame",
		"POST",
		"/gc/{gameName}",
		CreateGame,
	},

	Route{
		"DeleteGame",
		"DELETE",
		"/gc/{gameName}",
		DeleteGame,
	},

	Route{
		"EditGame",
		"PUT",
		"/gc/{gameName}",
		EditGame,
	},

	Route{
		"GetGame",
		"GET",
		"/gc/{gameName}",
		GetGame,
	},

	Route{
		"GetHealth",
		"GET",
		"/health",
		GetHealth,
	},
}
