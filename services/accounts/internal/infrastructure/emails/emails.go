package emails

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Emails struct {
	apiUrl string
}

func NewEmails(apiUrl string) *Emails {
	return &Emails{
		apiUrl: apiUrl,
	}
}

func (e *Emails) SendEmailCompanyRegistration(email string, name string) error {

	data := map[string]string{
		"email": email,
		"name":  name,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", e.apiUrl, bytes.NewBuffer(payload))
	if err != nil {

		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return err
	}

	defer resp.Body.Close()

	return nil
}
