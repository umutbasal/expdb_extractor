package extract

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type info struct {
	CWE    cwe
	Impact DefImpact
}

type NVD map[string]info

func parseNVD() NVD {
	nvd := NVD{}
	cwe_db := parseCWE()
	err := filepath.Walk("../nvd",
		func(path string, i os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".json") {
				j, err := ioutil.ReadFile(path)
				if err != nil {
					panic(err)
				}
				var data NvdCveFeedJson11Schema
				err = json.Unmarshal(j, &data)
				if err != nil {
					panic(err)
				}

				for _, cve := range data.CVEItems {
					cwe_item := cwe{}
					for _, problemType := range cve.CVE.Problemtype.ProblemtypeData {
						for _, description := range problemType.Description {
							id := strings.TrimLeft(description.Value, "CWE-")
							cwe_item = queryCWEbyId(id, cwe_db)
						}
					}
					nvd[cve.CVE.CVEDataMeta.ID] = info{CWE: cwe_item, Impact: cve.Impact}
				}
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	return nvd
}

func queryNVDbyCve(cve string, nvd NVD) info {
	if cve == "" {
		return info{}
	}
	return nvd[cve]
}
