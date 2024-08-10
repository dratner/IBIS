# IBIS
![IBIS image](static/IBIS.png)

IBIS - The Injured Bird Information System 

This app uses Twilio for text message communication and uses the [Twilio Go Helper Library](https://www.twilio.com/en-us/blog/introducing-twilio-go-helper-library)

## Interface

Any volunteer can text the system using the following messages:

```on``` - this tells the system to forward relevant messages

```off``` - this tells the system to stop forwarding messages

```status``` - replies with your current status (on/off)

```add <keyword>``` - adds a keyword for filtering

```remove <keyword>``` - removes a keyword for filtering

```all``` - overrides filtering and forwards all messages (wildcard)

```keywords``` - lists your current keywords

```register <number>``` - adds a new volunteer with the given number

```delete <number>``` - removes a volunteer with the given number

```block <number>``` - blocks a spammer

```unblock <number>``` - unblocks a spammer


## Environment Variables

```TWILIO_ACCOUNT_SID``` - The account SID issued by Twilio

```TWILIO_ACCOUNT_TOKEN``` - The secret token for the account issued by Twilio

```TWILIO_ACCOUNT_NUMBER``` - The inbound phone number for text messages
