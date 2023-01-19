package controller

import (
	"github.com/gofiber/fiber"
	"github.com/mohdaalam005/one-to-many/database"
	"github.com/mohdaalam005/one-to-many/helpers"
	"github.com/mohdaalam005/one-to-many/models"
	"strconv"
)

func CreateStudent(c *fiber.Ctx) {
	var student models.Student
	c.BodyParser(&student)
	database.DB.Create(&student)
	c.JSON(student)
}

func GetStudents(c *fiber.Ctx) {
	var student []models.Student
	database.DB.Model(&models.Student{}).Preload("Department").Find(&student)
	c.JSON(student)
}

func GetStudent(c *fiber.Ctx) {
	var student models.Student
	database.DB.Model(&models.Student{}).Preload("Department").Find(&student, c.Params("id"))
	c.JSON(student)
}
func UpdateDepartment(c *fiber.Ctx) {
	var student models.Student
	database.DB.Model(&models.Student{}).Preload("Department").Find(&student, c.Params("studentId"))
	departments := student.Department
	for _, department := range departments {
		id, err := strconv.Atoi(c.Params("departmentId"))
		helpers.CheckNilErr(err)
		if int(department.ID) == id {
			database.DB.Find(&department, c.Params("departmentId"))
			c.BodyParser(&department)
			database.DB.Save(&department)
			c.JSON(department)
		}

	}
}
func UpdateStudent(c *fiber.Ctx) {
	var student models.Student
	var department []models.Department
	database.DB.Find(&student, c.Params("studentId"))
	database.DB.Find(&department)
	c.BodyParser(&student)
	student.Department = department
	database.DB.Save(&student)
	c.JSON(student)
}

func GetStudentAndDepartment(c *fiber.Ctx) {
	var student models.Student
	database.DB.Model(&models.Student{}).Preload("Department").Find(&student, c.Params("studentId"))
	departments := student.Department

	for _, department := range departments {
		id, err := strconv.Atoi(c.Params("departmentId"))
		helpers.CheckNilErr(err)
		if int(department.ID) == id {
			database.DB.Find(&department, c.Params("departmentId"))
			c.JSON(department)
		}

	}
}

func DeleteStudent(c *fiber.Ctx) {
	var student models.Student
	database.DB.Model(models.Student{}).Delete(&student, c.Params("id"))
	c.JSON("deleted successfully ")
}
