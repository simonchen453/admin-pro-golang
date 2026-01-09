package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "admin-pro/docs"

	"admin-pro/internal/config"
	"admin-pro/internal/delivery/http/middleware"
	v1 "admin-pro/internal/delivery/http/v1"
	"admin-pro/internal/infrastructure/persistence"
	"admin-pro/internal/infrastructure/persistence/mysql"
	"admin-pro/internal/usecase"
)

func main() {
	// 1. 加载配置 (Load Config)
	// 读取 config.yaml 文件，包含数据库连接、端口等信息
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. 初始化数据库连接 (Init DB)
	// 使用 GORM 连接 MySQL
	db, err := persistence.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to init DB: %v", err)
	}

	// 3. 初始化 Repository 层 (Init Repositories)
	// Repository 负责直接操作数据库，实现了 Domain 层定义的接口
	// 依赖注入：将数据库连接实例 db 注入到 Repository 中
	userRepo := mysql.NewUserRepository(db)
	roleRepo := mysql.NewRoleRepository(db)
	menuRepo := mysql.NewMenuRepository(db)
	deptRepo := mysql.NewDeptRepository(db)
	postRepo := mysql.NewPostRepository(db)
	dictRepo := mysql.NewDictRepository(db)
	configRepo := mysql.NewConfigRepository(db)
	noticeRepo := mysql.NewNoticeRepository(db)
	logRepo := mysql.NewLogRepository(db)
	monitorRepo := mysql.NewMonitorRepository(db)
	jobRepo := mysql.NewJobRepository(db)
	genRepo := mysql.NewGenRepository(db)

	// 4. 初始化 Usecase 层 (Init Usecases)
	// Usecase 负责业务逻辑（如校验、流程控制）
	// 依赖注入：将 Repository 实例注入到 Usecase 中
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg)
	userUsecase := usecase.NewUserUsecase(userRepo, roleRepo, menuRepo)
	deptUsecase := usecase.NewDeptUsecase(deptRepo)
	postUsecase := usecase.NewPostUsecase(postRepo)
	dictUsecase := usecase.NewDictUsecase(dictRepo)
	configUsecase := usecase.NewConfigUsecase(configRepo)
	noticeUsecase := usecase.NewNoticeUsecase(noticeRepo)
	logUsecase := usecase.NewLogUsecase(logRepo)
	monitorUsecase := usecase.NewMonitorUsecase(monitorRepo)
	jobUsecase := usecase.NewJobUsecase(jobRepo)
	genUsecase := usecase.NewGenUsecase(genRepo)

	// 5. 初始化 Gin Web 框架 (Init Gin)
	r := gin.Default()
	r.Use(middleware.Cors()) // 开启跨域允许

	// 6. 注册路由 Handler (Init Handlers)
	// Handler 负责处理 HTTP 请求，解析参数，调用 Usecase，返回 JSON
	// 依赖注入：将 Usecase 实例注入到 Handler 中
	v1.NewAuthHandler(r, authUsecase, cfg)
	v1.NewUserHandler(r, userUsecase, middleware.JWTAuth(cfg)) // 需要 JWT 认证
	v1.NewDeptHandler(r, deptUsecase, middleware.JWTAuth(cfg))
	v1.NewPostHandler(r, postUsecase, middleware.JWTAuth(cfg))
	v1.NewDictHandler(r, dictUsecase, middleware.JWTAuth(cfg))
	v1.NewConfigHandler(r, configUsecase, middleware.JWTAuth(cfg))
	v1.NewNoticeHandler(r, noticeUsecase, middleware.JWTAuth(cfg))
	v1.NewLogHandler(r, logUsecase, middleware.JWTAuth(cfg))
	v1.NewMonitorHandler(r, monitorUsecase, middleware.JWTAuth(cfg))
	v1.NewJobHandler(r, jobUsecase, middleware.JWTAuth(cfg))
	v1.NewGenHandler(r, genUsecase, middleware.JWTAuth(cfg))
	
	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 7. 启动服务 (Run)
	log.Printf("Server starting on %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}
