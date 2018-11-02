package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/glog"
)

// TODO: Should this be passed as build flags?
var (
	trackCode = "UA-127388617-1"
	uuid      = "7c004a96-de84-11e8-9f32-f2801f1b9fd1"
)

const analyticsURL = "https://www.google-analytics.com/collect"

// PushSingleEvent sends a single event in a POST request
func pushSingleEvent(eventName, eventValue string) {
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
		"cid": []string{uuid}, // uuid-version-1
	}

	resp, err := http.PostForm(analyticsURL, queryParams)
	if err != nil {
		glog.Errorf(err.Error())
	} else {
		out, _ := ioutil.ReadAll(resp.Body)
		glog.Infof(string(out))
	}
}

func main() {
	// Installation HIT, mocked by a page-view
	x := []string{"/" + uuid + "/1.0.0/cstore2/1.0.0/create",
		"/" + uuid + "/1.0.0/cstor/1.0.0/create",
		"/" + uuid + "/1.0.0/cstore2/1.0.0/delete",
		"/" + uuid + "/0.6.0/cstore2/1.0.0/create",
		"/" + uuid + "/1.0.0/cstore2/1.0.0/create",
		"/" + uuid + "/0.8.0/cstore2/1.0.0/create",
		"/" + uuid + "/0.9.0/cstore2/1.0.0/create",
	}
	for _, i := range x {
		pushSingleEvent("pageview", i)
	}
}
