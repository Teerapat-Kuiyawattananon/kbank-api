package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/bill"
	"log"

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

func (repo billRepository) CreateBill(input mBill.Bill) (error) {
	bill, err := repo.clientDB.Bill.Create().
					SetBillerID(input.BillerId).
					SetRef1(input.Ref1).
					SetRef2(input.Ref2).
					Save(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Created Bill id: %d success", bill.ID)
	return nil
}

func (repo billRepository) GetBillByRef1Ref2(ref1_id int, ref2_id int) (mBill.Bill) {
	if (repo.clientDB.Bill.Query().Where(bill.Ref1(ref1_id), bill.Ref2(ref2_id)).
			CountX(context.Background()) == 0) {
				return mBill.Bill{}
		}
	model := repo.clientDB.Bill.Query().
				Where(bill.Ref1(ref1_id), bill.Ref2(ref2_id)).
				OnlyX(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	bill := mBill.Bill{
		ID: model.ID,
		BillerId: model.BillerID,
		Ref1: model.Ref1,
		Ref2: model.Ref2,
	}

	return bill
}