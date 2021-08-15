package domain_test

import (
	"testing"

	"github.com/ahsanulks/waitress/domain"
	"github.com/stretchr/testify/assert"
)

func TestOrderState_MarshalJSON(t *testing.T) {
	want, _ := domain.Pending.MarshalJSON()
	got := string(want)
	assert.Equal(t, "\"pending\"", got)

	var testInvalid domain.OrderState = 10
	_, err := testInvalid.MarshalJSON()
	assert.Error(t, err)
}

func TestOrderState_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    domain.OrderState
		wantErr bool
	}{
		{
			name:    "success",
			args:    `"paid"`,
			want:    domain.Paid,
			wantErr: false,
		},
		{
			name:    "invalid json format",
			args:    "requested",
			wantErr: true,
		},
		{
			name:    "not found",
			args:    `"on_going"`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var orderState domain.OrderState
			if err := orderState.UnmarshalJSON([]byte(tt.args)); (err != nil) != tt.wantErr {
				t.Errorf("OrderState.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.want, orderState)
		})
	}
}
