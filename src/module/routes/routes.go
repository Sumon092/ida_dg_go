package routes

import (
	"database/sql"
	"fmt"
	"ida_diag/src/module/user"
	"net/http"
)

// Route represents a route definition.
type Route struct {
	Path   string
	Handle http.HandlerFunc
}

// RegisterRoutes registers routes from a slice of Route.
func RegisteredRoutes(db *sql.DB) []Route {
	routes := []Route{
		{Path: "/route2", Handle: func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Route is heated")
		}},
	}
	userRoutesHandler := user.UserRoutes(db)
	routes = append(routes, Route{Path: "/users", Handle: userRoutesHandler.ServeHTTP})

	return routes
}
