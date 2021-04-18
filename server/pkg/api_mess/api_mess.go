package api_mess

import "go.mongodb.org/mongo-driver/bson"

func SuccessMessage(data interface{}) bson.M {
	return bson.M{
		"ok":   true,
		"data": data,
	}
}

func ErrorMessage(error interface{}) bson.M {
	return bson.M{
		"ok":   false,
		"data": error,
	}
}
