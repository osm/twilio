package twilio

import (
	"net/http"
	"net/url"
	"strings"
)

type twilio struct {
	accountSID string
	authToken  string
	from       string
}

func New(accountSID, authToken, from string) *twilio {
	return &twilio{accountSID, authToken, from}
}

func (t *twilio) SendSMS(to, body string) (int, error) {
	values := url.Values{}
	values.Set("Body", body)
	values.Set("From", t.from)
	values.Set("To", to)
	valuesReader := *strings.NewReader(values.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest(
		"POST",
		"https://api.twilio.com/2010-04-01/Accounts/"+t.accountSID+"/Messages.json",
		&valuesReader,
	)
	req.SetBasicAuth(t.accountSID, t.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	return resp.StatusCode, err
}
