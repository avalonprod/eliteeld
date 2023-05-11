package emails

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Emails struct {
	apiUrlCompanyRegistration string
	apiUrlDriverRegistration  string
}

func NewEmails(apiUrlCompanyRegistration string, apiUrlDriverRegistration string) *Emails {
	return &Emails{
		apiUrlCompanyRegistration: apiUrlCompanyRegistration,
		apiUrlDriverRegistration:  apiUrlDriverRegistration,
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

	req, err := http.NewRequest("POST", e.apiUrlCompanyRegistration, bytes.NewBuffer(payload))
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

func (e *Emails) SendEmailDriverRegistration(email string, name string, companyName string, companyEmail string) error {

	data := map[string]string{
		"email":        email,
		"name":         name,
		"companyName":  companyName,
		"companyEmail": companyEmail,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", e.apiUrlDriverRegistration, bytes.NewBuffer(payload))
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
