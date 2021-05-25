# gomsteam
This is the simplest implementation of gomsteams. It will send a message to the teams channel using webhook url with plain text or HTML in the message.

## Steps to get webhook url:
1. In MS Team channel, click on ... ( three dots )
2. Click on Connectors
3. Incoming Webhook 
4. Configure and Create new webhook.
5. Copy URL after creating webhook in same window.

## Run Code
1. Clone gomsteam repo 
2. Open gomsteam.go file in any editor
3. update webhook url.
4. Run below command 
```
go run gomsteam.go
```
5. verify message in team channel.
