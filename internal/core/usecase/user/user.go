package user

import (
	"context"
	"fmt"
	"os"

	"github.com/loganrk/worker-engine/internal/core/port"
	"github.com/loganrk/worker-engine/internal/utils"
)

// userusecase implements user-related operations such as sending activation and password reset emails.
type userusecase struct {
	logger                port.Logger  // Logger interface for structured logging
	activationEmailTpl    string       // Email template content for activation emails
	passwordResetEmailTpl string       // Email template content for password reset emails
	activationSmsTpl      string       // SMS template content for activation emails
	passwordResetSmsTpl   string       // SMS template content for password reset emails
	emailer               port.Emailer // Interface to send emails
	emailRateLimiter      port.RateLimiter
	smsSender             port.SmsSender
}

// New initializes a new userusecase instance by loading email templates and setting dependencies.
func New(userConf port.ConfUser, loggerIns port.Logger, emailerIns port.Emailer, emailRateLimitIns port.RateLimiter) (*userusecase, error) {
	// Read activation email template from file
	activationEmailTpl, err := os.ReadFile(userConf.GetActivationEmailTemplatePath())
	if err != nil {
		return nil, fmt.Errorf("failed to load activation template: %w", err)
	}

	// Read password reset email template from file
	passwordResetEmailTpl, err := os.ReadFile(userConf.GetPasswordResetEmailTemplatePath())
	if err != nil {
		return nil, fmt.Errorf("failed to load password reset template: %w", err)
	}

	// Read activation sms template from file
	activationSmsTpl, err := os.ReadFile(userConf.GetActivationSmsTemplatePath())
	if err != nil {
		return nil, fmt.Errorf("failed to load activation template: %w", err)
	}

	// Read password reset sms template from file
	passwordResetSmsTpl, err := os.ReadFile(userConf.GetPasswordResetSmsTemplatePath())
	if err != nil {
		return nil, fmt.Errorf("failed to load password reset template: %w", err)
	}

	// Return the fully initialized userusecase
	return &userusecase{
		logger:                loggerIns,
		emailer:               emailerIns,
		emailRateLimiter:      emailRateLimitIns,
		activationEmailTpl:    string(activationEmailTpl),
		passwordResetEmailTpl: string(passwordResetEmailTpl),
		activationSmsTpl:      string(activationSmsTpl),
		passwordResetSmsTpl:   string(passwordResetSmsTpl),
	}, nil
}

func (u *userusecase) ActivationEmail(ctx context.Context, to, subject string, macros map[string]string) error {
	u.logger.Infow(ctx, "Processing Activation Email", "to", to, "subject", subject, "macros", macros)

	if u.emailRateLimiter != nil {
		err := u.emailRateLimiter.WaitUntilAllowed(context.Background())

		if err != nil {
			u.logger.Errorw(ctx, "Failed to send activation email due to rate limit error", "error", err)
			return err
		}
	}

	emailBody := utils.ReplaceMacros(u.activationEmailTpl, macros)
	if err := u.emailer.SendEmail(to, subject, emailBody); err != nil {
		u.logger.Errorw(ctx, "Failed to send activation email", "error", err)
	}
	return nil
}

func (u *userusecase) ActivationSms(ctx context.Context, contryCode, to string, macros map[string]string) error {
	u.logger.Infow(ctx, "Processing Activation SMS", "to", to, "macros", macros)

	message := utils.ReplaceMacros(u.activationSmsTpl, macros)
	if err := u.smsSender.SendSMS(contryCode, to, message); err != nil {
		u.logger.Errorw(ctx, "Failed to send activation SMS", "error", err)
		return err
	}
	return nil
}

func (u *userusecase) PasswordResetEmail(ctx context.Context, to, subject string, macros map[string]string) error {
	u.logger.Infow(ctx, "Processing Password Reset Email", "to", to, "subject", subject, "macros", macros)

	if u.emailRateLimiter != nil {
		err := u.emailRateLimiter.WaitUntilAllowed(context.Background())

		if err != nil {
			u.logger.Errorw(ctx, "Failed to send password reset email due to rate limit error", "error", err)
			return err
		}
	}

	emailBody := utils.ReplaceMacros(u.passwordResetEmailTpl, macros)
	if err := u.emailer.SendEmail(to, subject, emailBody); err != nil {
		u.logger.Errorw(ctx, "Failed to send password reset email", "error", err)
		return err
	}
	return nil
}

func (u *userusecase) PasswordResetSms(ctx context.Context, contryCode, to string, macros map[string]string) error {
	u.logger.Infow(ctx, "Processing Password Reset SMS", "to", to, "contryCode", contryCode, "macros", macros)

	message := utils.ReplaceMacros(u.passwordResetSmsTpl, macros)
	if err := u.smsSender.SendSMS(contryCode, to, message); err != nil {
		u.logger.Errorw(ctx, "Failed to send password reset SMS", "error", err)
		return err
	}
	return nil
}
