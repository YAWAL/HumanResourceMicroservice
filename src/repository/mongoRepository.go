package repository

import (
	"context"
	"time"

	"github.com/YAWAL/ERP-common-lib/models"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	employeesDB         = "hr"
	employeesCollection = "employees"
)

// MongoRepository structure with client to MongoBD
type MongoRepository struct {
	client *mongo.Client
}

// NewMongoRepository creates MongoDB repository with provided client
func NewMongoRepository(cl *mongo.Client) MongoRepository {
	return MongoRepository{client: cl}
}

// CreateEmployee store employee documents in Mongo database "hr", collection "employees"
func (mg MongoRepository) CreateEmployee(emp *models.Employee) error {

	collection := mg.client.Database(employeesDB).Collection(employeesCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, emp)
	if err != nil {
		return err
	}
	logging.Log.Infof("employee with id %s has been created", res.InsertedID)
	return nil
}

// UpdateEmployee
func (mg MongoRepository) UpdateEmployee(emp *models.Employee) error {
	return nil
}

// DeleteEmployee
func (mg MongoRepository) DeleteEmployee(id int64) (bool, error) {
	return false, nil
}

// GetEmployees returns all employees stored in Mongo database "hr", collection "employees"
func (mg MongoRepository) GetEmployees() (emps []*models.Employee, err error) {
	collection := mg.client.Database(employeesDB).Collection(employeesCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
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
