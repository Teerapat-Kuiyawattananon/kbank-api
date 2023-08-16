package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/billdetail"
	"log"
	"time"

	mBillDetail "kapi/model"
)

type billDetailRepository struct {
	clientDB *ent.Client
}

func NewBillDetailRepository(DB *ent.Client) billDetailRepository {
	return billDetailRepository{
		clientDB: DB,
	}
}

func (repo billDetailRepository) CreateBillDetail(input mBillDetail.BillDetailRequest) (*ent.BillDetail, error) {
	billDetail, err := repo.clientDB.BillDetail.Create().
		SetCustomerID(input.CustomerId).
		SetTranAmount(input.TranAmount).
		SetChannelCode(input.ChannelCode).
		SetSenderBankCode(input.SenderBankCode).
		SetCreatedAt(time.Now().UTC().Add(time.Hour * 7)).
		SetUpdatedAt(time.Now().UTC().Add(time.Hour * 7)).
		SetStatus(input.Status).
		Save(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("Created BillDetail id: %d success", billDetail.ID)
	return billDetail, nil
}

func (repo billDetailRepository) GetBillDetailByRef2(ref2_id int) *ent.BillDetail {
	// Check database
	if repo.clientDB.BillDetail.Query().
		Where(billdetail.ID(ref2_id)).CountX(context.Background()) == 0 {
		return nil
	}

	model, err := repo.clientDB.BillDetail.Query().
		Where(billdetail.ID(ref2_id)).
		Only(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return model
}
