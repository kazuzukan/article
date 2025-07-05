package Service

import (
	"articles-consumer/Config"
	"articles-consumer/Controller"
	"articles-consumer/Controller/Dto"
	"articles-consumer/Modules"
	"articles-consumer/Repository"
	"articles-consumer/Routes"
	"articles-consumer/Service/RabbitMQ"
	"github.com/go-chi/chi"
	"net/http"
)

func InitApplication() {
	connection := Config.BuildConnection()
	rabbitConfigPublish := RabbitMQ.GetRabbitMQService(connection.RabbitMQConnection())
	repoInit := Repository.InitRepository(connection)
	utilities := Dto.Utilities{
		RabbitMQ: rabbitConfigPublish,
		Modules:  Modules.InitModule(repoInit),
	}

	router := Routes.Routes{
		Chi:        chi.NewRouter(),
		Controller: Controller.InitControllerAPI(utilities),
	}

	subscriber := Routes.NewRabbitMQSubscriber(connection.RabbitMQConnection(), Controller.InitController(utilities))
	subscriber.StartSubscriber()
	err := http.ListenAndServe(":9000", router.GetRoutes())
	if err != nil {
		panic(err)
	}
}
