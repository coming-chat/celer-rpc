package celerrpc

import (
	"testing"
)

func TestRestClient_GetTransferConfigs(t *testing.T) {
	type fields struct {
		rpcUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "case1",
			fields:  fields{rpcUrl: "https://cbridge-prod2.celer.app"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRestClient(tt.fields.rpcUrl)
			gotTransferConfig, err := c.GetTransferConfigs()
			if (err != nil) != tt.wantErr {
				t.Errorf("RestClient.GetTransferConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTransferConfig == nil {
				t.Errorf("RestClient.GetTransferConfigs() got nil")
			}
		})
	}
}

func TestRestClient_EstimateAmount(t *testing.T) {
	type fields struct {
		rpcUrl string
	}
	type args struct {
		estimateAmountReq EstimateAmountReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "case 1",
			fields: fields{rpcUrl: "https://cbridge-prod2.celer.app"},
			args: args{
				estimateAmountReq: EstimateAmountReq{
					SrcChainId:        1,
					DstChainId:        56,
					TokenSymbol:       "USDT",
					UsrAddr:           "0x0e9D66A7008ca39AE759569Ad1E911d29547E892",
					SlippageTolerance: 500,
					Amt:               "1000000", // 1usdc
					IsPegged:          false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRestClient(tt.fields.rpcUrl)
			gotRes, err := c.EstimateAmount(tt.args.estimateAmountReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("RestClient.EstimateAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes == nil {
				t.Errorf("RestClient.EstimateAmount() got nil")
			}
		})
	}
}
