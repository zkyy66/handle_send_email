package main

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

const (
	Sender            = "activity@5eplay.com"
	AuthorizationCode = "VDMSmoQE92KnewP2"
	SmtpServer        = "smtp.exmail.qq.com"
	SendAddr          = SmtpServer + ":465"
)

func main() {
	//handleMail()
	handleExcel()
}

// 主题标题
func themeTitle(key int32) string {
	var themeArr = []string{
		"【5E对战平台】元旦福利大放送",
		"【5E对战平台】元旦盛典福利",
		"【5E对战平台】元旦三重活动赢豪礼",
		"【5E对战平台】元旦盛典享三重豪礼",
		"【5E对战平台】元旦盛典享豪礼",
		"【5E对战平台】元旦福利大放送",
		"【5E对战平台】元旦盛典福利",
		"【5E对战平台】元旦三重活动赢豪礼",
		"【5E对战平台】元旦盛典享三重豪礼",
		"【5E对战平台】元旦盛典享豪礼",
	}
	return themeArr[key]
}

// 处理发送邮件逻辑
func handleMail(toUserEmail string, logger *log.Logger, logFile *os.File) {

	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "5E对战平台 <activity@5eplay.com>"
	// 设置接收方的邮箱
	e.To = []string{toUserEmail}
	//设置主题
	randKey := generateRandomNumber(0, 10, 1)
	subjectTitle := themeTitle(int32(randKey[0]))
	e.Subject = subjectTitle
	//设置文件发送的内容
	e.HTML = []byte(`
					<img data-imagetype="External" src="https://oss.5eplay.com/images/act/3db413d2c02246088b6efde4c443ec9e.jpg" class="x_fullMobileWidth" width="620" alt="5E对战平台" title="5E对战平台" style="display:block; height:auto; border:0; width:700px; max-width:100%">`)
	//设置服务器相关的配置
	err := e.SendWithTLS(SendAddr, smtp.PlainAuth("", Sender, AuthorizationCode, SmtpServer), &tls.Config{ServerName: SmtpServer})

	log.Printf("开始发送-发送邮件用户:%s\n", toUserEmail)
	logger = log.New(logFile, "[send_email_res]", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Printf("开始发送-发送邮件用户：%s\n", toUserEmail)
	if err != nil {
		//发送失败
		logger.Printf("email %s;err：%s\n", toUserEmail, err)
		log.Fatalf("send error email:%s; %s\n", toUserEmail, err)
	}
	//发送成功
	logger.Printf("email %s；success：%s\n", toUserEmail, err)
	log.Println("send success.......")
}

func handleExcel() {

	var logger *log.Logger
	logPath := "./send_res.log"
	logFile, errs := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errs != nil {
		panic(errs)
	}

	fmt.Println("处理用户邮箱的excel...")
	filesTwo, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var sheetCountEmail []string
	valueRow, err := filesTwo.GetRows("Sheet1")
	for _, row := range valueRow {
		for _, colCell := range row {
			sheetCountEmail = append(sheetCountEmail, colCell)
			handleMail(colCell, logger, logFile)
		}
	}
	fmt.Println("Email总数：", len(sheetCountEmail))

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
