package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

// @title Admin Pro Golang API
// @version 1.0
// @description 基于 Clean Architecture 的企业级后台管理系统 API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email simonchen453@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.



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

	// 3. 初始化 Repository层 (Init Repositories)
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
	
	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 7. 启动服务 (Run) - 支持优雅关闭
	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: r,
	}

	// 在 Goroutine 中启动服务器
	go func() {
		log.Printf("Server starting on %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen failed: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 5 秒超时的 Context 用于优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
