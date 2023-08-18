package model

type InquiryRequest struct {
	FunctionName 				string	`json:"functionName"`
	TransitionId 				string	`json:"transactionId"`
	TransactionDateTime			string	`json:"transactionDateTime"`
	BillerType 					string	`json:"billerType"`
	BillerId 					string	`json:"billerId"`
	TerminalNo 					string	`json:"terminalNo"`
	PromptPayReferenceNumber 	string	`json:"promptPayReferenceNumber"`
	ChannelCode					string	`json:"channelCode"`
	TranAmount					string	`json:"tranAmount"`
	SenderBankCode				string	`json:"senderBankCode"`
	Reference1					string	`json:"reference1"`
	Reference2					string	`json:"reference2"`
	Reference3					string	`json:"reference3"`
	Language					string	`json:"language"`
	ApiKey						string	`json:"apiKey"`
	AdditionalFieldRequest		InquiryAdditionalFieldRequest	`json:"additional"`
}

type InquiryAdditionalFieldRequest struct {
	PayerFee			string	`json:"payerFee"`
	BillerFee			string	`json:"billerFee"`
	TransgerWithTax		string	`json:"transferWithTax"`
	FromProxyValue		string	`json:"fromProxyValue"`
	FromProxyType		string	`json:"fromProxyType"`
	SenderTaxID			string	`json:"senderTaxID"`
	SenderAcctName		string	`json:"senderAcctName"`
	VatRates			string	`json:"vatRates"`
	Vat 				string	`json:"vat"`
	TypeofTaxIncome		string	`json:"typeofTaxIncome"`
	WithholdingTaxRates	string	`json:"withholdingTaxRates"`
	WithholdingTax 		string	`json:"withholdingTax"`
	WithholdingTaxConditions string `json:"withholdingTaxConditions"`
	SettlementDate		string	`json:"settlementDate"`
	PromptPayFlag		string	`json:"promptPayFlag"`
	TypeofSender		string	`json:"typeofSender"`
	DueDate				string	`json:"dueDate"`
	RtpReference		string	`json:"rtpReference"`
	RqAppId				string	`json:"rqAppId"`
}	

type InquiryResponse struct {
	FunctionName 				string	`json:"functionName"`
	TransitionId 				string	`json:"transactionId"`
	TransactionDateTime			string	`json:"transactionDateTime"`
	BillerTransactionId			string  `json:"billerTransactionId"`
	ResponseCode				string	`json:"responseCode"`
	ResponseDescription			string	`json:"responseDescription"`
	BillerType 					string	`json:"billerType"`
	BillerId 					string	`json:"billerId"`
	TerminalNo 					string	`json:"terminalNo"`
	PromptPayTransactionId		string	`json:"promptPayTransactionId"`
	TypeofReceiver				string	`json:"typeofReceiver"`
	Reference1					string	`json:"reference1"`
	Reference2					string	`json:"reference2"`
	Reference3					string	`json:"reference3"`
	TranAmount					string	`json:"tranAmount"`
	Info1						string	`json:"info1"`
	Info2						string	`json:"info2"`
	Info3						string	`json:"info3"`
	AdditionalFieldResponse		InquiryAdditionalFieldResponse `json:"additional"`
}

type InquiryAdditionalFieldResponse struct {
	ToBillerAccountName	string	`json:"toBillerAccountName"`
	ToBillerServiceName	string	`json:"toBillerServiceName"`
	PayerFee			string	`json:"payerFee"`
	SettlementDate		string	`json:"settlementDate"`
	ReceiverTaxID		string	`json:"receiverTaxID"`
	DueDate				string	`json:"dueDate"`
	RtpReference		string	`json:"rtpReference"`
	RsAppId				string	`json:"rsAppId"`
}