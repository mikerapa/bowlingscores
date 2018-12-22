package bowlingscore

import (
	"reflect"
	"testing"
)

func Test_isValidRollString(t *testing.T) {
	type args struct {
		rollString string
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
	}{
		{"C", args{rollString: "C"}, false},
		{"3", args{rollString: "3"}, true},
		{"blank", args{rollString: ""}, false},
		{"X", args{rollString: "X"}, true},
		{"/ (Spare)", args{rollString: "/"}, true},
	}
	for _, tt := range tests {
		t.Run("value is "+tt.name, func(t *testing.T) {
			if gotValid := isValidRollString(tt.args.rollString); gotValid != tt.wantValid {
				t.Errorf("isValidRollString() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func Test_getCleanedRollsData(t *testing.T) {
	type args struct {
		rolls []string
	}
	tests := []struct {
		name         string
		args         args
		wantRollData []string
		wantErr      bool
	}{

		{name: "extra spaces", args: args{rolls: []string{"4 ", "X "}}, wantRollData: []string{"4", "X"}, wantErr: false},
		{name: "Invalid char", args: args{rolls: []string{"4 ", "X ", "5", "7", "1", "7", "1", "7", "&", "X"}}, wantRollData: []string{}, wantErr: true},
		{name: "double chars", args: args{rolls: []string{"4 ", "X ", "5", "71", "1", "7", "1", "7", "&", "XX"}}, wantRollData: []string{}, wantErr: true},
		{name: "lower case x", args: args{rolls: []string{"4 ", "X ", "5", "1", "7", "1", "7", "x"}}, wantRollData: []string{"4", "X", "5", "1", "7", "1", "7", "X"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRollData, err := getCleanedRollsData(tt.args.rolls)
			if (err != nil) != tt.wantErr {
				// wanted an error but didn't get one
				t.Errorf("getCleanedRollsData() name=%s, error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRollData, tt.wantRollData) && !tt.wantErr {
				t.Errorf("getCleanedRollsData() name=%s, = %v, want %v", tt.name, gotRollData, tt.wantRollData)
			}
		})
	}
}

func Test_nomalizeRollString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "leading spaces", input: "  5", want: "5"},
		{name: "trailing", input: "5  ", want: "5"},
		{name: "lower case x", input: "  x", want: "X"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nomalizeRollString(tt.input); got != tt.want {
				t.Errorf("nomalizeRollString() = %v, want %v", got, tt.want)
			}
		})
	}
}
