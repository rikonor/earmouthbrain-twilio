EarMouthBrain - Twilio
---

##### TwilioSMSMouth

Use to send SMS messages:
```
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
```
