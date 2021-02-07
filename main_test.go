package main

import "testing"

type ChooseMock struct {
	Result int
}

func (rng *ChooseMock) TimeChooseIntnList(n int) int {
	result := rng.Result
	return result
}
func (rng *ChooseMock) CryptoChooseIntnList(n int) int {
	result := rng.Result
	return result
}

// Chooseをモックにすることで確実に望みの結果に対する挙動をテストできる
func TestAttack_DiceのMockで確実に成功(t *testing.T) {
	l := []string{"0", "1", "2", "3"}
	// 0 が出るように仕込む
	dice := &ChooseMock{Result: 0}
	a := &Chooser{c: dice}
	got := a.timeRandChooseList(l)
	want := "0"

	if got != want {
		t.Errorf("timeRandChooseList() = %v, want %v", got, want)
	}
}

func TestAttack_ケース網羅(t *testing.T) {
	type fakes struct {
		FakeResult int
	}
	tests := []struct {
		l          []string
		name       string
		fakes      fakes
		wantResult string
	}{
		{[]string{"0", "1", "2", "3"}, "0", fakes{FakeResult: 0}, "0"},
		{[]string{"0", "1", "2", "3"}, "1", fakes{FakeResult: 1}, "1"},
		{[]string{"0", "1", "2", "3"}, "2", fakes{FakeResult: 2}, "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dice := &ChooseMock{Result: tt.fakes.FakeResult}
			a := &Chooser{c: dice}
			if gotResult := a.timeRandChooseList(tt.l); gotResult != tt.wantResult {
				t.Errorf("timeRandChooseList() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
