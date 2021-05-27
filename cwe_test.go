package extract

import (
	"log"
	"testing"
)

func Test_parseCWE(t *testing.T) {
	cwe := parseCWE()
	out := queryCWEbyId("200", cwe)
	log.Print(out)
}
