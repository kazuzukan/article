package ConfigSubscriber

type articlesConfig interface {
	CreateArticlesConfig() ConfigSubs
}

func (s ConfigSubs) CreateArticlesConfig() ConfigSubs {
	s.ExchangeName = "articles"
	s.QueueName = "articles.queue"
	s.BindingKey = "articles.create.key"
	s.ExchangeType = "topic"
	return s
}
