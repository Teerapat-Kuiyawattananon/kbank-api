package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/bill"
	"log"
	"time"

	model "kapi/model"
)

type billRepository struct {
	clientDB *ent.Client
}

func NewBillRepository(DB *ent.Client) billRepository {
	return billRepository{
		clientDB: DB,
	}
}

func (repo billRepository) CreateBill(input model.Bill) (*ent.Bill, error) {
	bill, err := repo.clientDB.Bill.Create().
		SetBillerID(input.BillerId).
		SetReference1(input.Reference1).
		SetReference2(input.Reference2).
		SetTranAmount(input.TranAmount).
		SetChannelCode(input.ChannelCode).
		SetSenderBankCode(input.SenderBankCode).
		SetStatus(input.Status).
		Save(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("Created Bill id: %d success", bill.ID)
	return bill, nil
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

func (repo billRepository) UpdateBill(input model.Bill, id int) (*ent.Bill, error) {
	var billRes *ent.Bill
	bill, err := repo.clientDB.Bill.Query().
					Where(bill.ID(id)).
					Only(context.Background())
	if err != nil {
		log.Println(err)
	}
	if (input.BillerId != 0) {
		err = bill.Update().SetBillerID(input.BillerId).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	} 
	if (input.Reference1 != 0) {
		err = bill.Update().SetReference1(input.Reference1).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	if (input.Reference2 != 0) {
		err = bill.Update().SetReference2(input.Reference2).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	if (input.TransactionID != "") {
		err = bill.Update().SetTransactionID(input.TransactionID).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	if (input.ChannelCode != "") {
		err = bill.Update().SetChannelCode(input.ChannelCode).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	if (input.SenderBankCode != "") {
		err = bill.Update().SetSenderBankCode(input.SenderBankCode).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	if (input.Status != "") {
		err = bill.Update().SetStatus(input.Status).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	if (input.TranAmount != 0) {
		err = bill.Update().SetTranAmount(input.TranAmount).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}

	billRes = bill.Update().SetUpdatedAt(func () time.Time {
		strTime := time.Now().Add(time.Hour * 7).Format(time.RFC3339)
		t, _ := time.Parse(time.RFC3339, strTime)
		return t
	}()).
	SaveX(context.Background())

	return billRes, nil
}

func (repo billRepository) GetBillByDate (date model.DateParams) ([]*ent.Bill, error) {

	bills, err := repo.clientDB.Bill.Query().
					Where(bill.CreatedAtGTE(date.StartDate), bill.CreatedAtLTE(date.EndDate)).
					All(context.Background())
	if err != nil {
		log.Println(err)
	}

	return bills, nil
}