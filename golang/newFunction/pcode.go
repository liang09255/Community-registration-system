package newFunction

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

var codeMap = make(map[string]string, 32)
var outdateMap = make(map[string]time.Time, 32)

func RandomCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000-100000) + 100000)
	return vcode
}

var judgeCheck = 0

func checkOutdate () {
	ticker := time.NewTicker(60 * time.Minute)
	for _ = range ticker.C {
		var deleteJudge = 0
		for k, v := range outdateMap{
			t := time.Now()
			if v.Before(t) {
				delete(outdateMap, k)
				delete(codeMap, k)
				deleteJudge += 1
			}
		}
		fmt.Println("定时任务：删除了", deleteJudge, "条过期验证码")
	}
}

func GetPhoneCode(w http.ResponseWriter, r *http.Request)  {

	randomCode := RandomCode()
	fmt.Println("随机验证码为:", randomCode)

	myPhone := r.FormValue("myPhone")

	//还要存个过期时间
	nowTime := time.Now()
	//获取五分钟后的过期时间s
	m, _ := time.ParseDuration("5m")
	outdateTime := nowTime.Add(m)

	//把过期时间也存个map
	outdateMap[myPhone] = outdateTime

	//map存储验证码
	codeMap[myPhone] = randomCode

	//发送手机验证码
	//省点钱把这个功能先注释掉
	err, information := SendPhoneCode(tea.StringSlice(os.Args[1:]), myPhone, randomCode)
	if err != nil {
		panic(err)
	}else{
		fmt.Println("获取手机验证码")
	}
	checkInformation, err := regexp.MatchString("流控", information)
	if err != nil {
		fmt.Println("checkInformation failed, err: ", err)
	}else if checkInformation {
		w.Write([]byte("so fast"))
	}

	if judgeCheck == 0 {
		judgeCheck = 1
		checkOutdate()
	}
}

func CheckCode(w http.ResponseWriter, r *http.Request)  {
	myCodeNumber := string(r.FormValue("myCodeNumber"))
	myPhoneNumber := r.FormValue("myPhoneNumber")

	fmt.Println("用户输入的验证码为:", myCodeNumber)

	outDateTime, ok := outdateMap[myPhoneNumber]
	if ok {
		nowTime := time.Now()
		if outDateTime.Before(nowTime) {
			w.Write([]byte("验证失败"))
			delete(codeMap, myCodeNumber)
			delete(outdateMap, myPhoneNumber)
			return
		}
	}else {
		w.Write([]byte("验证失败"))
		return
	}

	codeNumber, ok := codeMap[myPhoneNumber]
	if ok {
		if codeNumber == myCodeNumber {
			w.Write([]byte("验证成功"))
			delete(codeMap, myCodeNumber)
			delete(outdateMap, myPhoneNumber)
		}else {
			w.Write([]byte("验证失败"))
		}
	}
}