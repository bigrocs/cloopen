package cloopen

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	cloopen := &Cloopen{
		AccountSid:   "********",
		AppID:        "********",
		AccountToken: "********",
	}
	req := &Request{
		Mobile:       "15550251272",
		TemplateCode: "453946",
		Datas: []string{
			"123456",
			"30",
		},
	}
	valid, err := cloopen.Send(req)
	fmt.Println(valid, err)
}
