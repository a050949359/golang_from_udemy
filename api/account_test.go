package api_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/golang_from_udemy/db/mock"
	db "github.com/golang_from_udemy/db/sqlc"
	"github.com/golang_from_udemy/util"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
