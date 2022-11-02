package repositories

import (
	"errors"
	mongo_db "github.com/sanzharanarbay/golang_mongoDB_example/internal/configs/mongo-db"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type EmployeeRepository struct {
	MongoDB *mongo_db.MongoDB
}

func NewEmployeeRepository(MongoDB *mongo_db.MongoDB) *EmployeeRepository {
	return &EmployeeRepository{
		MongoDB: MongoDB,
	}
}

type EmployeeRepositoryInterface interface {
	SaveEmployee(employee *models.Employee) (interface{}, error)
	GetEmployee(ID string) (*models.Employee, error)
	GetAllEmployees() (*[]models.Employee, error)
	UpdateEmployee(employee *models.Employee, ID string) (interface{}, error)
	DeleteEmployee(ID string) error
	CheckEmployeeExistence(TabNum int64) (*models.Employee, error)
}

func (e *EmployeeRepository) SaveEmployee(employee *models.Employee) (interface{}, error) {
	employee.CreatedAt = time.Now()
	result, err := e.MongoDB.EmployeesCollection.InsertOne(e.MongoDB.Context, employee)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (e *EmployeeRepository) GetEmployee(ID string) (*models.Employee, error) {
	var employee models.Employee
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = e.MongoDB.EmployeesCollection.FindOne(e.MongoDB.Context, bson.D{{"_id", objectId}}).Decode(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (e *EmployeeRepository) GetAllEmployees() (*[]models.Employee, error) {
	var employee models.Employee
	var employees []models.Employee

	cursor, err := e.MongoDB.EmployeesCollection.Find(e.MongoDB.Context, bson.D{})
	if err != nil {
		defer cursor.Close(e.MongoDB.Context)
		return nil, err
	}

	for cursor.Next(e.MongoDB.Context) {
		err := cursor.Decode(&employee)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return &employees, nil
}

func (e *EmployeeRepository) UpdateEmployee(employee *models.Employee, ID string) (interface{}, error) {
	employee.UpdatedAt = time.Now()
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"tab_num", employee.TabNum}, {"first_name", employee.FirstName},
		{"last_name", employee.LastName}, {"age", employee.Age},
		{"department", employee.Department}, {"address", employee.Address},
		{"courses", employee.Courses}, {"is_active", employee.IsActive},
		{"is_gpx", employee.IsGpx}, {"updated_at", employee.UpdatedAt},
	}}}

	res := e.MongoDB.EmployeesCollection.FindOneAndUpdate(
		e.MongoDB.Context,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	var employeeUpdated models.Employee

	if err := res.Decode(&employeeUpdated); err != nil {
		return nil, err
	}

	return &employeeUpdated, nil
}

func (e *EmployeeRepository) DeleteEmployee(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	res, err := e.MongoDB.EmployeesCollection.DeleteOne(e.MongoDB.Context, bson.D{{"_id", objectId}})

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (e *EmployeeRepository) CheckEmployeeExistence(TabNum int64) (*models.Employee, error){
	var employee models.Employee
	filter := bson.D{{"tab_num", TabNum}}
	err := e.MongoDB.EmployeesCollection.FindOne(e.MongoDB.Context, filter).Decode(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}