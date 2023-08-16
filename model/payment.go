package model

type PaymentRequest struct {
	FunctionName 				string	`json:"functionName"`
	TransactionId 				string	`json:"transactionId"`
	TransactionDateTime			string	`json:"transactionDateTime"`
	BillerType 					string	`json:"billerType"`
	BillerId 					string	`json:"billerId"`
	TerminalNo 					string	`json:"terminalNo"`
	PromptPayReferenceNumber 	string	`json:"promptPayReferenceNumber"`
	ChannelCode					string	`json:"channelCode"`
	TranAmount					string	`json:"tranAmount"`
	SenderBankCode				string	`json:"senderBankCode"`
	IsRetry						string	`json:"isRetry"`
	Reference1					string	`json:"reference1"`
	Reference2					string	`json:"reference2"`
	Reference3					string	`json:"reference3"`
	ApiKey						string	`json:"apiKey"`
	AdditionalFieldRequest		PaymentAdditionalFieldRequest	`json:"additional"`
}

type PaymentAdditionalFieldRequest struct {
	CustomerFee			string	`json:"customerFee"`
	PartnerFee			string	`json:"partnerFee"`
	SponsorBankFee		string	`json:"sponsorBankFee"`
	TransgerWithTax		string	`json:"transferWithTax"`
	FromProxyValue		string	`json:"fromProxyValue"`
	FromProxyType		string	`json:"fromProxyType"`
	SenderReferenceNo	string	`json:"senderReferenceNo"`
	SenderTaxID			string	`json:"senderTaxID"`
	SenderAcctName		string	`json:"senderAcctName"`
	ReceiverTaxID		string	`json:"receiverTaxID"`
	VatRates			string	`json:"vatRates"`
	Vat 				string	`json:"vat"`
	TypeofTaxIncome		string	`json:"typeofTaxIncome"`
	WithholdingTaxRates	string	`json:"withholdingTaxRates"`
	WithholdingTax 		string	`json:"withholdingTax"`
	WithholdingTaxConditions string `json:"withholdingTaxConditions"`
	SettlementDate		string	`json:"settlementDate"`
	PromptPayFlag		string	`json:"promptPayFlag"`
	TypeofSender		string	`json:"typeofSender"`
	TypeofReceriver		string	`json:"typeofReceiver"`
	DueDate				string	`json:"dueDate"`
	RtpReference		string	`json:"rtpReference"`
	RqAppId				string	`json:"rqAppId"`
}

type PaymentResponse struct {
	FunctionName 			string	`json:"functionName"`
	TransitionId 			string	`json:"transactionId"`
	ResponseDateTime		string	`json:"responseDateTime"`
	BillerTransactionId		string	`json:"billerTransactionId"`
	ResponseCode			string	`json:"responseCode"`
	ResponseDescription		string	`json:"responseDescription"`
	TerminalNo 				string	`json:"terminalNo"`
	PromptPayTransactionId	string	`json:"promptPayTransactionId"`
	AdditionalFieldResponse PaymentAdditionalFieldResponse `json:"additional"`
}

type PaymentAdditionalFieldResponse struct {
	SettlementDate	string `json:"settlementDate"`
	RsAppId			string `json:"rsAppId"`
}