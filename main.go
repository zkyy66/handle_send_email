package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/jordan-wright/email"
	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
)

const (
	Sender            = "xxx"
	AuthorizationCode = "xxx"
	SmtpServer        = "smtp.exmail.qq.com"
	SendAddr          = SmtpServer + ":465"
)

type sheetCountEmail struct {
	Email string
	Code  string
}

func main() {
	ch := make(chan *sheetCountEmail, 10)

	var logger *log.Logger
	logPath := "./send_res.log"
	logFile, errs := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errs != nil {
		panic(errs)
	}

	fmt.Println("处理用户邮箱的excel...")
	filesTwo, err := excelize.OpenFile("./bindEmail.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()

	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()

	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()
	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()
	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()
	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()
	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()
	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()
	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()
	go func() {
		for v := range ch {
			handleMail(v, logger, logFile)
		}
	}()

	var emailCount []*sheetCountEmail

	valueRow, err := filesTwo.GetRows("SheetJS")
	for _, row := range valueRow {
		mailBox := &sheetCountEmail{
			Email: row[0],
			Code:  row[1],
		}
		emailCount = append(emailCount, mailBox)
		ch <- mailBox
		fmt.Printf("Email总数：%d；excel总数：%d\n", len(emailCount), len(valueRow))
	}
	fmt.Println("Email总数：", len(emailCount))
	select {}
}

// 处理发送邮件逻辑
func handleMail(mail *sheetCountEmail, logger *log.Logger, logFile *os.File) {
	randKey := generateRandomNumber(0, 5, 1)
	emailAccount := randomItem(int32(randKey[0]))
	fromUserEmail := emailAccount[0]     //邮箱账号
	authorizationCode := emailAccount[1] //邮箱授权码

	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "5E对战平台 <" + fromUserEmail + ">"
	// 设置接收方的邮箱
	e.To = []string{mail.Email}
	//设置主题
	e.Subject = "【5E对战平台】兔年春节福利"
	//设置文件发送的内容
	e.HTML = []byte(getEmailContext(mail.Code))
	////设置服务器相关的配置
	fmt.Println("fromUserEmail:", fromUserEmail)
	err := e.SendWithTLS(SendAddr, smtp.PlainAuth("", fromUserEmail, authorizationCode, SmtpServer), &tls.Config{ServerName: SmtpServer})

	log.Printf("开始发送-发送邮件用户:%s\n", mail.Email)
	logger = log.New(logFile, "[send_email_res]", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Printf("开始发送-发送邮件用户：%s\n", mail.Email)
	if err != nil {
		//发送失败
		logger.Printf("email %s;err：%s\n", mail.Email, err)
		log.Println("send error email:%s; %s\n", mail.Email, err)
	}
	// 发送成功
	logger.Printf("email %s；success：%s\n", mail.Email, err)
	log.Println("send success.......")
	time.Sleep(time.Second * 10)
}

// 随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end-start) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
func getEmailContext(exchangeCode string) string {
	sendContext := fmt.Sprintf(`
 <div style="margin: 24px 0; display: flex; justify-content: center">
      <div style="width: 799px">
        <div
          style="
            width: 799px;
            height: 898px;
            background-image: url(xxx);
            background-size: 799px 898px;
            overflow: hidden;
          "
        >
          <div
            style="
              width: 799px;
              margin-top: 610px;
              text-align: center;
              font-size: 24px;
              font-weight: bold;
              color: #ffffff;
              white-space: nowrap;
            "
          >
            %s
          </div>
        </div>
        <img
          style="width: 799px; display: block"
          src="xxx"
        />
      </div>
    </div>
    `, exchangeCode)
	return sendContext
}

func randomItem(key int32) []string {
	stringItem := [...][5]string{
		{"xxx.com", "xxx"},
		{"xxx.com", "xxx"},
		{"xxx.com", "xxx"},
		{"xxx.com", "xxx"},
		{"xxx.com", "xxx"},
	}
	emailItem := stringItem[key]
	return emailItem[:]
}
