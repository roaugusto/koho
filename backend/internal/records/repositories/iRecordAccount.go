package records

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (

	//IRecordAccount - Record Account interface
	IRecordAccount interface {

		//Interface to insert one register on the database
		InsertOne(ctx context.Context, document interface{},
			opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

		//Interface to find registers on the database
		Find(ctx context.Context, filter interface{},
			opts ...*options.FindOptions) (*mongo.Cursor, error)

		//Interface to delete registers on the database
		DeleteMany(ctx context.Context, filter interface{},
			opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	}
)

//RecordHandler a record handler
type RecordHandler struct {
	Repo IRecordAccount
}
