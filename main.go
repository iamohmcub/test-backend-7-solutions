package main

import (
	"seven-solutions-backend-test/test1"
	"seven-solutions-backend-test/test2"
	"seven-solutions-backend-test/test3"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(`
			Path "/test1/:testCase" (1 for testcase 1 AND 2 for testcase 2)
			Path "/test2/:encodeInput" (encodeInput ex. LLRR=)
			Path "/beef/summary" (for test3 beef summary)
			Path "/run-test" (for get all test result)
		`)
	})

	app.Get("/test1/:testCase", Test1)
	app.Get("/test2/:input", Test2)
	app.Get("/beef/summary", Test3)

	app.Get("/run-test", RunTest)
	app.Use(func(c fiber.Ctx) error {
		return c.Status(404).SendString("NotFound")
	})
	app.Listen(":3000")
}

type SuccessResponse struct {
	Output string `json:"output"`
}

func Test1(c fiber.Ctx) error {
	testCase := c.Params("testCase")
	result := test1.GetGreatestPath(testCase)
	response := SuccessResponse{
		Output: result,
	}
	return c.Status(200).JSON(response)
}

func Test2(c fiber.Ctx) error {
	input := c.Params("input")
	result := test2.CatchMeIfYouCan(input)
	response := SuccessResponse{
		Output: result,
	}
	return c.Status(200).JSON(response)
}

func Test3(c fiber.Ctx) error {
	response := test3.GetResponseBeef()
	return c.JSON(response)
}

func RunTest(c fiber.Ctx) error {
	result := make(map[string]any)

	result["1.1"] = test1.GetGreatestPath("1") //Test Case 1.1
	result["1.2"] = test1.GetGreatestPath("2") //Test Case 1.2

	result["2.1"] = test2.CatchMeIfYouCan("LLRR=") //Test Case 2.1
	result["2.2"] = test2.CatchMeIfYouCan("==RLL") //Test Case 2.1
	result["2.3"] = test2.CatchMeIfYouCan("=LLRR") //Test Case 2.1
	result["2.4"] = test2.CatchMeIfYouCan("RRL=R") //Test Case 2.1

	result["3"] = test3.GetResponseBeef() //Test Case 3.1

	return c.JSON(result)
}
