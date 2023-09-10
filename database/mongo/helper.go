package mongo // import "github.com/SimifiniiCTO/simfiny-core-lib/database/mongo"

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectIDFromHexIDString takes a pointer to a string as input and returns a primitive.ObjectID and an error. It
// converts the input string, which is expected to be a hexadecimal representation of an ObjectID, to a
// primitive.ObjectID. The function is a method of the MongoDatabase struct and can be used to convert
// a string representation of an ObjectID to a primitive.ObjectID that can be used in MongoDB queries.
func ObjectIDFromHexIDString(str *string) (primitive.ObjectID, error) {
	if str == nil {
		return primitive.ObjectID{}, fmt.Errorf("invalid input argument. str: %v", str)
	}
	return primitive.ObjectIDFromHex(*str)
}
