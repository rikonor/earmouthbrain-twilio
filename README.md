EarMouthBrain - Twilio
---

##### TwilioSMSMouth

Use to send SMS messages:
```
import twilio "github.com/rikonor/earmouthbrain-twilio"

twilioSMSConfig := twilio.TwilioSMSConfig{
  FromNumber: "+12345678900",
  AccountSID: "AccountSID",
  AuthToken:  "AuthToken",
}

phoneNumbers := []string{
  "(123) 456-7890",
}

sm := twilio.NewTwilioSMSMouth(twilioSMSConfig, phoneNumbers...)

msg := emb.StringToMessage("Test")
sm.Say(msg)
```

Register as message handler:
```
import (
  "github.com/rikonor/earmouthbrain-common/file"
  twilio "github.com/rikonor/earmouthbrain-twilio"
)

fe := file.NewFileEar("./input")

twilioSMSConfig := twilio.TwilioSMSConfig{
  FromNumber: "+12345678900",
  AccountSID: "AccountSID",
  AuthToken:  "AuthToken",
}

phoneNumbers := []string{
  "(123) 456-7890",
}

sm := twilio.NewTwilioSMSMouth(twilioSMSConfig, phoneNumbers...)
fe.RegisterMessageHandler(sm.Say)

fmt.Scanln()
```
