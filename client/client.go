package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	token      string
	apiTimeout int
}

func (c Client) GetUpdates(offset int) ([]Update, error) {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", c.token)

	payload := url.Values{
		"offset":  {strconv.Itoa(offset)},
		"timeout": {strconv.Itoa(c.apiTimeout)},
	}
	data := bytes.NewBufferString(payload.Encode())
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response GetUpdatesReturn
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if !response.Ok {
		return nil, errors.New("Cannot read the last updates, try again")
	}

	return response.Result, nil
}

func (c Client) SendMessage(chatId, messageId int, message string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.token)

	payload := url.Values{
		"chat_id":             {strconv.Itoa(chatId)},
		"reply_to_message_id": {strconv.Itoa(messageId)},
		"text":                {message},
	}
	data := bytes.NewBufferString(payload.Encode())

	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response SendMessageReturn
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if !response.Ok {
		return errors.New(string(body))
	}

	return nil
}

func New(token string) *Client {
	return &Client{token: token, apiTimeout: 60}
}
