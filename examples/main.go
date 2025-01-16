package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gerror"
	"github.com/ysfgrl/gmodel"
)

type ResponseContent struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {

}

func fiberOkResponse(ctx *fiber.Ctx) error {
	return gmodel.Fiber.Ok(ctx, ResponseContent{
		Value: "value",
		Key:   "key",
	})
}

func fiberErrorResponse(ctx *fiber.Ctx) error {
	return gmodel.Fiber.NotFound(ctx, &gerror.Error{
		Code:   "not.found",
		Detail: "Not Found",
	})
}

func fiberResponse(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(gmodel.Response{
		Code:    404,
		Content: nil,
		Error: &gerror.Error{
			Code:   "not.found",
			Detail: "Not Found",
		},
	})
}

func fiberResponse2(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(gmodel.Response{
		Code: 200,
		Content: ResponseContent{
			Value: "value",
			Key:   "key",
		},
	})
}
