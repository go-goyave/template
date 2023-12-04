package user

import (
	"context"
	"net/http"
	"strconv"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/typeutil"
	"goyave.dev/template/dto"
	"goyave.dev/template/service"
)

type Service interface {
	First(ctx context.Context, id uint) (*dto.User, error)
	Paginate(ctx context.Context, page int, pageSize int) (*database.PaginatorDTO[*dto.User], error)
}

type UserController struct {
	goyave.Component
	UserService Service
}

func (ctrl *UserController) Init(server *goyave.Server) {
	ctrl.UserService = server.Service(service.User).(Service)
	ctrl.Component.Init(server)
}

func (ctrl *UserController) RegisterRoutes(router *goyave.Router) {
	subrouter := router.Subrouter("/users")
	subrouter.CORS(cors.Default())

	subrouter.Get("/", ctrl.Index).ValidateQuery(IndexRequest)
	subrouter.Get("/{userID:[0-9+]}", ctrl.Show)
}

func (ctrl *UserController) Index(response *goyave.Response, request *goyave.Request) {

	query := typeutil.MustConvert[dto.Index](request.Query)

	paginator, err := ctrl.UserService.Paginate(request.Context(), query.Page.Default(1), query.PerPage.Default(20))
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, paginator)
}

func (ctrl *UserController) Show(response *goyave.Response, request *goyave.Request) {

	userID, err := strconv.ParseUint(request.RouteParams["userID"], 10, 64)
	if err != nil {
		response.Status(http.StatusNotFound)
		return
	}

	user, err := ctrl.UserService.First(request.Context(), uint(userID))
	if response.WriteDBError(err) {
		return
	}

	response.JSON(http.StatusOK, user)
}
