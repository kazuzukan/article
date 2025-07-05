package Repository

import (
	"articles-consumer/Config"
	"articles-consumer/Repository/Articles"
)

type Repository struct {
	Articles Articles.Repository
}

func InitRepository(dbCon Config.ConnectionInterface) Repository {
	return Repository{
		Articles: Articles.NewRepository(dbCon),
	}
}
