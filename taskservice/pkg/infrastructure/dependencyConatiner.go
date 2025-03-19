package infrastructure

import (
	responsemodel "github.com/TemaStatham/TaskService/taskservice/pkg/app/response/model"
	responsequery "github.com/TemaStatham/TaskService/taskservice/pkg/app/response/query"
	responseservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/response/service"
	postgres2 "github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/db/postgres"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/grpc"

	taskmodel "github.com/TemaStatham/TaskService/taskservice/pkg/app/task/model"
	taskquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/task/query"
	taskservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/task/service"

	organizationquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/organization/query"

	commentmodel "github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/model"
	commentquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/query"
	commentservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/service"
	userquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/user/query"

	approvemodel "github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/model"
	approveservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/service"

	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/config"
)

type Container struct {
	taskReadRepository taskmodel.TaskReadRepositoryInterface
	taskRepository     taskmodel.TaskRepositoryInterface
	TaskQuery          taskquery.TaskQueryInterface
	TaskService        taskservice.TaskServiceInterface

	responseReadRepository responsemodel.ResponseRepositoryReadInterface
	responseRepository     responsemodel.ResponseRepositoryInterface
	ResponseQuery          responsequery.ResponseQueryInterface
	ResponseService        responseservice.ResponseServiceInterface

	commentReadRepository commentmodel.CommentReadRepositoryInterface
	commentRepository     commentmodel.CommentRepositoryInterface
	CommentQuery          commentquery.CommentQueryInterface
	CommentService        commentservice.CommentServiceInterface

	approveReadRepository approvemodel.ApproveRepositoryInterface
	ApproveService        approveservice.ApproveServiceInterface

	TaskUserReadRepository taskmodel.TaskUserReadRepositoryInterface
	TaskUserRepository     taskmodel.TaskUserRepositoryInterface
	TaskUserQuery          taskquery.TaskUserQueryInterface
	TaskUserService        taskservice.TaskUserServiceInterface

	Client            grpc.ClientInterface
	UserQuery         userquery.UserQueryInterface
	OrganizationQuery organizationquery.OrganizationQueryInterface
}

func NewContainer(config config.Config) *Container {
	db, err := postgres2.NewPostgresGormDB(postgres2.Config{
		Host:     config.DBConfig.Host,
		Port:     config.DBConfig.Port,
		Username: config.DBConfig.Username,
		Password: config.DBConfig.Password,
		DBName:   config.DBConfig.DBName,
		SSLMode:  config.DBConfig.SSLMode,
	})
	if err != nil {
		panic(err)
	}

	responseRepository := postgres2.NewResponsePostgresRepository(db)
	responseQuery := responsequery.NewResponseQuery(responseRepository)
	responseService := responseservice.NewResponseService(responseRepository)

	commentResponse := postgres2.NewCommentsRepository(db)
	commentQuery := commentquery.NewCommentQuery(commentResponse)
	commentService := commentservice.NewCommentService(commentResponse)

	approveRepository := postgres2.NewApproveRepository(db)
	approveStatusRepository := postgres2.NewApproveStatusRepository(db)
	approveService := approveservice.NewApproveService(approveRepository, approveStatusRepository)

	grpcClient, err := grpc.NewGrpcClient(config.Address)
	if err != nil {
		panic(err)
	}

	organizationQuery := organizationquery.NewOrganization(grpcClient)
	userQuery := userquery.NewUserQuery(grpcClient)

	taskUserRepository := postgres2.NewTaskUserPostgresRepository(db)
	taskCategoryRepository := postgres2.NewTaskCategoryRepository(db)
	taskRepository := postgres2.NewTaskPostgresRepository(db, taskUserRepository, taskCategoryRepository)
	taskQuery := taskquery.NewTaskQuery(taskRepository, organizationQuery, userQuery)
	taskService := taskservice.NewTaskService(taskRepository, organizationQuery)

	taskUserQuery := taskquery.NewTaskUserQuery(taskUserRepository)
	taskUserService := taskservice.NewTaskUserService(taskUserRepository)

	return &Container{
		taskReadRepository: taskRepository,
		taskRepository:     taskRepository,
		TaskQuery:          taskQuery,
		TaskService:        taskService,

		responseReadRepository: responseRepository,
		responseRepository:     responseRepository,
		ResponseQuery:          responseQuery,
		ResponseService:        responseService,

		commentReadRepository: commentResponse,
		commentRepository:     commentResponse,
		CommentQuery:          commentQuery,
		CommentService:        commentService,

		approveReadRepository: approveRepository,
		ApproveService:        approveService,

		TaskUserReadRepository: taskUserRepository,
		TaskUserRepository:     taskUserRepository,
		TaskUserQuery:          taskUserQuery,
		TaskUserService:        taskUserService,

		Client:            grpcClient,
		UserQuery:         userQuery,
		OrganizationQuery: organizationQuery,
	}
}
