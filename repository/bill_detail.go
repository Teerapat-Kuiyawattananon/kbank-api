package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/billdetail"
	"log"

	mBillDetail "kapi/model/bill_detail"
)


type billDetailRepository struct {
	clientDB *ent.Client
}

func NewBillDetailRepository(DB *ent.Client) billDetailRepository {
	return billDetailRepository{
		clientDB: DB,
	}
}

func (repo billDetailRepository) CreateBillDetail(input mBillDetail.BillDetail) (error) {
	// create bill detail with id from bill.ref_2
	billDetail, err := repo.clientDB.BillDetail.Create().
					SetCustomerID(input.Customer_id).
					SetTranAmount(input.TranAmount).
					SetChannelCode(input.ChannelCode).
					SetSenderBankCode(input.SenderBankCode).
					SetCreatedAt(input.CreatedAt).
					SetUpdatedAt(input.UpdatedAt).
					SetStatus(input.Status).
					Save(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Created BillDetail id: %d success", billDetail.ID)
	return nil
}

func (repo billDetailRepository) GetBillDetailByRef2(ref2_id int) *ent.BillDetail {
	model, err := repo.clientDB.BillDetail.Query().
					Where(billdetail.ID(ref2_id)).
					Only(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return model
}