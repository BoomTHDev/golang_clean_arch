package middleware

import (
	"errors"
	"log"

	"github.com/BoomTHDev/golang_clean_arch/pkg/custom"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		if err == nil {
			return ctx.Next()
		}

		appErr := &custom.AppError{}
		fiberErr := &fiber.Error{}

		if errors.As(err, &appErr) {
			if appErr.StatusCode >= fiber.StatusInternalServerError && appErr.Err != nil {
				log.Printf("Internal AppError: %s, Original Err: %s\n", appErr.Message, appErr.Err.Error())
			} else if appErr.Err != nil {
				log.Printf("AppError: %s, Original Err: %s\n", appErr.Message, appErr.Err.Error())
			} else {
				log.Printf("AppError: %s\n", appErr.Message)
			}

			return ctx.Status(appErr.StatusCode).JSON(fiber.Map{
				"success": false,
				"message": appErr.Message,
			})
		}

		if errors.As(err, &fiberErr) {
			log.Printf("Fiber Error: Code=%d, Message=%s\n", fiberErr.Code, fiberErr.Message)
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"success": false,
				"message": fiberErr.Message,
			})
		}

		log.Printf("Unhandled Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "An unexpected internal server error occurred.",
		})
	}
}
