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
	//out := queryRefMap(eidCveMap, "11")
	log.Println(eidCveMap)
}
