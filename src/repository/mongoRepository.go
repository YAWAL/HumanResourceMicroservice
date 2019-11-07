package repository

import (
	"context"
	"time"

	"github.com/YAWAL/ERP-common-lib/models"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	employeesDB         = "hr"
	employeesCollection = "employees"
	idField             = "_id"
	positionField       = "employee_info.position"
)

// MongoRepository structure with client to MongoBD
type MongoRepository struct {
	client *mongo.Client
}

// NewMongoRepository creates MongoDB repository with provided client
func NewMongoRepository(cl *mongo.Client) MongoRepository {
	return MongoRepository{client: cl}
}

func (mg MongoRepository) prepare() (*mongo.Collection, context.Context, context.CancelFunc) {
	collection := mg.client.Database(employeesDB).Collection(employeesCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return collection, ctx, cancel
}

// CreateEmployee store employee documents in Mongo database "hr", collection "employees"
func (mg MongoRepository) CreateEmployee(emp *models.Employee) error {
	collection, ctx, cancel := mg.prepare()
	defer cancel()

	emp.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, &emp)
	if err != nil {
		return err
	}
	logging.Log.Infof("employee with id %s has been created", res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

// UpdateEmployee searchs employee document by it's _id in Mongo database "hr", collection "employees"
// and updates all fields, provided in request body
func (mg MongoRepository) UpdateEmployee(emp *models.Employee) error {
	collection, ctx, cancel := mg.prepare()
	defer cancel()

	_, err := collection.ReplaceOne(ctx, bson.M{idField: &emp.ID}, &emp)
	if err != nil {
		return err
	}
	logging.Log.Infof("employee with id %s has been updated", emp.ID.Hex())
	return nil
}

// DeleteEmployee deletes employee document by _id(ObjectID) from Mongo database "hr", collection "employees"
func (mg MongoRepository) DeleteEmployee(id string) (bool, error) {
	collection, ctx, cancel := mg.prepare()
	defer cancel()

	odjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	deletedResult, err := collection.DeleteOne(ctx, bson.M{idField: odjectID})
	if err != nil {
		return false, err
	}
	if deletedResult.DeletedCount == 0 {
		logging.Log.Infof("DeleteEmployee: employee with id %s does not exists", id)
		return false, nil
	}
	logging.Log.Infof("%d employee with id %s has been deleted", deletedResult.DeletedCount, id)
	return true, nil
}

func (mg MongoRepository) find(filter bson.M) (emps []*models.Employee, err error) {
	collection, ctx, cancel := mg.prepare()
	defer cancel()

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem models.Employee
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		emps = append(emps, &elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return emps, nil

}

// GetEmployees retrieves all employees stored in Mongo database "hr", collection "employees"
func (mg MongoRepository) GetEmployees() (emps []*models.Employee, err error) {
	return mg.find(bson.M{})
}

// GetEmployeesByPosition retrieves all employees by provided position
// stored in Mongo database "hr", collection "employees"
func (mg MongoRepository) GetEmployeesByPosition(position string) (emps []*models.Employee, err error) {
	return mg.find(bson.M{positionField: position})
}
