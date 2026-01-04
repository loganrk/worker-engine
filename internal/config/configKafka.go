package config

func (k kafka) GetBrokers() []string {
	return k.Brokers
}

func (k kafka) GetActivationTopic() string {
	return k.Topics.Activation
}

func (k kafka) GetPasswordResetTopic() string {
	return k.Topics.PasswordReset
}

func (k kafka) GetConsumerGroupName() string {
	return k.ConsumerGroupName
}
