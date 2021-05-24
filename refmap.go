package extract

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html/charset"
)

type eidCveMap map[string][]string

func ParseExploitCvesMap() (eidCveMap eidCveMap, err error) {
	cveXML, err := ioutil.ReadFile("../allitems.xml")
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
				eidCveMap[eid] = append(eidCveMap[eid])
				continue
			}

			ss := strings.Split(ref.Text, ":")
			if len(ss) != 2 {
				continue
			}
			refType, exploitID := ss[0], ss[1]
			// https://cve.mitre.org/data/refs/index.html
			if refType != "EXPLOIT-DB" {
				continue
			}
			eidCveMap[exploitID] = append(eidCveMap[exploitID], vuln.Name)
		}
	}
	return eidCveMap, nil
}

func queryRefMap(eidCveMap eidCveMap, exploitID string) []string {
	return eidCveMap[exploitID]
}
