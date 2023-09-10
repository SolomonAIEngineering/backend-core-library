package mongo

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestObjectIDFromHexIDString(t *testing.T) {
	type args struct {
		str *string
	}
	tests := []struct {
		name    string
		args    args
		want    primitive.ObjectID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ObjectIDFromHexIDString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectIDFromHexIDString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObjectIDFromHexIDString() = %v, want %v", got, tt.want)
			}
		})
	}
}
