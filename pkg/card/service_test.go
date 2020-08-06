package card

import "testing"

func TestService_Sum(t *testing.T) {
	type fields struct {
		Cards []*Card
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "first", fields: fields{Cards: []*Card{}}, want: 0},
		{name: "second", fields: fields{Cards: []*Card{
			{Balance: 10},
		}}, want: 10},
		{name: "third", fields: fields{Cards: []*Card{
			{Balance: 10},
			{Balance: 100},
		}}, want: 110},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Cards: tt.fields.Cards,
			}
			if got := s.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
