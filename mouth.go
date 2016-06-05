package twilio

import (
	emb "github.com/rikonor/earmouthbrain"
	"github.com/subosito/twilio"
)

type TwilioSMSMouth struct {
	emb.Mouth

	// Phone - Example "(203) 451-1578"
	PhoneNumbers []string

	// Twilio API client
	twilioAPI *twilio.Client

	// Sender number (basically your twilio number)
	fromNumber string
}

type TwilioSMSConfig struct {
	FromNumber string
	AccountSID string
	AuthToken  string
}

func NewTwilioSMSMouth(cfg TwilioSMSConfig, phoneNumbers ...string) *TwilioSMSMouth {
	twilioAPI := twilio.NewClient(cfg.AccountSID, cfg.AuthToken, nil)

	m := TwilioSMSMouth{
		PhoneNumbers: phoneNumbers,
		twilioAPI:    twilioAPI,
		fromNumber:   cfg.FromNumber,
	}

	m.Init(m.OutputToSMS)
	return &m
}

func (m *TwilioSMSMouth) OutputToSMS(msg emb.Message) {
	for _, phoneNumber := range m.PhoneNumbers {
		m.twilioAPI.Messages.SendSMS(m.fromNumber, phoneNumber, msg.String())
	}
}
