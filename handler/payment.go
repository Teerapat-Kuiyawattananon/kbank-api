package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	model "kapi/model"
	db "kapi/progresql"
	repo "kapi/repository"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	paymentRequest model.PaymentRequest
	paymentResponse model.PaymentResponse
)
func HandlerPayment(c echo.Context) error {
	if err := c.Bind(&paymentRequest) ; err != nil {
		log.Fatal(err)
		return err
	}

	paymentResponse = model.PaymentResponse{
		FunctionName: "BillPaymentResponse",
		TransitionId: paymentRequest.TransactionId,
		ResponseDateTime: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		BillerTransactionId: paymentRequest.TransactionId,
		ResponseCode: "0000",
		ResponseDescription: "Success",
		TerminalNo: paymentRequest.TerminalNo,
	}
	resCode, resDes := paymentConfirm(paymentRequest)
	if (paymentResponse.ResponseCode != resCode) {
		paymentResponse.ResponseCode = resCode
		paymentResponse.ResponseDescription = resDes
	}
	
	return c.JSON(http.StatusOK, paymentResponse)
}

func paymentConfirm(input model.PaymentRequest) (string, string) {
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	billRepo := repo.NewBillRepository(clientDB)

	// string to int
	ref1_id, err := strconv.Atoi(input.Reference1)
	if err != nil {
		return "0001", "Invalid Payment reference number"
	}
	ref2_id, err := strconv.Atoi(input.Reference2)
	if err != nil {
		return "0001", "Invalid Payment reference number"
	}
	tranAmount, err := strconv.ParseFloat(input.TranAmount, 64)
	if err != nil {
		return "0004", "Invalid payment amount"
	}
	bill := billRepo.GetBillByRef1Ref2(ref1_id, ref2_id)
	// if (bill.Status == "waiting") {
	// 	bill.Update().SetStatus("paid").ExecX(context.Background())
	// 	return "0000", "Success"
	// }
	defer func () {
		bill.Update().SetTransactionID(paymentRequest.TransactionId).
				SetUpdatedAt(func () time.Time {
					strTime := time.Now().Add(time.Hour * 7).Format(time.RFC3339)
					t, _ := time.Parse(time.RFC3339, strTime)
					return t
				}()).
				ExecX(context.Background())
	}()

	if (bill == nil) {
		return "9001", "Unauthorized"
	}
	if (tranAmount != bill.TranAmount) {
		return "0004", "Invalid payment amount"
	}
	if (bill.Status == "unavailable") {
		return "1000", "Other Merchant Error"
	}

	if (bill.Status == "already_paid") {
		return "0002", "Already paid"
	}
	// defer func () {
	// 	log.Println("defer")
	// 	bill.Update().SetTransactionID("paymentRequest.TransactionId").ExecX(context.Background())
	// }()
	
	// Waiting for payment
	bill.Update().SetStatus("already_paid").
					SetChannelCode(input.ChannelCode).
					SetSenderBankCode(input.SenderBankCode).
					SetUpdatedAt(func () time.Time {
						strTime := time.Now().Add(time.Hour * 7).Format(time.RFC3339)
						t, _ := time.Parse(time.RFC3339, strTime)
						return t
					}()).
					ExecX(context.Background())
	return "0000", "Success"
}
