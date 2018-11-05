package main

import (
	"github.com/golang/glog"
	"github.com/vharsh/ga-test/pkg/ga"
)

func main() {
	// Installation HIT, mocked by a page-view
	glog.Infof("using uuid: %s", ga.UUID)
	x := []string{"/" + ga.UUID + "/1.0.0/cstore2/1.0.0/create",
		"/" + ga.UUID + "/1.0.0/cstor/1.0.0/create",
		"/" + ga.UUID + "/1.0.0/cstore2/1.0.0/delete",
		"/" + ga.UUID + "/0.6.0/cstore2/1.0.0/create",
		"/" + ga.UUID + "/1.0.0/cstore2/1.0.0/create",
		"/" + ga.UUID + "/0.8.0/cstore2/1.0.0/create",
		"/" + ga.UUID + "/0.9.0/cstore2/1.0.0/create",
	}
	for _, i := range x {
		ga.PushSingleEvent("pageview", i)
	}
}
