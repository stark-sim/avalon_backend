# 更新了数据库字段后，生成 ent 代码，再根据新的 ent 代码生成迁移文件，最后同步更新 graphql 定义
go generate ./pkg/ent &
go generate ./internal/db &
go generate ./pkg/graphql