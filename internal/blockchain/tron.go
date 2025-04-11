package blockchain

import (
	"DigitalCurrency/internal/blockchain/model"
	"DigitalCurrency/internal/config"
	"DigitalCurrency/internal/util"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Tron struct {
	Blockchain
	httpClient http.Client
}

var (
	tron     Tron
	tronOnce sync.Once
)

func NewTron(ctx context.Context, conf *config.EthConfig) *Tron {
	tronOnce.Do(func() {
		tron = Tron{
			Blockchain{
				Chain:  "tron",
				Config: conf,
				ctx:    ctx,
			},
			http.Client{
				Timeout: 60 * time.Second, // 设置超时时间
			},
		}
	})
	return &tron
}

// GetNowBlockId 获取当前区块 ID
// @param blockID: 可选的区块 ID，若未提供则获取最新区块
// @return: 当前区块 ID, 当前时间, 错误信息
func (t *Tron) GetNowBlockId(blockID ...int) (int, time.Time, error) {
	url := "/walletsolidity/getblock"
	method := "POST"
	params := make(map[string]any) // 请求参数
	params["detail"] = false
	if len(blockID) > 0 && blockID[0] > 0 { // 如果提供了区块 ID
		params["id_or_num"] = strconv.Itoa(blockID[0]) // 将 ID 转为字符串
	}
	res, err := t.request(method, url, params) // 发送请求
	if err != nil {
		return 0, time.Time{}, err // 返回错误
	}
	var data model.GetNowBlock
	err = json.Unmarshal(res, &data) // 解析响应
	if err != nil {
		return 0, time.Time{}, err // 返回错误
	}
	if data.BlockHeader.RawData.Number == 0 && data.Error != "" {
		return 0, time.Time{}, errors.New(data.Error) // 返回错误信息
	}
	return data.BlockHeader.RawData.Number, util.GetDateByTimestamp(data.BlockHeader.RawData.Timestamp), nil // 返回区块 ID 和时间
}

func (t *Tron) GetBlockByNum(num int) (model.GetBlockByNum, error) {
	url := "/walletsolidity/getblockbynum"
	params := make(map[string]any) // 请求参数
	params["num"] = num
	res, err := t.request("POST", url, params) // 发送请求
	var data model.GetBlockByNum
	if err != nil {
		return data, err // 返回错误
	}
	err = json.Unmarshal(res, &data) // 解析响应
	return data, err                 // 返回响应体和错误信息
}

// request 发送 HTTP 请求
// @param method: HTTP 方法 (GET, POST 等)
// @param action: 请求的 URL 路径
// @param param: 请求参数
// @return: 响应体, 错误信息
func (t *Tron) request(method string, action string, param any) ([]byte, error) {
	jsonData, err := json.Marshal(param) // 将参数序列化为 JSON
	if err != nil {
		return nil, err // 返回错误
	}
	reqBody := bytes.NewBuffer(jsonData)              // 创建请求体
	uri := fmt.Sprintf("%v%v", t.Config.Node, action) // 构建完整 URL
	req, err := http.NewRequest(method, uri, reqBody) // 创建 HTTP 请求
	if len(t.Config.ApiKey) > 0 {
		req.Header.Add("TRON-PRO-API-KEY", t.Config.ApiKey) // 添加 API 密钥
	}
	if err != nil {
		return nil, err // 返回错误
	}
	// proxyAddress, _ := url.Parse("http://10.0.5.124:45613/")
	// t.httpClient.Transport = &http.Transport{
	// 	Proxy: http.ProxyURL(proxyAddress),
	// }
	httpRes, err := t.httpClient.Do(req) // 发送请求
	if err != nil {
		return nil, err // 返回错误
	}
	defer httpRes.Body.Close()            // 确保在函数结束时关闭响应体
	body, err := io.ReadAll(httpRes.Body) // 读取响应体
	return body, err                      // 返回响应体和错误信息
}
