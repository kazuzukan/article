package Routes

import (
	"articles-publisher/Controller"
	"github.com/go-chi/chi"
)

type Routes struct {
	Controller Controller.Controller
	Chi        *chi.Mux
}

func (r *Routes) GetRoutes() *chi.Mux {
	router := r.Chi
	router.Group(func(router chi.Router) {
		router.Route("/article", func(router chi.Router) {
			router.Post("/create", r.Controller.Articles.CreateArticles)
			router.Get("/list", r.Controller.Articles.GetArticlesList)
		})
	})

	return router
}
