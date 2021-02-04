# Gin Multifactor Authentication


v 0.0.1


Flexible authentication for web, mobile, desktop and hybrid apps. Can be used for 1fa, 2fa and mfa scenarios. Easily configurable and extendable with new authentication methods, called `services`. All authenticaton scenarios, called `flows`, are based on `identifiers` and `secrets`, which can be used or not used in multiple combinations:
- username, email, phone, ...
- password, passcode (aka one-time pass or token), hardcode (aka device or card id), ...

Full list of supported services (devices):
- Email (soon)
- Phone (as Sms)
- WhatsApp (soon)
- Google Authenticator
- Microsoft Authenticator
- Authy, andOTP, etc
- Yubikey (soon)
- ...add yours

and service providers:  
- Twilio
- Vonage (Nexmo) (soon)
- Amazon SNS (soon)
- ...add yours



### Usage
See an example app in the `/example` folder.


```
// Init with specific flow(s):
// authenticate user if all (username, password, passcode) params are valid
auth := multauth.Auth{
        Flows: []multauth.Flow{{"Username", "Password", "Passcode"}},
}

app := gin.Default()

app.POST("/signin", func(c *gin.Context) {
        // ...Grab params from the context and store them in the "data" map

        err := auth.Authenticate(map[string]interface{}{
                "Username": data["username"],
                "Password": data["password"],
                "Passcode": data["passcode"], // with Google Authenticator or so
        }, user)

        if err == nil {
                c.JSON(200, gin.H{
                        "message": "Welcome " + user.Username,
                        "token": "YOUR_JWT_TOKEN",
                })
        }
})

app.Run()
```
