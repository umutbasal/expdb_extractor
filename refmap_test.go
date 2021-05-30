package extract

import (
	"testing"
)

func TestParseExploitCvesMap(t *testing.T) {
	eidCveMap, err := ParseExploitCvesMap()
	if err != nil {
		panic(err)
	}
	actual := queryRefMap(eidCveMap, "8545")[0].Id
	expect := "CVE-2009-1623"
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}

}
