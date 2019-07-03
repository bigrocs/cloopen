# www.yuntongxun.com 
# cloopen 容联云通讯 
## sms 短信 golang sdk
```
    cloopen := &Cloopen{
		AccountSid:   "********",
		AppID:        "********",
		AccountToken: "********",
	}
	req := &Request{
		Mobile:       "15512345678",
		TemplateCode: "453946",
		Datas: []string{
			"123456",
			"30",
		},
	}
	valid, err := cloopen.Send(req)
	fmt.Println(valid, err)
```
- 详情请查看 test 车市文件