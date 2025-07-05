package Routes

import (
	"articles-consumer/Controller"
	"github.com/go-chi/chi"
)

type Routes struct {
	Controller Controller.ControllerAPI
	Chi        *chi.Mux
}

func (r *Routes) GetRoutes() *chi.Mux {
	router := r.Chi
	router.Group(func(router chi.Router) {
		router.Route("/articles", func(router chi.Router) {
			router.Post("/list", r.Controller.Articles.GetArticlesList)
		})
	})

	return router
}
