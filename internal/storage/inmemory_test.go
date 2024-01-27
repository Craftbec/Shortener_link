package storage

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInMemory_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInMemory := NewMockStorage(ctrl)
	original := "https://example.com"
	short := "short"
	mockInMemory.EXPECT().Get(context.Background(), short).Return(original, nil)
	result, err := mockInMemory.Get(context.Background(), short)
	if err != nil {
		t.Error("Unexpected error\n")
	}
	if result != original {
		t.Error("Expected 'original'\n")
	}
}

func TestInMemory_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInMemory := NewMockStorage(ctrl)
	mockInMemory.EXPECT().Post(context.Background(), "original", "short").Return(nil)
	err := mockInMemory.Post(context.Background(), "original", "short")
	if err != nil {
		t.Error("Unexpected error\n")
	}
}

func TestCheckPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInMemory := NewMockStorage(ctrl)
	original := "https://example.com"
	short := "short"
	mockInMemory.EXPECT().CheckPost(context.Background(), original).Return(short, nil)
	resultURL, err := mockInMemory.CheckPost(context.Background(), original)
	if resultURL != short {
		t.Error("Expected 'short'\n")
	}
	if err != nil {
		t.Errorf("Unexpected error\n")
	}
}
