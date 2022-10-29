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

type DepartmentRepository struct {
	MongoDB *mongo_db.MongoDB
}

func NewDepartmentRepository(MongoDB *mongo_db.MongoDB) *DepartmentRepository {
	return &DepartmentRepository{
		MongoDB: MongoDB,
	}
}

type DepartmentRepositoryInterface interface {
	SaveDepartment(department *models.Department) (interface{}, error)
	GetDepartment(ID string) (*models.Department, error)
	GetAllDepartments() (*[]models.Department, error)
	UpdateDepartment(department *models.Department, ID string) (interface{}, error)
	DeleteDepartment(ID string) error
}

func (d *DepartmentRepository) SaveDepartment(department *models.Department) (interface{}, error) {
	department.CreatedAt = time.Now()
	result, err := d.MongoDB.DepartmentsCollection.InsertOne(d.MongoDB.Context, department)
	if err != nil {
		panic(err)
	}
	return result.InsertedID, nil
}

func (d *DepartmentRepository) GetDepartment(ID string) (*models.Department, error) {
	var department models.Department
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = d.MongoDB.DepartmentsCollection.FindOne(d.MongoDB.Context, bson.D{{"_id", objectId}}).Decode(&department)
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (d *DepartmentRepository) GetAllDepartments() (*[]models.Department, error) {
	var department models.Department
	var departments []models.Department

	cursor, err := d.MongoDB.DepartmentsCollection.Find(d.MongoDB.Context, bson.D{})
	if err != nil {
		defer cursor.Close(d.MongoDB.Context)
		return nil, err
	}

	for cursor.Next(d.MongoDB.Context) {
		err := cursor.Decode(&department)
		if err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	return &departments, nil
}

func (d *DepartmentRepository) UpdateDepartment(department *models.Department, ID string) (interface{}, error) {
	department.UpdatedAt = time.Now()
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"name", department.Name}, {"short_name", department.ShortName},
		{"total.employees", department.Total.Employees}, {"total.gpx", department.Total.Gpx},
		{"is_active", department.IsActive}, {"order", department.Order},
		{"updated_at", department.UpdatedAt},
	}}}

	res := d.MongoDB.DepartmentsCollection.FindOneAndUpdate(
		d.MongoDB.Context,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	var departmentUpdated models.Department

	if err := res.Decode(&departmentUpdated); err != nil {
		return nil, err
	}

	return &departmentUpdated, nil
}

func (d *DepartmentRepository) DeleteDepartment(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	res, err := d.MongoDB.DepartmentsCollection.DeleteOne(d.MongoDB.Context, bson.D{{"_id", objectId}})

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
