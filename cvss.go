package extract

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type info struct {
	Problemtype Problemtype
	Impact      DefImpact
}

type NVD map[string]info

func parseNVD() NVD {
	nvd := NVD{}
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
					nvd[cve.CVE.CVEDataMeta.ID] = info{Problemtype: cve.CVE.Problemtype, Impact: cve.Impact}
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
