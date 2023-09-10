package consumer

import (
	"reflect"
	"testing"
)

func Test_createFullBufferedChannel(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want chan bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFullBufferedChannel(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFullBufferedChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}
