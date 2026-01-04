package port

import "context"

type SvrList struct {
	User UserSvr
}

type UserSvr interface {
	ActivationEmail(ctx context.Context, to, subject string, macros map[string]string) error
	ActivationSms(ctx context.Context, contryCode, to string, macros map[string]string) error

	PasswordResetEmail(ctx context.Context, to, subject string, macros map[string]string) error
	PasswordResetSms(ctx context.Context, contryCode, to string, macros map[string]string) error
}
