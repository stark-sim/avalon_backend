package main

import (
	"avalon_backend/configs"
	"avalon_backend/internal/db"
	"avalon_backend/pkg/ent"
	"avalon_backend/pkg/graphql"
	"avalon_backend/pkg/graphql/middlewares"
	"avalon_backend/tools"
	"context"
	"database/sql"
	"database/sql/driver"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	err = configs.InitLogger()
	err = configs.InitConfig()
	err = tools.Init()
	err = db.InitDB()
	if err != nil {
		return
	}
	// 结合 gin 启动 http 服务
	r := gin.Default()
	r.Use(middlewares.WriterMiddleware())
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	err = r.Run(fmt.Sprintf(":%v", configs.Conf.APIConfig.HttpPort))
	if err != nil {
		return
	}
}

func graphqlHandler() gin.HandlerFunc {
	// 创建数据库链接
	client := db.NewDBClient()
	// 初始化 graphql server
	srv := handler.NewDefaultServer(graphql.NewSchema(client))
	// 自定义事务隔离等级
	srv.Use(entgql.Transactioner{
		TxOpener: entgql.TxOpenerFunc(func(ctx context.Context) (context.Context, driver.Tx, error) {
			tx, err := client.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
			if err != nil {
				return nil, nil, err
			}
			ctx = ent.NewTxContext(ctx, tx)
			ctx = ent.NewContext(ctx, tx.Client())
			return ctx, tx, nil
		}),
	})
	// 接上 cookie 校验中间件
	srv.Use(middlewares.NewAuthenticationMiddleware())
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Token, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Hello")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Hello", "World")
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("Test", "/graphql")
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
