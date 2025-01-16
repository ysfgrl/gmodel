
# go_blue_model RestApi standardizes requests and responses
[![Go Report Card](https://goreportcard.com/badge/github.com/ysfgrl/go_blue_model)](https://goreportcard.com/report/github.com/ysfgrl/go_blue_model)
[![GoDoc](https://godoc.org/github.com/ysfgrl/go_blue_model?status.svg)](https://godoc.org/github.com/ysfgrl/go_blue_model)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ysfgrl/go_blue_model/blob/master/LICENSE)



## ⚙️ Installation

```
go get -u github.com/ysfgrl/go_blue_model
```


## ⚡️ [Examples](https://github.com/ysfgrl/go_blue_model/tree/master/examples)

```go

func fiberOkResponse(ctx *fiber.Ctx) error {
    return  gmodel.Fiber.Ok(ctx, ResponseContent{
        Value: "value",
        Key: "key",
    })
}


func fiberErrorResponse(ctx *fiber.Ctx) error {
    return  gmodel.Fiber.NotFound(ctx, &gerror.Error{
        Code: "not.found",
        Detail: "Not Found",
    })
}


func fiberResponse(ctx *fiber.Ctx) error {
    return ctx.Status(fiber.StatusNotFound).JSON(gmodel.Response{
        Code: 404,
        Content: nil,
        Error: &gerror.Error{
            Code: "not.found",
            Detail: "Not Found",
        },
    })
}


func fiberResponse2(ctx *fiber.Ctx) error {
    return ctx.Status(fiber.StatusOK).JSON(gmodel.Response{
        Code: 200,
        Content: ResponseContent{
            Value: "value",
            Key: "key",
        },
    })
}
```
---

