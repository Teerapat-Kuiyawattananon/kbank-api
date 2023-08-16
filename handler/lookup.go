package handler

import (
	model "kapi/model"
	db "kapi/progresql"
	repo "kapi/repository"
	"strconv"

	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// type InquiryHandler struct {
// 	inquiryRequest model.InquiryRequest
// 	inquiryResponse model.InquiryResponse
// 	BillRepo repo.BillRepository
// }

// func NewInquiryHandler() *InquiryHandler {
// 	return &InquiryHandler{}
// }

var (
	inquiryRequest model.InquiryRequest
	inquiryResponse model.InquiryResponse

)

// @Summary 	Show inquiry response.
// @Description Show inquiry response from inquiry request.
// @Tags 		Inquiry
// @Accept 		*/*
// @Produce 	json
// @Param body 	body model.InquiryRequest true "JSON request body for inquiry request"
// @Success 	200 {array} model.InquiryResponse "Success"
// @Router 		/api/billpayment/lookup [post]
func HandlerLookup(c echo.Context) error {
	if err := c.Bind(&inquiryRequest) ; err != nil {
		log.Fatal(err)
		return err
	}
	inquiryResponse = model.InquiryResponse{
		FunctionName: "BillLookupResponse",
		TransitionId: 	inquiryRequest.TransitionId,
		TransactionDateTime: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		BillerTransactionId: inquiryRequest.TransitionId,
		ResponseCode: "0000",
		ResponseDescription: "Success",
		BillerType: inquiryRequest.BillerType,
		BillerId: inquiryRequest.BillerId,
		TerminalNo: inquiryRequest.TerminalNo,
		// PromptPayTransactionId: inquiryRequest.PromptPayReferenceNumber,
		Reference1: inquiryRequest.Reference1,
		Reference2: inquiryRequest.Reference2,
		AdditionalFieldResponse: model.InquiryAdditionalFieldResponse{
			DueDate: "",
		},
	}
	checkCode, checkDes := checkRef1Ref2(inquiryRequest)
	if (checkCode != "") {
		inquiryResponse.ResponseCode = checkCode
		inquiryResponse.ResponseDescription = checkDes
		return c.JSON(http.StatusOK, inquiryResponse)
	}
	resCode, resDes := checkTranAmount(inquiryRequest)
	if (resCode != "") {
		inquiryResponse.ResponseCode = resCode
		inquiryResponse.ResponseDescription = resDes
		return c.JSON(http.StatusOK, inquiryResponse)
	}
	resCode, resDes = checkStatus(inquiryRequest)
	inquiryResponse.ResponseCode = resCode
	inquiryResponse.ResponseDescription = resDes

	if (inquiryRequest.BillerType == "BILLERID") {
		inquiryResponse.TypeofReceiver = "C"
		inquiryResponse.PromptPayTransactionId = inquiryRequest.PromptPayReferenceNumber
	}

	return c.JSON(http.StatusOK, inquiryResponse)
}

func checkRef1Ref2(input model.InquiryRequest) (string, string) {
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	billRepo := repo.NewBillRepository(clientDB)

	// strint to int
	ref1_id, err := strconv.Atoi(input.Reference1)
	if err != nil {
		return "0001", "Invalid Payment reference number"
	}
	ref2_id, err := strconv.Atoi(input.Reference2)
	if err != nil {
		return "0001", "Invalid Payment reference number"
	}
	check := billRepo.GetBillByRef1Ref2(ref1_id, ref2_id)
	if (check == model.Bill{}) {
		log.Println("false")
		return "9001", "Unauthorized"
	}

	return "", ""
}

func checkStatus(input model.InquiryRequest) (string, string) {
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	billDetailRepo := repo.NewBillDetailRepository(clientDB)

	// strint to int
	ref2_id, _ := strconv.Atoi(input.Reference2)
	bill := billDetailRepo.GetBillDetailByRef2(ref2_id)

	if (bill.Status == "paid") {
		return "0002", "Already paid"
	}
	if (bill.Status == "unavailable") {
		return "1000", "Other Merchant Error"
	}
	return "0000", "Success"
}

func checkTranAmount(input model.InquiryRequest) (string, string){
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	billDetailRepo := repo.NewBillDetailRepository(clientDB)
	ref2_id, _ := strconv.Atoi(input.Reference2)
	bill := billDetailRepo.GetBillDetailByRef2(ref2_id)
	tranAmount, err := strconv.ParseFloat(input.TranAmount, 64)
	if err != nil {
		return "0004", "Invalid payment amount"
	}

	if (tranAmount != bill.TranAmount) {
		return "0004", "Invalid payment amount"
	}
	
	return "", ""
}

