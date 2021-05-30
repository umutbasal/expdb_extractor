package extract

import (
	"testing"
)

func Test_parseNVD(t *testing.T) {
	NVD := parseNVD()
	actual := queryNVDbyCve("CVE-2021-32925", NVD).CWE.ID
	expect := 200
	if expect != actual {
		t.Errorf("expected %v, got %v", expect, actual)
	}
}
