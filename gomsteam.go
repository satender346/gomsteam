package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

//WebookURL : MSTeam webhook url of the channel
const WebhookURL = "paste msteam webhook url here"

func PostMSTeamMessage(MessageBody string) {

  method := "POST"

  payload := strings.NewReader(`{"text": "`+ MessageBody +`"}`)

  fmt.Println(payload)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, WebhookURL, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}

func main() {
  //Message can be a plan text or HTML formatted. Below example of HTML formatted.
  
  Message := "<p> Hello world </p> <br> <b>This is second line </b>"
  PostMSTeamMessage(Message)
}
