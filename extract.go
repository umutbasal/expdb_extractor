package extract

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func read(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func write(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	return err
}

type Eid struct {
	Info     Entry    `json:"info"`
	Requests []string `json:"requests"`
}

func Walker() []Eid {
	csvFile, err := os.Open("../files_exploits.csv")
	db := csvRead(csvFile)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	out := []Eid{}
	err = filepath.Walk("../",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasPrefix(path, "../exploits") {
				fileName := strings.TrimPrefix(path, "../")
				row, err := getRowByFileName(db, fileName)
				if err != nil {
					panic(err)
				}
				entry := Eid{Info: row, Requests: Extract(read(path))}
				if len(entry.Requests) > 0 {
					out = append(out, entry)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return out
}

func Extract(content string) []string {
	content = content + "\n\n"
	content = strings.Replace(content, "\r\n", "\n", -1)
	content = strings.Replace(content, "\n", "\\n", -1)
	r, _ := regexp.Compile(`\\n(GET|POST|PUT|PATCH|DELETE) \/.*?\\n\\n`)
	res := r.FindAllString(content, -1)
	for k, v := range res {
		res[k] = strings.Replace(v, "\\n", "\n", -1)
		res[k] = strings.TrimSpace(res[k])
	}
	return res
}
