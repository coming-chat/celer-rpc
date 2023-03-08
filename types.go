package celerrpc

import "fmt"

type RestError struct {
	Code int
}

func (e RestError) Error() string {
	return fmt.Sprintf("{\"code\":%d}", e.Code)
}

type TransferConfig struct {
	Err                       Err                       `json:"err"`
	Chains                    []Chain                   `json:"chains"`
	ChainToken                map[string]ChainTokenInfo `json:"chain_token"`
	FarmingRewardContractAddr string                    `json:"farming_reward_contract_addr"`
	PeggedPairConfigs         []PeggedPairConfig        `json:"pegged_pair_configs"`
}

type Chain struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	BlockDelay     uint   `json:"block_delay"`
	GasTokenSymbol string `json:"gas_token_symbol"`
	ExploreUrl     string `json:"explore_url"`
	ContractAddr   string `json:"contract_addr"`
}

type ChainTokenInfo struct {
	TokenList []TokenInfo `json:"token"`
}

type TokenInfo struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Token Token  `json:"token"`
}

type Token struct {
	Symbol       string `json:"symbol"`
	Address      string `json:"address"`
	Decimal      int    `json:"decimal"`
	XferDisabled bool   `json:"xfer_disabled"` // token transfer disabled
}

type PeggedPairConfig struct {
	OrgChainId                uint      `json:"org_chain_id"`
	OrgToken                  TokenInfo `json:"org_token"`
	PeggedChainId             uint      `json:"pegged_chain_id"`
	PeggedToken               TokenInfo `json:"pegged_token"`
	PeggedDepositContractAddr string    `json:"pegged_deposit_contract_addr"`
	PeggedBurnContractAddr    string    `json:"pegged_burn_contract_addr"`
	VaultVersion              int       `json:"vault_version"`
	BridgeVersion             int       `json:"bridge_version"`
}

type EstimateAmountReq struct {
	SrcChainId        int
	DstChainId        int
	TokenSymbol       string
	UsrAddr           string
	SlippageTolerance int // slippage_tolerance_rate = slippage_tolerance / 1M, eg.  0.05% = 500
	Amt               string
	IsPegged          bool
}

type EstimateAmountRes struct {
	Err               Err     `json:"err"`
	EqValueTokenAmt   string  `json:"eq_value_token_amt"`
	BridgeRate        float64 `json:"bridge_rate"`
	PercFee           string  `json:"perc_fee"`
	BaseFee           string  `json:"base_fee"`
	SlippageTolerance int     `json:"slippage_tolerance"`
	MaxSlippage       int     `json:"max_slippage"`
}

type Err struct {
	Code int    `json:"code"` // 0 is success
	Msg  string `json:"msg"`
}
