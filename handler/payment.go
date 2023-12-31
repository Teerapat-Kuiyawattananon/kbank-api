package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	model "kapi/model"
	repo "kapi/repository"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	paymentRequest model.PaymentRequest
	paymentResponse model.PaymentResponse
)

// @Summary 	Show payment response.
// @Description Show payment response after payment process.
// @Tags 		Payment
// @Accept 		*/*
// @Produce 	json
// @Param body 	body model.PaymentRequest true "JSON request body for payment request"
// @Success 	200 {array} model.PaymentResponse "Success"
// @Router 		/api/billpayment/payment [post]
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
	billRepo := repo.NewBillRepository(clientDB)

	// Convert string to int.
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
	defer func () {
		bill.Update().SetTransactionID(paymentRequest.TransactionId).
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

	// Waiting for payment.
	bill.Update().SetStatus("already_paid").
					SetChannelCode(input.ChannelCode).
					SetSenderBankCode(input.SenderBankCode).
					ExecX(context.Background())
	return "0000", "Success"
}
