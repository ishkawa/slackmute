package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetChannels returns ids that the authenticated user joins.
func GetChannels(token string) ([]string, error) {
	form := url.Values{
		"token": {token},
	}

	resp, err := http.Get("https://slack.com/api/channels.list?" + form.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code " + string(resp.StatusCode))
	}

	var channels ListChannelsResponse
	err = json.NewDecoder(resp.Body).Decode(&channels)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, channel := range channels.Channels {
		ids = append(ids, channel.ID)
	}

	return ids, nil
}

// MuteChannels calls mute API for ids.
func MuteChannels(token string, ids []string) error {
	form := url.Values{
		"token": {token},
		"name":  {"muted_channels"},
		"value": {strings.Join(ids, ",")},
	}

	resp, err := http.PostForm("https://slack.com/api/users.prefs.set", form)
	if err != nil {
		return err
	}
	d, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(d))

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status code " + string(resp.StatusCode))
	}

	return nil
}
