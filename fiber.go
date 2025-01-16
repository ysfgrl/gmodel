package gmodel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gerror"
)

type FiberC fiber.Ctx

var Fiber *fiberResponse

func init() {
	Fiber = &fiberResponse{}
}

type fiberResponse struct {
}

func (f *fiberResponse) base(c *fiber.Ctx, status int, content any, err *gerror.Error) error {
	return c.Status(status).JSON(Response{
		Code:    status,
		Content: content,
		Error:   err,
	})
}

func (f *fiberResponse) Ok(c *fiber.Ctx, content any) error {
	return f.base(c, fiber.StatusOK, content, nil)
}

func (f *fiberResponse) Created(c *fiber.Ctx, content any) error {
	return f.base(c, fiber.StatusCreated, content, nil)
}

func (f *fiberResponse) Unauthorized(c *fiber.Ctx, err *gerror.Error) error {
	return f.base(c, fiber.StatusUnauthorized, nil, err)
}
func (f *fiberResponse) Forbidden(c *fiber.Ctx, err *gerror.Error) error {
	return f.base(c, fiber.StatusForbidden, nil, err)
}
func (f *fiberResponse) NotAllowed(c *fiber.Ctx, err *gerror.Error) error {
	return f.base(c, fiber.StatusMethodNotAllowed, nil, err)
}

func (f *fiberResponse) NotImplemented(c *fiber.Ctx) error {
	return f.base(c, fiber.StatusNotImplemented, nil, gerror.UserError("NotImplemented", "api.not_implemented"))
}

func (f *fiberResponse) NotFound(c *fiber.Ctx, err *gerror.Error) error {
	return f.base(c, fiber.StatusNotFound, nil, err)
}

func (f *fiberResponse) BadRequest(c *fiber.Ctx, err *gerror.Error) error {
	return f.base(c, fiber.StatusBadRequest, nil, err)
}

func (f *fiberResponse) InternalServerError(c *fiber.Ctx, err *gerror.Error) error {
	return f.base(c, fiber.StatusInternalServerError, nil, err)
}
