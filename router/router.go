package router

import (
	"github.com/gofiber/fiber"
	"github.com/mohdaalam005/one-to-many/controller"
)

func Route(app *fiber.App) {
	app.Post("/students", controller.CreateStudent)
	app.Get("/students", controller.GetStudents)
	app.Get("/students/:id", controller.GetStudent)
	app.Put("/students/:id", controller.UpdateStudent)
	app.Put("/students/:studentId/department/:departmentId", controller.UpdateDepartment)
	app.Get("/students/:studentId/department/:departmentId", controller.GetStudentAndDepartment)
	app.Delete("/students/:id", controller.DeleteStudent)
	app.Listen(":8080")
}
