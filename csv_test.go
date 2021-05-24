package extract

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_csvRead(t *testing.T) {
	csvFile, err := os.Open("../files_exploits.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	out := csvRead(csvFile)
	if len(out) < 1 {
		t.Error("no data")
	}
}

func Test_getRowByFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantRow Entry
		wantErr bool
	}{
		{"example file", args{fileName: "exploits/linux/dos/11.c"}, Entry{Id: "11", File: "exploits/linux/dos/11.c", Description: "Apache 2.0.44 (Linux) - Remote Denial of Service", Date: "2003-04-11", Author: "Daniel Nystram", Type: "dos", Platform: "linux"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvFile, err := os.Open("../files_exploits.csv")
			db := csvRead(csvFile)
			if err != nil {
				fmt.Println(err)
			}
			defer csvFile.Close()

			gotRow, err := getRowByFileName(db, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRowByFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRow, tt.wantRow) {
				t.Errorf("getRowByFileName() = %v, want %v", gotRow, tt.wantRow)
			}
		})
	}
}
