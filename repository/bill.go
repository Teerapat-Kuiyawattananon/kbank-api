package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/bill"
	"log"
	"time"

	mBill "kapi/model"
)

type billRepository struct {
	clientDB *ent.Client
}

func NewBillRepository(DB *ent.Client) billRepository {
	return billRepository{
		clientDB: DB,
	}
}

func (repo billRepository) CreateBill(input mBill.Bill) error {
	bill, err := repo.clientDB.Bill.Create().
		SetBillerID(input.BillerId).
		SetReference1(input.Reference1).
		SetReference2(input.Reference2).
		SetTranAmount(input.TranAmount).
		SetChannelCode(input.ChannelCode).
		SetSenderBankCode(input.SenderBankCode).
		SetCreatedAt(func () time.Time {
			strTime := time.Now().Add(time.Hour * 7).Format(time.RFC3339)
			t, _ := time.Parse(time.RFC3339, strTime)
			return t
		}()).
		SetUpdatedAt(func () time.Time {
			strTime := time.Now().Add(time.Hour * 7).Format(time.RFC3339)
			t, _ := time.Parse(time.RFC3339, strTime)
			return t
		}()).
		SetStatus(input.Status).
		Save(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Created Bill id: %d success", bill.ID)
	return nil
}

func (repo billRepository) GetBillByRef1Ref2(ref1_id int, ref2_id int) *ent.Bill {
	if repo.clientDB.Bill.Query().Where(bill.Reference1(ref1_id), bill.Reference2(ref2_id)).
		CountX(context.Background()) == 0 {
		return nil
	}
	bill, err := repo.clientDB.Bill.Query().
		Where(bill.Reference1(ref1_id), bill.Reference2(ref2_id)).
		Only(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return bill
}

func (repo billRepository) GetAllBills() ([]*ent.Bill, error) {
	bills, err := repo.clientDB.Bill.Query().
					All(context.Background())
	if err != nil {
		log.Println(err)
	}

	return bills, nil
}