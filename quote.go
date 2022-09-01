package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// quote struct types

type quote struct {
	ID          primitive.ObjectID
	Description string
	Writer      string
	Book        string
}
