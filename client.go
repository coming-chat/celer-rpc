package celerrpc

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RestClient struct {
	c      *http.Client
	rpcUrl string
}

func NewRestClient(host string) *RestClient {
	return &RestClient{
		rpcUrl: strings.TrimRight(host, "/"),
		c: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    3,
				IdleConnTimeout: 30 * time.Second,
			},
			Timeout: 30 * time.Second,
		},
	}
}

func (c *RestClient) GetTransferConfigs() (transferConfig *TransferConfig, err error) {
	transferConfig = &TransferConfig{}
	req, err := http.NewRequest("GET", c.rpcUrl+"/v2/getTransferConfigs", nil)
	if err != nil {
		return
	}
	err = c.doReq(req, transferConfig)
	return
}

func (c *RestClient) EstimateAmount(estimateAmountReq EstimateAmountReq) (res *EstimateAmountRes, err error) {
	req, err := http.NewRequest("GET", c.rpcUrl+"/v2/estimateAmt", nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("src_chain_id", strconv.Itoa(estimateAmountReq.SrcChainId))
	q.Add("dst_chain_id", strconv.Itoa(estimateAmountReq.DstChainId))
	q.Add("token_symbol", estimateAmountReq.TokenSymbol)
	q.Add("usr_addr", estimateAmountReq.UsrAddr)
	q.Add("slippage_tolerance", strconv.Itoa(estimateAmountReq.SlippageTolerance))
	q.Add("amt", estimateAmountReq.Amt)
	req.URL.RawQuery = q.Encode()
	res = &EstimateAmountRes{}
	err = c.doReq(req, res)
	return
}

// doReq send request and unmarshal response body to result
func (c *RestClient) doReq(req *http.Request, result interface{}) error {
	return doReqWithClient(req, result, c.c)
}

func doReqWithClient(req *http.Request, result interface{}, client *http.Client) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return handleResponse(result, resp)
}

// handleResponse read response data and unmarshal to result
// if http status code >= 400, function will return error with raw content
func handleResponse(result interface{}, resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		restError := RestError{}
		json.Unmarshal(body, &restError)
		restError.Code = resp.StatusCode
		return restError
	}
	return json.Unmarshal(body, &result)
}
