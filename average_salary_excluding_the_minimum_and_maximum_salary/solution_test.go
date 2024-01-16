package averagesalaryexcludingtheminimumandmaximumsalary

import "testing"

func Test_average(t *testing.T) {
	tests := []struct {
		name   string
		salary []int
		want   float64
	}{
		{
			name:   "example 1",
			salary: []int{4000, 3000, 1000, 2000},
			want:   2500.00000,
		},
		{
			name:   "example 2",
			salary: []int{1000, 2000, 3000},
			want:   2000.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
