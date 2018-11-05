package ga

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/glog"
)

// TODO: Should this be passed as build flags?
var (
	trackCode = "UA-127388617-1"
	UUID      = "7c004a96-de84-11e8-9f32-f2801f1b9fd1"
)

const analyticsURL = "https://www.google-analytics.com/collect"

// PushSingleEvent sends a single event in a POST request
func PushSingleEvent(eventCategory, eventAction, eventName, eventValue string) {
	queryParams := url.Values{
		// Version of Measurement protocol, default = 1
		"v": []string{"1"},

		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#tid
		"tid": []string{trackCode}, // constant code for tracking users for an application

		// hit type: any of
		// 'pageview', 'screenview', 'event', 'transaction', 'item', 'social', 'exception', 'timing'
		"t": []string{eventName}, // Hit Type, eg. install, volume_create

		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#dp
		"dp": []string{eventValue}, // Value of the custom eventA

		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#ds
		"ds": []string{"m-apiserver"}, // Data-source

		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#cid
		"cid": []string{UUID}, // UUID-version-1

		// Page title set to UUID
		"dt": []string{UUID},

		// Category of event
	}

	if eventCategory != "" {
		queryParams.Set("ec", eventCategory)
		queryParams.Set("ea", eventAction)
	}

	resp, err := http.PostForm(analyticsURL, queryParams)
	if err != nil {
		glog.Errorf(err.Error())
	} else {
		out, _ := ioutil.ReadAll(resp.Body)
		glog.Infof(string(out))
	}
}
