package usecase

import (
	"context"
	"errors"
	"testing"

	mock "github.com/SethukumarJ/sellerapp-order-svc/pkg/mock/repoMock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestUpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockrepo := mock.NewMockOrderRepository(ctrl)

	usecase := NewOrderUseCase(mockrepo)
	ctx := context.Background()
	stringid := uuid.New().String()
	testData := []struct {
		name       string
		orderid    string
		status     string
		beforeTest func(userRepo *mock.MockOrderRepository)
		expectErr  error
	}{
		{
			name:    "Test success response",
			orderid: stringid,
			status:  "pending",
			beforeTest: func(userRepo *mock.MockOrderRepository) {
				userRepo.EXPECT().UpdateOrder(ctx, stringid, "pending").Return(stringid, nil)
			},
			expectErr: nil,
		},
		{
			name:    "Test when user alredy exist response",
			orderid: stringid,
			status:  "pending",
			beforeTest: func(userRepo *mock.MockOrderRepository) {
				userRepo.EXPECT().UpdateOrder(ctx, stringid, "pending").Return(stringid, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(mockrepo)
			actualId, err := usecase.UpdateOrder(ctx, tt.orderid, tt.status)
			assert.Equal(t, tt.expectErr, err)
			if err == nil {
				assert.Equal(t, tt.orderid, actualId)
			}
		})
	}
}

