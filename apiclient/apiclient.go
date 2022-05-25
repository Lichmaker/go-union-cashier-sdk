package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"

	"strings"
	"time"

	"github.com/lichmaker/go-union-cashier-sdk/apirequest"
	"github.com/lichmaker/go-union-cashier-sdk/signaturebuilder"
	"github.com/pkg/errors"
)

const (
	VERSION  = "1.0.1"
	SIGN_ALG = 1
)

type Client struct {
	Conf Config
}

type Config struct {
	Host           string
	AppID          string
	MchId          string
	PrivateKeyPath string
	PublicKeyPath  string
	BizType        string
	Timeout        time.Duration
}

type RequestBody struct {
	App_id      string `json:"app_id"`
	Method      string `json:"method"`
	Timestamp   string `json:"timestamp"`
	V           string `json:"v"`
	Sign_alg    int    `json:"sign_alg"`
	Biz_content string `json:"biz_content"`
}

func NewClient(c Config) Client {
	return Client{
		Conf: c,
	}
}

func (c Client) Do(req apirequest.Request) (map[string]interface{}, error) {
	body := RequestBody{
		App_id:      c.Conf.AppID,
		Method:      req.ApiInterfaceId + "." + req.MethodName,
		Timestamp:   time.Now().Format("2006-01-02 15:06:07"),
		V:           VERSION,
		Sign_alg:    SIGN_ALG,
		Biz_content: req.BizContent,
	}
	bodyString := BuildBodyString(body)
	sign, err := signaturebuilder.Sign(bodyString, c.Conf.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	bodyString = fmt.Sprintf("%s&sign=%s", bodyString, sign)

	httpClient := new(http.Client)
	httpClient.Timeout = time.Second * 5

	httpReq, _ := http.NewRequest("POST", c.Conf.Host, strings.NewReader(bodyString))
	httpReq.Header.Add("Content-type", "test/json")
	httpReq.Header.Add("charset", "utf8")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrapf(err, "请求失败")
	}

	defer resp.Body.Close()
	dataByte, _ := ioutil.ReadAll(resp.Body)
	respBody := make(map[string]interface{})
	json.Unmarshal(dataByte, &respBody)
	return respBody, nil
}

func (c Client) GenMerOrdrNo() string {
	// 商户交易订单号要求32位[15位商户号+8位交易日期+9位数字序列号]
	// 9位数字，为了避免冲突，使用时+分+秒+3位随机数 例：140227123
	mchId := c.Conf.MchId
	dateTime := time.Now().Format("20060102150405")
	return mchId + dateTime + RandomNumberString(3)
}

func BuildBodyString(body RequestBody) string {
	b := url.Values{}
	b.Set("app_id", body.App_id)
	b.Set("method", body.Method)
	b.Set("timestamp", body.Timestamp)
	b.Set("v", body.V)
	b.Set("sign_alg", fmt.Sprintf("%d", body.Sign_alg))
	b.Set("biz_content", body.Biz_content)
	return b.Encode()
}

func RandomNumberString(strLen int) string {
	rand.Seed(time.Now().UnixNano())
	full := "0123456789"
	arr := strings.Split(full, "")
	genStr := ""
	for i := 0; i < strLen; i++ {
		randNum := rand.Intn(len(arr))
		genStr = genStr + arr[randNum]
	}
	return genStr
}
