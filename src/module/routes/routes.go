package routes

import (
	"fmt"
	"net/http"
)

// Route represents a route definition.
type Route struct {
	Path   string
	Handle http.HandlerFunc
}

// RegisterRoutes registers routes from a slice of Route.
func RegisteredRoutes() []Route {
	routes := []Route{
		{Path: "/route1", Handle: func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Woow I am flying with gin server")
		}},
		{Path: "/route2", Handle: func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Route is heated")
		}},
	}

	return routes
}
