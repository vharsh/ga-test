package main

import (
	"github.com/golang/glog"
	"github.com/vharsh/ga-test/pkg/ga"
)

func main() {
	// Installation HIT, mocked by a page-view
	glog.Infof("using uuid: %s", ga.UUID)
	x := []string{
		"/" + ga.UUID + "/install/v1.8/0.7.0/linux/gke",
		"/" + ga.UUID + "/vc/cstor/0.7.0",
		"/" + ga.UUID + "/vc/jiva/0.5.0",
		"/" + ga.UUID + "/vd/jiva/0.5.0",
		"/" + ga.UUID + "/vd/cstor/0.7.0",
	}
	for _, i := range x {
		ga.PushSingleEvent("", "pageview", i)
	}
}
