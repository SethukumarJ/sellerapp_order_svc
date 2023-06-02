
package routes

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/SethukumarJ/sellerapp/cmd/docs"
	"github.com/SethukumarJ/sellerapp/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// @Summary Register new user
// @ID User Registration
// @Tags Authentication-Service
// @Produce json
// @param RegisterUser body RegisterRequestBody{} true "User registration"
// @Success 200 {object} pb.RegisterResponse{}
// @Failure 422 {object} pb.RegisterResponse{}
// @Failure 502 {object} pb.RegisterResponse{}
// @Router /auth/register [post]
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:     body.Email,
		Password:  body.Password,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		UserName:  body.UserName,
	})

    fmt.Println("res",body)

	if err != nil {
		fmt.Println("er////r", err)
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
