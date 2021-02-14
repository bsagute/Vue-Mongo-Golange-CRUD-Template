package robber

import "testing"

func TestGetMaxRobAmount(t *testing.T) {
	type args struct {
		houseOfAmounts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				houseOfAmounts: []int{1, 2, 3, 1},
			},
			want: 4,
		}, {
			name: "test2",
			args: args{
				houseOfAmounts: []int{2, 3, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxRobAmount(tt.args.houseOfAmounts); got != tt.want {
				t.Errorf("GetMaxRobAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
