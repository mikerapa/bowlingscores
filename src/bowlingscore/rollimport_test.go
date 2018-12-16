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
		// TODO: Add more test cases.
		{name: "extra spaces", args: args{rolls: []string{"4 ", "X "}}, wantRollData: []string{"4", "X"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRollData, err := getCleanedRollsData(tt.args.rolls)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCleanedRollsData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRollData, tt.wantRollData) {
				t.Errorf("getCleanedRollsData() = %v, want %v", gotRollData, tt.wantRollData)
			}
		})
	}
}
