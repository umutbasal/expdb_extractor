package extract

import (
	"log"
	"testing"
)

func TestParseExploitCvesMap(t *testing.T) {
	eidCveMap, err := ParseExploitCvesMap()
	if err != nil {
		panic(err)
	}
	out := queryRefMap(eidCveMap, "8545")
	log.Println(out)
}
