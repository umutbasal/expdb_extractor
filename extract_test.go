package extract

import (
	"reflect"
	"testing"
)

func testData() string {
	return read("testdata")
}

func Test_extract(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"should extract raw requests", args{content: testData()}, []string{
			`GET /auth/oauth/authorize?response_type=code&client_id=R2dpxQ3vPrtfgF72&scope=user_info&state=HPRbfRgJLWdmLMi9KXeLJDesMLfPC3vZ0viEkeIvGuQ%3D&redirect_uri=http://localhost:8086/login/oauth2/code/ HTTP/1.1`,
			`GET /auth/oauth/authorize?response_type=code&client_id=R2dpxQ3vPrtfgF72&scope=user_info&state=HPRbfRgJLWdmLMi9KXeLJDesMLfPC3vZ0viEkeIvGuQ%3D&redirect_uri=http://%localhost:9000/login/oauth2/code/ HTTP/1.1
Host: localhost:8085
User-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Referer: http://localhost:8085/auth/login
Connection: close
Cookie: JSESSIONID=3394FD89204BE407CB585881755C0828; JSESSIONID=C0F1D5A2F1944DCB43F2BFFA416B7A63
Upgrade-Insecure-Requests: 1`,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

// //Runner
// func Test_Walker(t *testing.T) {
// 	out := Walker()
// 	if cap(out) < 1 {
// 		t.Errorf("no file found %v", out)
// 	}
// 	j, err := json.Marshal(out)
// 	write("out.json", j)
// 	if err != nil {
// 		panic(err)
// 	}
// }
