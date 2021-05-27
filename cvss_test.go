package extract

import (
	"log"
	"testing"
)

func Test_parseNVD(t *testing.T) {
	nvd := parseNVD()
	out := queryNVDbyCve("CVE-2021-32925", nvd)
	log.Print(out)
}
