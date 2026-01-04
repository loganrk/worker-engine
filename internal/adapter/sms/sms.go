package sms

type sms struct {
}

func New() *sms {
	return &sms{}
}

func (s *sms) SendSms(contryCode, to, message string) error {

	return nil
}
