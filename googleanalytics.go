package main

import (
	"net/http"
	"net/url"

	"github.com/golang/glog"
)

// TODO: Should this be passed as build flags?
var trackCode = "UA-127388617-1"

const analyticsURL = "https://www.google-analytics.com/collect"

// PushSingleEvent sends a single event in a POST request to OpenEBS core-developers
func pushSingleEvent(eventName, eventValue string) {
	// uid, err := getAnonymousID()
	queryParams := url.Values{
		"v": []string{"1"}, // Version of Measurement protocol. default = 1
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
		"cid": []string{"7c004a96-de84-11e8-9f32-f2801f1b9fd1"}, // uuid-version-1
	}

	_, err := http.PostForm(analyticsURL, queryParams)
	if err != nil {
		glog.Errorf(err.Error())
	} else {
		// TODO: Remove before merging code?
		glog.Infof("Sent request to GA")
	}
}

func pushCustomEvent(mKey, mValue string) {
	queryParams := url.Values{
		"v": []string{"1"}, // Version of Measurement protocol. default = 1
		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#tid
		"tid": []string{trackCode}, // constant code for tracking users for an application

		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#cid
		"cid": []string{"7c004a96-de84-11e8-9f32-f2801f1b9fd1"}, // uuid-version-1

		// key -> value
		mKey: []string{mValue},
	}
	_, err := http.PostForm(analyticsURL, queryParams)
	if err != nil {
		glog.Errorf(err.Error())
	} else {
		// TODO: Remove before merging code?
		glog.Infof("Sent request to GA")
	}

}

func main() {
	// Installation HIT, mocked by a page-view
	pushSingleEvent("pageview", "/openebs/installed")

	// Event HITS mocked by up to 20 events
	x := map[string]string{
		"cm1": "100",  // We can create 20 custom metrics in Google Analytics
		"cm2": "1000", // number of cstor volumes
		"cm3": "99",
	}

	for eventName, eventValue := range x {
		pushCustomEvent(eventName, eventValue)
	}
}
