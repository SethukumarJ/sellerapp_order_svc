package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SethukumarJ/sellerapp/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @ID Delete user
// @Tags User-Section
// @Produce json
// @param id path string true "Id of the user to be deleted"
// @Success 200 {object} pb.DeleteUserResponse{}
// @Failure 422 {object} pb.DeleteUserResponse{}
// @Failure 502 {object} pb.DeleteUserResponse{}
// @Router /user/delete/{id} [delete]
// DeleteUser deletes a user with the given ID.
func DeleteUser(ctx *gin.Context, c pb.AuthServiceClient) {
	paramsID := ctx.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.DeleteUserResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}

	res, err := c.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: int64(id)})
	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
