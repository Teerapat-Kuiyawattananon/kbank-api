package handler

import (
	model "kapi/model/inquiry"
	mBill "kapi/model/bill"
	db "kapi/progresql"
	repo "kapi/repository"
	"strconv"

	// db "kapi/progresql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	inquiryRequest model.InquiryRequest
	inquiryResponse model.InquiryResponse

)

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
		Reference1: inquiryRequest.Reference1,
		Reference2: inquiryRequest.Reference2,
		AdditionalFieldResponse: model.AdditionalFieldResponse{
			DueDate: "",
		},
	}
	check := checkRef1Ref2(inquiryRequest)
	if (!check) {
		inquiryResponse.ResponseCode = "9001"
		inquiryResponse.ResponseDescription = "Unauthorized"
		return c.JSON(http.StatusOK, inquiryResponse)
	}

	resCode, resDes := checkStatus(inquiryRequest)
	inquiryResponse.ResponseCode = resCode
	inquiryResponse.ResponseDescription = resDes

	return c.JSON(http.StatusOK, inquiryResponse)
}

func checkRef1Ref2(input model.InquiryRequest) bool {
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	billRepo := repo.NewBillRepository(clientDB)

	// strint to int
	ref1_id, _ := strconv.Atoi(input.Reference1)
	ref2_id, _ := strconv.Atoi(input.Reference2)
	check := billRepo.GetBillByRef1Ref2(ref1_id, ref2_id)
	if (check == mBill.Bill{}) {
		log.Println("false")
		return false
	}

	return true
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

