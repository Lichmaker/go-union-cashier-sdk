package apirequest

import (
	"encoding/json"
	"time"
)

type BizContentBody struct {
	SndDt     string      `json:"sndDt"`
	BusiMerNo string      `json:"busiMerNo"`
	MsgBody   interface{} `json:"msgBody"`
	NotifyUrl string      `json:"notifyUrl,omitempty"`
}

type PrePayMsgBody struct {
	MerOrdrNo        string `json:"merOrdrNo"`
	TrxTtlAmt        string `json:"trxTtlAmt"`
	BizFunc          string `json:"bizFunc" default:"111011"`
	BizType          string `json:"bizType"`
	Subject          string `json:"subject"`
	OrdrDesc         string `json:"ordrDesc"`
	OprId            string `json:"oprId" default:"001"`
	TimeoutExpress   string `json:"timeoutExpress" default:"30m"`
	FrontRedirectURL string `json:"frontRedirectURL,omitempty"`
	TrxChnlType      string `json:"trxChnlType"`
	UserNo           string `json:"userNo"`
	UserType         string `json:"userType"`
	AssignPayType    string `json:"assignPayType,omitempty"`
	Remark           string `json:"remark,omitempty"`
	Remark1          string `json:"remark1,omitempty"`
	Remark2          string `json:"remark2,omitempty"`
}

type PayQueryMsgBody struct {
	MerOrdrNo    string `json:"merOrdrNo"`
	OriMerOrdrNo string `json:"oriMerOrdrNo"`
	OriTrxDt     string `json:"oriTrxDt"`
}

type OrderCancelMsgBody struct {
	MerOrdrNo    string `json:"merOrdrNo"`
	OriMerOrdrNo string `json:"oriMerOrdrNo"`
	OriTrxDt     string `json:"oriTrxDt"`
}

type RefundMsgBody struct {
	MerOrdrNo    string `json:"merOrdrNo"`
	OriMerOrdrNo string `json:"oriMerOrdrNo"`
	OriTrxDt     string `json:"oriTrxDt"`
	TrxAmt     string `json:"trxAmt"`
	BizFunc     string `json:"bizFunc" default:"411011"`
}

type RefundQueryMsgBody struct {
	MerOrdrNo    string `json:"merOrdrNo"`
	OriMerOrdrNo string `json:"oriMerOrdrNo"`
	OriTrxDt     string `json:"oriTrxDt"`
	BizFunc     string `json:"bizFunc" default:"611021"`
}

func PrePay(busiMerNo string, notifyUrl string, msgBody PrePayMsgBody) Request {
	bizContent := BizContentBody{
		SndDt:     time.Now().Format("20060102150405"),
		BusiMerNo: busiMerNo,
		NotifyUrl: notifyUrl,
		MsgBody:   msgBody,
	}
	bizContentJson, _ := json.Marshal(bizContent)

	req := Request{
		ApiInterfaceId: "gnete.upbc.cashier.trade",
		MethodName:     "prePay",
		BizContent:     string(bizContentJson),
	}
	return req
}

func PayQuery(busiMerNo string, msgBody PayQueryMsgBody) Request {
	bizContent := BizContentBody{
		SndDt:     time.Now().Format("20060102150405"),
		BusiMerNo: busiMerNo,
		MsgBody:   msgBody,
	}
	bizContentJson, _ := json.Marshal(bizContent)

	req := Request{
		ApiInterfaceId: "gnete.upbc.cashier.trade",
		MethodName:     "payQuery",
		BizContent:     string(bizContentJson),
	}
	return req
}

func OrderCancel(busiMerNo string, msgBody OrderCancelMsgBody) Request {
	bizContent := BizContentBody{
		SndDt:     time.Now().Format("20060102150405"),
		BusiMerNo: busiMerNo,
		MsgBody:   msgBody,
	}
	bizContentJson, _ := json.Marshal(bizContent)

	req := Request{
		ApiInterfaceId: "gnete.upbc.cashier.trade",
		MethodName:     "cancel",
		BizContent:     string(bizContentJson),
	}
	return req
}

func Refund(busiMerNo string, msgBody RefundMsgBody) Request {
	bizContent := BizContentBody{
		SndDt:     time.Now().Format("20060102150405"),
		BusiMerNo: busiMerNo,
		MsgBody:   msgBody,
	}
	bizContentJson, _ := json.Marshal(bizContent)

	req := Request{
		ApiInterfaceId: "gnete.upbc.cashier.trade",
		MethodName:     "refund",
		BizContent:     string(bizContentJson),
	}
	return req
}

func RefundQuery(busiMerNo string, msgBody RefundQueryMsgBody) Request {
	bizContent := BizContentBody{
		SndDt:     time.Now().Format("20060102150405"),
		BusiMerNo: busiMerNo,
		MsgBody:   msgBody,
	}
	bizContentJson, _ := json.Marshal(bizContent)

	req := Request{
		ApiInterfaceId: "gnete.upbc.cashier.trade",
		MethodName:     "refundQuery",
		BizContent:     string(bizContentJson),
	}
	return req
}
