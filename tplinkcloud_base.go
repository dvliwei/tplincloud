/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplinkcloud_basev2
 * @Version: 1.0.0
 * @Date: 2025/03/20 10:06
 * @Update liwei 2025/3/20 10:06
 */

package tplinkcloud

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type TplinkCloudBase struct {
	Ak         string
	Sk         string
	TerminalId string
	Path       string
	Payload    any
	PayloadStr string
}

func (res *TplinkCloudBase) PostRequest() ([]byte, error) {
	host := "api-smbcloud.tp-link.com.cn"
	path := res.Path
	authorization := res.authorization()

	ct := "application/json"
	payload := res.PayloadStr
	// 创建请求体
	body := bytes.NewBuffer([]byte(payload))
	// 构建请求 URL
	url := fmt.Sprintf("https://%s%s", host, path)
	// 创建 HTTP POST 请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("创建请求时出错:", err)
		return nil, err
	}
	// 设置请求头
	req.Header.Set("X-Authorization", authorization)
	req.Header.Set("Content-Type", ct)
	req.Header.Set("Host", host)
	// 创建 HTTP 客户端
	client := &http.Client{}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求时出错:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容时出错:", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("tplink cloud 响应失败")
	}
	return respBody, nil
}

func (res *TplinkCloudBase) authorization() string {
	var timestamp int64 = time.Now().Unix()
	nonce := getUUID()
	algorithm := "HmacSHA256"
	bt, _ := json.Marshal(res.Payload)
	payloadStr := string(bt)
	res.PayloadStr = payloadStr
	method := "POST"
	hashedRequestPayload := sha256hex(res.PayloadStr)
	credentialScope := fmt.Sprintf("%s %s tp-link_request", method, res.Path)
	// ************* 步骤 2：拼接待签名字符串 *************
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedRequestPayload)
	// ************* 步骤 3：计算签名 *************
	secretDate := hmacsha256(strconv.FormatInt(timestamp, 10), res.Sk)
	secretService := hmacsha256(res.Path, secretDate)
	secretSigning := hmacsha256("tp-link", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	// ************* 步骤 4：拼接 Authorization *************
	authorization := fmt.Sprintf("Timestamp=%d,Nonce=%s,AccessKey=%s,Signature=%s,TerminalId=%s",
		timestamp,
		nonce,
		res.Ak,
		signature,
		res.TerminalId)

	return authorization
}

func getUUID() (uuid string) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	uuid = fmt.Sprintf("%x", b)
	return
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
