package api

import (
	"bytes"
	// "database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/golang_from_udemy/db/mock"
	db "github.com/golang_from_udemy/db/sqlc"
	"github.com/golang_from_udemy/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserAPI(t *testing.T) {
	user := randomUser()

	testCase := []struct {
		name          string
		username      string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			username: user.Username,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(
					gomock.Any(),
					gomock.Any(),
				).Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		// , {
		// 	name:      "NotFound",
		// 	accountID: account.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().GetAccount(
		// 			gomock.Any(),
		// 			gomock.Eq(account.ID),
		// 		).Times(1).Return(db.Account{}, sql.ErrNoRows)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusNotFound, recorder.Code)
		// 		// requireBodyMatchAccount(t, recorder.Body, account)
		// 	},
		// }, {
		// 	name:      "InternalError",
		// 	accountID: account.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().GetAccount(
		// 			gomock.Any(),
		// 			gomock.Eq(account.ID),
		// 		).Times(1).Return(db.Account{}, sql.ErrConnDone)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusInternalServerError, recorder.Code)
		// 		// requireBodyMatchAccount(t, recorder.Body, account)
		// 	},
		// }, {
		// 	name:      "InvalidID",
		// 	accountID: 0,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().GetAccount(
		// 			gomock.Any(),
		// 			gomock.Any(),
		// 		).Times(0)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 		// requireBodyMatchAccount(t, recorder.Body, account)
		// 	},
		// },
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			userReq := CreateUserRequest{
				Username: user.Username,
				Password: user.HashedPassword,
				FullName: user.FullName,
				Email: user.Email,
			}

			userBytes, err := json.Marshal(userReq)
			require.NoError(t, err)

			bodyReader := bytes.NewReader(userBytes)

			url := "/user"
			request, err := http.NewRequest(http.MethodPost, url, bodyReader)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomUser() db.User {
	hashPassword, _ := util.NewPassword(util.RandomString(6))

	return db.User{
		Username:       util.RandomString(6),
		HashedPassword: hashPassword,
		FullName:  		util.RandomString(6),
		Email: 			util.RandomEmail(),
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getUser db.User
	err = json.Unmarshal(data, &getUser)
	require.NoError(t, err)
	require.Equal(t, user.Username, getUser.Username)
	require.Equal(t, user.FullName, getUser.FullName)
	require.Equal(t, user.Email, getUser.Email)
	require.Equal(t, user.CreatedAt, getUser.CreatedAt)
	require.Equal(t, user.PasswordChangedAt, getUser.PasswordChangedAt)
}
