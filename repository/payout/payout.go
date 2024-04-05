package payout

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pradiptarana/simple-payment/model"
)

var baseURL = "https://660ed898356b87a55c50498c.mockapi.io/api/v1/"

type PayoutResponse struct {
	Status      string `json:"status"`
	ReferenceNo string `json:"reference_no"`
}

// PayoutRepository is responsible for storing and retrieving Payout information
type PayoutsRepository struct {
	db     *sql.DB
	client *http.Client
}

// NewPayoutRepository creates a new PayoutRepository instance
func NewPayoutRepository(db *sql.DB, client *http.Client) *PayoutsRepository {
	return &PayoutsRepository{db, client}
}

func (tr *PayoutsRepository) Payout(req *model.PayoutRequest) (*model.BankPayoutResponse, error) {
	var err error
	var data model.BankPayoutResponse

	reqBody := PayoutResponse{
		ReferenceNo: "123456789",
		Status:      "queued",
	}
	payoutJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	var payload = bytes.NewBuffer(payoutJSON)

	request, err := http.NewRequest("POST", baseURL+"/payout", payload)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := tr.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (tr *PayoutsRepository) GetAccountValidation(accountReq *model.AccountValidationRequest) (*model.BankAccount, error) {
	var err error
	var data []*model.BankAccount

	// request, err := http.NewRequest("GET", baseURL+"account_validation?account="+accountReq.Account+"&bank="+accountReq.Bank, nil)
	request, err := http.NewRequest("GET", baseURL+"account_validation", nil)
	if err != nil {
		return nil, err
	}

	response, err := tr.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	fmt.Println(response)
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data[0], nil
}

func (tr *PayoutsRepository) CreateTransaction(data *model.Transaction) error {
	stmt, err := tr.db.Prepare("insert into transaction (amount, status, reference_no, account_no, account_name, bank, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(data.Amount, data.Status, data.ReferenceNo, data.AccountNo, data.AccountName, data.Bank, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (tr *PayoutsRepository) UpdateTransaction(data *model.Transaction) error {
	stmt, err := tr.db.Prepare("update transaction set status = $1, updated_at = $2 where reference_no = $3 AND amount = $4")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(data.Status, data.UpdatedAt, data.ReferenceNo, data.Amount)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
