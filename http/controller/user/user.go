package user

import (
	"net/http"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/typeutil"
	"goyave.dev/template/http/dto"
	"goyave.dev/template/service"
	"goyave.dev/template/service/user"
)

type UserController struct {
	goyave.Component
	UserService *user.Service
}

func (ctrl *UserController) Init(server *goyave.Server) {
	ctrl.UserService = server.Service(service.User).(*user.Service)
	ctrl.Component.Init(server)
}

func (ctrl *UserController) RegisterRoutes(router *goyave.Router) {
	subrouter := router.Subrouter("/users")
	subrouter.CORS(cors.Default())

	subrouter.Get("/", ctrl.Index).ValidateQuery(IndexRequest)
}

func (ctrl *UserController) Index(response *goyave.Response, request *goyave.Request) {

	query := typeutil.MustConvert[dto.Index](request.Query)

	paginator, err := ctrl.UserService.Paginate(query.Page.Default(1), query.PerPage.Default(20))
	if response.WriteDBError(err) {
		return
	}

	// Convert to DTO and write response
	dto := typeutil.MustConvert[database.PaginatorDTO[dto.User]](paginator)
	response.JSON(http.StatusOK, dto)
}
