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
