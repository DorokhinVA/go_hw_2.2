package card

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "first", args: args{[]Card{}}, want: 0},
		{name: "second", args: args{cards: []Card{
			{Balance: 10},
		}}, want: 10},
		{name: "third", args: args{cards: []Card{
			{Balance: 10},
			{Balance: 100},
		}}, want: 110},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.cards); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
