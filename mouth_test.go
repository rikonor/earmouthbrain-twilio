package twilio

import (
	"fmt"
	"time"

	emb "github.com/rikonor/earmouthbrain"
	"github.com/rikonor/earmouthbrain-common/file"
)

func ExampleTwilioSMSMouth() {
	twilioSMSConfig := TwilioSMSConfig{
		FromNumber: "+12345678900",
		AccountSID: "AccountSID",
		AuthToken:  "AuthToken",
	}

	phoneNumbers := []string{
		"(123) 456-7890",
	}

	sm := NewTwilioSMSMouth(twilioSMSConfig, phoneNumbers...)

	msg := emb.StringToMessage("Test")
	sm.Say(msg)

	// Give mouth time to process message
	time.Sleep(1 * time.Millisecond)
}

func ExampleTwilioSMSMouth_from_ear() {
	fe := file.NewFileEar("./input")

	twilioSMSConfig := TwilioSMSConfig{
		FromNumber: "+12345678900",
		AccountSID: "AccountSID",
		AuthToken:  "AuthToken",
	}

	phoneNumbers := []string{
		"(123) 456-7890",
	}

	sm := NewTwilioSMSMouth(twilioSMSConfig, phoneNumbers...)
	fe.RegisterMessageHandler(sm.Say)

	fmt.Scanln()
}
