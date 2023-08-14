package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	model "kapi/model/payment"
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

	billDetailRepo := repo.NewBillDetailRepository(clientDB)

	// string to int
	ref2_id, _ := strconv.Atoi(input.Reference2)
	bill := billDetailRepo.GetBillDetailByRef2(ref2_id)
	// if (bill.Status == "waiting") {
	// 	bill.Update().SetStatus("paid").ExecX(context.Background())
	// 	return "0000", "Success"
	// }

	if (bill.Status == "unavailable") {
		return "1000", "Other Merchant Error"
	}

	if (bill.Status == "paid") {
		return "0002", "Already paid"
	}

	// Waiting for payment
	bill.Update().SetStatus("paid").ExecX(context.Background())
	return "0000", "Success"
}