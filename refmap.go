package extract

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html/charset"
)

type cve struct {
	Id   string `json:"id"`
	Desc string `json:"desc"`
	NVD  info   `json:"nvd"`
}

type eidCveMap map[string][]cve

func ParseExploitCvesMap() (ecm eidCveMap, err error) {
	eidCveMap := eidCveMap{}
	cveXML, err := ioutil.ReadFile("../allitems.xml")
	nvd := parseNVD()
	if err != nil {
		panic(err)
	}

	var mitreCve Cve
	decoder := xml.NewDecoder(bytes.NewReader(cveXML))
	decoder.CharsetReader = charset.NewReaderLabel
	if err = decoder.Decode(&mitreCve); err != nil {
		panic(fmt.Errorf("Failed to Unmarshal XML: err: %s", err))
	}
	for _, vuln := range mitreCve.Item {
		for _, ref := range vuln.Refs.Ref {
			if strings.HasPrefix(ref.Url, "https://www.exploit-db.com/exploits/") {
				eid := strings.TrimSuffix(strings.TrimPrefix(ref.Url, "https://www.exploit-db.com/exploits/"), "/")
				cvss := queryNVDbyCve(vuln.Name, nvd)
				eidCveMap[eid] = append(eidCveMap[eid], cve{Id: vuln.Name, Desc: vuln.Desc, NVD: cvss})
			}
		}
	}
	return eidCveMap, nil
}

func queryRefMap(eidCveMap eidCveMap, exploitID string) []cve {
	return eidCveMap[exploitID]
}
