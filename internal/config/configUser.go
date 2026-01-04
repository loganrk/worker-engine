package config

func (u user) GetActivationEmailTemplatePath() string {

	return u.Activation.EmailTemplatePath
}

func (u user) GetPasswordResetEmailTemplatePath() string {

	return u.PasswordReset.EmailTemplatePath
}

func (u user) GetActivationSmsTemplatePath() string {

	return u.Activation.EmailTemplatePath
}

func (u user) GetPasswordResetSmsTemplatePath() string {

	return u.PasswordReset.EmailTemplatePath
}
