package main

import (
	"github.com/gorilla/mux"
	"github.com/johnmcdnl/darts-backend/board"
	"github.com/johnmcdnl/darts-backend/practice/targets"
	"net/http"
)

func main() {
	migrate()
	http.ListenAndServe(":9000", registerRoutes())
}

func registerRoutes() *mux.Router {
	r := mux.NewRouter()
	board.BoardRouter(r)
	targets.TargetRouter(r)
	return r
}

func migrate() {
	targets.Migrate()
}
