package sample

import (
	"github.com/gofiber/fiber/v2"
	"initial/domain/sample/samplemodel"
	"initial/domain/sample/sampleservice"
	"initial/exception"
	"initial/infrastructure/shared/response"
)

type HttpSampleHandler interface {
	CreateData(c *fiber.Ctx) error
	UpdateData(c *fiber.Ctx) error
	DeleteData(c *fiber.Ctx) error
	GetDataById(c *fiber.Ctx) error
	GetAllData(c *fiber.Ctx) error
}

type httpSampleHandler struct {
	service sampleservice.Service
}

func NewSampleHandler(sampleService sampleservice.Service) HttpSampleHandler {
	return &httpSampleHandler{
		service: sampleService,
	}
}

func (s *httpSampleHandler) CreateData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var request samplemodel.SampleDataCreateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	request = samplemodel.SampleDataCreateModel{}

	result, err := s.service.CreateData(ctx, request)
	exception.PanicLogging(err)
	return c.Status(fiber.StatusCreated).JSON(response.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (s *httpSampleHandler) UpdateData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var request samplemodel.SampleDataUpdateModel
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	request = samplemodel.SampleDataUpdateModel{}

	result, err := s.service.UpdateData(ctx, request, id)
	exception.PanicLogging(err)
	return c.Status(fiber.StatusOK).JSON(response.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (s *httpSampleHandler) DeleteData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	err := s.service.DeleteData(ctx, id)
	exception.PanicLogging(err)
	return c.Status(fiber.StatusOK).JSON(response.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

func (s *httpSampleHandler) GetDataById(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	result, err := s.service.GetDataById(ctx, id)
	exception.PanicLogging(err)
	return c.Status(fiber.StatusOK).JSON(response.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (s *httpSampleHandler) GetAllData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	result, err := s.service.GetAllData(ctx)
	exception.PanicLogging(err)
	return c.Status(fiber.StatusOK).JSON(response.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
