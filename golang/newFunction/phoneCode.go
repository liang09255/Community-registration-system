package newFunction

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendPhoneCode (args []*string, phoneNumber string, phoneCode string) (_err error, information string) {
	client, _err := CreateClient(tea.String("这里需要自己加配置信息"), tea.String("这里需要自己加配置信息"))
	if _err != nil {
		return _err,"nil"
	}

	codeMessage := "{code:" + phoneCode + "}"

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		TemplateCode: tea.String("这里需要自己加配置信息"),
		SignName: tea.String("这里需要自己加配置信息"),
		PhoneNumbers: tea.String(phoneNumber),
		TemplateParam: tea.String(codeMessage),
	}
	// 复制代码运行请自行打印 API 的返回值
	resp, _err := client.SendSms(sendSmsRequest)
	if _err != nil {
		return _err,"nil"
	}

	//下面自己改了下，以字符串的形式返回json中人看的message
	var message1 string
	message, ok := tea.ToMap(resp)["body"]
	if ok {
		message0, ok := message.(map[string]interface{})["Message"]
		if ok {
			message1 = message0.(string)
		}
	}
	return _err,message1
}