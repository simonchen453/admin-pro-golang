# 权限标识符映射表

本文档定义了系统中所有接口对应的权限标识符，采用 `模块:资源:操作` 格式。

## 权限命名规范

格式：`{module}:{resource}:{action}`

- **module**: 模块名（system, monitor等）
- **resource**: 资源名（dept, user, role等）
- **action**: 操作类型（list, add, edit, remove, query, export）

## 系统管理模块 (system)

### 部门管理 (dept)
- `system:dept:list` - 查看部门列表
- `system:dept:query` - 查看部门详情
- `system:dept:add` - 新增部门
- `system:dept:edit` - 编辑部门
- `system:dept:remove` - 删除部门

### 用户管理 (user)
- `system:user:list` - 查看用户列表
- `system:user:query` - 查看用户详情
- `system:user:add` - 新增用户
- `system:user:edit` - 编辑用户
- `system:user:remove` - 删除用户
- `system:user:reset` - 重置密码

### 角色管理 (role)
- `system:role:list` - 查看角色列表
- `system:role:query` - 查看角色详情
- `system:role:add` - 新增角色
- `system:role:edit` - 编辑角色
- `system:role:remove` - 删除角色

### 菜单管理 (menu)
- `system:menu:list` - 查看菜单列表
- `system:menu:query` - 查看菜单详情
- `system:menu:add` - 新增菜单
- `system:menu:edit` - 编辑菜单
- `system:menu:remove` - 删除菜单

### 岗位管理 (post)
- `system:post:list` - 查看岗位列表
- `system:post:query` - 查看岗位详情
- `system:post:add` - 新增岗位
- `system:post:edit` - 编辑岗位
- `system:post:remove` - 删除岗位

### 字典管理 (dict)
- `system:dict:list` - 查看字典列表
- `system:dict:query` - 查看字典详情
- `system:dict:add` - 新增字典
- `system:dict:edit` - 编辑字典
- `system:dict:remove` - 删除字典

### 参数配置 (config)
- `system:config:list` - 查看参数列表
- `system:config:query` - 查看参数详情
- `system:config:add` - 新增参数
- `system:config:edit` - 编辑参数
- `system:config:remove` - 删除参数

### 通知公告 (notice)
- `system:notice:list` - 查看通知列表
- `system:notice:query` - 查看通知详情
- `system:notice:add` - 新增通知
- `system:notice:edit` - 编辑通知
- `system:notice:remove` - 删除通知

### 日志管理 (log)
- `system:log:list` - 查看日志列表
- `system:log:query` - 查看日志详情
- `system:log:remove` - 删除日志

## 系统监控模块 (monitor)

### 在线用户 (online)
- `monitor:online:list` - 查看在线用户
- `monitor:online:forceLogout` - 强制下线

### 服务监控 (server)
- `monitor:server:query` - 查看服务器信息

## 定时任务模块 (job)

### 任务管理 (job)
- `job:job:list` - 查看任务列表
- `job:job:query` - 查看任务详情
- `job:job:add` - 新增任务
- `job:job:edit` - 编辑任务
- `job:job:remove` - 删除任务
- `job:job:execute` - 执行任务

### 任务日志 (log)
- `job:log:list` - 查看任务日志
- `job:log:query` - 查看日志详情

## 代码生成模块 (tool)

### 代码生成 (gen)
- `tool:gen:list` - 查看表列表
- `tool:gen:code` - 生成代码

## 通配符权限

- `*:*:*` - 超级管理员（所有权限）
- `system:*:*` - 系统管理模块所有权限
- `system:dept:*` - 部门管理所有权限

## 使用示例

```go
// 在 Handler 路由注册时使用
g.GET("/list", permMw("system:dept:list"), handler.List)
g.POST("", permMw("system:dept:add"), handler.Add)
```

## 数据库配置

这些权限标识符应该在 `sys_menu_tbl` 表的 `col_permission` 字段中配置：

```sql
INSERT INTO sys_menu_tbl (col_menu_name, col_permission, col_menu_type) 
VALUES ('部门列表', 'system:dept:list', 'F');
```
