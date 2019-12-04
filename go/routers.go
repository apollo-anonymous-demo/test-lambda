/*
 * PetStore
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2019-12-04T02:15:00Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

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
	fmt.Fprintf(w, "Welcome!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/Alpha/",
		Index,
	},

	Route{
		"CreatePet",
		strings.ToUpper("Post"),
		"/Alpha/pets",
		CreatePet,
	},

	Route{
		"GetPet",
		strings.ToUpper("Get"),
		"/Alpha/pets/{petId}",
		GetPet,
	},

	Route{
		"PetsGet",
		strings.ToUpper("Get"),
		"/Alpha/pets",
		PetsGet,
	},

	Route{
		"PetsOptions",
		strings.ToUpper("Options"),
		"/Alpha/pets",
		PetsOptions,
	},

	Route{
		"PetsPetIdOptions",
		strings.ToUpper("Options"),
		"/Alpha/pets/{petId}",
		PetsPetIdOptions,
	},

	Route{
		"RootGet",
		strings.ToUpper("Get"),
		"/Alpha/",
		RootGet,
	},
}
