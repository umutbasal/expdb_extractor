package extract

import (
	"reflect"
	"testing"
)

func Test_parseCWE(t *testing.T) {
	CWE := parseCWE()
	actual := queryCWEbyId("200", CWE)
	expect := cwe{
		ID:      200,
		Name:    "Exposure of Sensitive Information to an Unauthorized Actor",
		Related: "::NATURE:ChildOf:CWE ID:668:VIEW ID:1000:ORDINAL:Primary::",
	}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("expected() = %v, want %v", expect, actual)
	}
}
