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
		"v":   []string{"1"},                                    // Version of Measurement protocol
		"tid": []string{trackCode},                              // a constant
		"t":   []string{eventName},                              // Hit Type, eg. install, volume_create
		"dp":  []string{eventValue},                             // Value of the custom event
		"ds":  []string{"m-apiserver"},                          // Data-source
		"cid": []string{"7c004a96-de84-11e8-9f32-f2801f1b9fd1"}, // uuid-version-1
	}

	resp, err := http.PostForm(analyticsURL, queryParams)
	if err != nil {
		glog.Errorf(err.Error())
	} else {
		// TODO: Remove before merging code?
		glog.Infof("Sent request to GA")
	}
}

func main() {
	x := map[string]string{
		"pageview": "/openebs/installed",
		"cm1":      "100",  // We can create 20 custom metrics in Google Analytics
		"cm2":      "1000", // number of cstor volumes
		"cm3":      "99",
	}

	for eventName, eventValue := range x {
		pushSingleEvent(eventName, eventValue)
	}
}
