package adapter

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func AdminAdapter(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println(time.Now())
	return next
}
