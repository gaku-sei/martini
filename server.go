package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func Server(db *DB) *martini.ClassicMartini {
	m := martini.Classic()

	m.Map(db)

	m.Use(func(res http.ResponseWriter) {
		res.Header().Set("ContentType", "application/json; charset=utf-8")
	})

	m.Group("/games", func(r martini.Router) {
		r.Get("/search/:name", func(params martini.Params, res http.ResponseWriter) {
			if err := PrepareXML(GetGamesList, Query(params)).Stream(new(Games), res); err != nil {
				panic(err)
			}
		})

		r.Get("/:id", func(params martini.Params, res http.ResponseWriter) {
			if err := PrepareXML(GetGame, Query(params)).Stream(new(Game), res); err != nil {
				panic(err)
			}
		})
	})

	m.Group("/platforms", func(r martini.Router) {
		r.Get("", func(res http.ResponseWriter) {
			if err := PrepareXML(GetPlatformsList, Query{}).Stream(new(Platforms), res); err != nil {
				panic(err)
			}
		})

		r.Get("/:id", func(params martini.Params, res http.ResponseWriter) {
			if err := PrepareXML(GetPlatform, Query(params)).Stream(new(Platform), res); err != nil {
				panic(err)
			}
		})

		r.Get("/:platform/games", func(params martini.Params, res http.ResponseWriter) {
			if err := PrepareXML(GetPlatformGames, Query(params)).Stream(new(Games), res); err != nil {
				panic(err)
			}
		})
	})

	m.Group("/personal", func(r martini.Router) {
		r.Get("/favorites", func(db *DB, res http.ResponseWriter) {
			var favorites []Favorite
			if err := db.Fetch(&favorites).All(res); err != nil {
				panic(err)
			}
		})

		r.Get("/owned-games", func(db *DB, res http.ResponseWriter) {
			var ownedGames []OwnedGame
			if err := db.Fetch(&ownedGames).All(res); err != nil {
				panic(err)
			}
		})
	})

	m.Get("/", func(res http.ResponseWriter) {
		res.Write([]byte(fmt.Sprintf(`{"version": "%s"}`, version)))
	})

	return m
}
