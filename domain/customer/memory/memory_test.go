package memory

import (
	"errors"
	"go-ddd-architecture/aggregate"
	"go-ddd-architecture/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetId()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("142e5494-9f58-4e43-a657-4373637d738c"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expect error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
