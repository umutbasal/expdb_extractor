package extract

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type cwe struct {
	ID      int    `json:"CWE-ID"`
	Name    string `json:"Name"`
	Related string `json:"Related Weaknesses"`
}

type CWE map[string]cwe

func parseCWE() CWE {
	CWE := CWE{}
	j, err := ioutil.ReadFile("../635.json")
	if err != nil {
		panic(err)
	}
	var data []cwe
	err = json.Unmarshal(j, &data)
	if err != nil {
		panic(err)
	}
	for _, cve := range data {
		id := strconv.Itoa(cve.ID)
		CWE[id] = cwe{ID: cve.ID, Name: cve.Name, Related: cve.Related}
	}
	return CWE
}

func queryCWEbyId(cve string, CWE CWE) cwe {
	if cve == "" {
		return cwe{}
	}
	return CWE[cve]
}
