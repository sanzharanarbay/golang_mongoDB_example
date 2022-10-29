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

type CourseRepository struct {
	MongoDB *mongo_db.MongoDB
}

func NewCourseRepository(MongoDB *mongo_db.MongoDB) *CourseRepository {
	return &CourseRepository{
		MongoDB: MongoDB,
	}
}

type CourseRepositoryInterface interface {
	SaveCourse(course *models.Course) (interface{}, error)
	GetCourse(ID string) (*models.Course, error)
	GetAllCourses() (*[]models.Course, error)
	UpdateCourse(course *models.Course, ID string) (interface{}, error)
	DeleteCourse(ID string) error
}

func (c *CourseRepository) SaveCourse(course *models.Course) (interface{}, error) {
	course.CreatedAt = time.Now()
	result, err := c.MongoDB.CoursesCollection.InsertOne(c.MongoDB.Context, course)
	if err != nil {
		panic(err)
	}
	return result.InsertedID, nil
}

func (c *CourseRepository) GetCourse(ID string) (*models.Course, error) {
	var course models.Course
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = c.MongoDB.CoursesCollection.FindOne(c.MongoDB.Context, bson.D{{"_id", objectId}}).Decode(&course)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (c *CourseRepository) GetAllCourses() (*[]models.Course, error) {
	var course models.Course
	var courses []models.Course

	cursor, err := c.MongoDB.CoursesCollection.Find(c.MongoDB.Context, bson.D{})
	if err != nil {
		defer cursor.Close(c.MongoDB.Context)
		return nil, err
	}

	for cursor.Next(c.MongoDB.Context) {
		err := cursor.Decode(&course)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return &courses, nil
}

func (c *CourseRepository) UpdateCourse(course *models.Course, ID string) (interface{}, error) {
	course.UpdatedAt = time.Now()
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"title", course.Title}, {"hours", course.Hours},
		{"author", course.Author}, {"price", course.Price},
		{"order", course.Order}, {"is_active", course.IsActive},
		{"updated_at", course.UpdatedAt},
	}}}

	res := c.MongoDB.CoursesCollection.FindOneAndUpdate(
		c.MongoDB.Context,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	var courseUpdated models.Course

	if err := res.Decode(&courseUpdated); err != nil {
		return nil, err
	}

	return &courseUpdated, nil
}

func (c *CourseRepository) DeleteCourse(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	res, err := c.MongoDB.CoursesCollection.DeleteOne(c.MongoDB.Context, bson.D{{"_id", objectId}})

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
