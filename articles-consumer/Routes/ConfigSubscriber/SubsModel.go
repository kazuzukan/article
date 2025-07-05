package ConfigSubscriber

type SubsConfig struct {
	Articles articlesConfig
}

// config subs
type ConfigSubs struct {
	ExchangeName string `json:"exchange_name"`
	BindingKey   string `json:"binding_key"`
	QueueName    string `json:"queue_name"`
	ExchangeType string `json:"exchange_type"`
}
