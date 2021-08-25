package routes

import (
	"context"
	"log"

	"github.com/ferjoaguilar/rest-go.git/src/config"
	"github.com/ferjoaguilar/rest-go.git/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupStudentRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/get-students", getStudents)
	route.Post("/add-student", addStudent)
	route.Put("/edit-student/:id", editStudent)
	route.Delete("/delete-student/:id", deleteStudent)
}

func getStudents(c *fiber.Ctx) error {
	var studentsDB []models.Student 
	studentQuery := config.Query.DB.Collection("students")
	res, err := studentQuery.Find(context.TODO(), bson.D{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
			"error": err.Error(),
		})
	}
	for res.Next(context.TODO()) {
		var student models.Student
		err := res.Decode(&student)

		if err != nil {
			log.Fatal(err)
		}
		studentsDB = append(studentsDB, student)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": studentsDB,
		"status": "OK",
	})
}

func addStudent(c *fiber.Ctx) error{
	studentQuery := config.Query.DB.Collection("students")
	studentDB := new(models.Student)

	if err := c.BodyParser(studentDB); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parser body",
			"status": "Bad",
		})
	}

	res, err := studentQuery.InsertOne(context.TODO(), studentDB)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
			"status": "Bad",
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": res,
		"status": "OK",
		"message": "New student inserted successfully",
	})
}

func editStudent(c *fiber.Ctx) error{
	studentQuery := config.Query.DB.Collection("students")
	studentDB := new(models.Student)

	if err := c.BodyParser(studentDB); err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parser body",
			"status": "Bad",
		})
	}

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Student not found",
			"status": "Bad",
			"error": err.Error(),
		})
	}

	update := bson.M{
        "$set": studentDB,
    }
	_, err = studentQuery.UpdateOne(context.TODO(), bson.M{"_id": objId}, update)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Student failed to update",
			"status": "Bad",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Student updated successfully",
		"status": "OK",
	})
}

func deleteStudent(c *fiber.Ctx)error {
	studentQuery := config.Query.DB.Collection("students")

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Student no found",
			"status": "Bad",
		})
	}

	_, err = studentQuery.DeleteOne(context.TODO(), bson.M{"_id": objId})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Student failed to delete",
			"status": "Bad",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Student deleted successfully",
		"status": "OK",
	})
}