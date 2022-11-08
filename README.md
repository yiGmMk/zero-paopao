# zero-paopao

paopao-ce implement by go-zero & vue3 & typescript

## reference

- [paopao-ce](https://github.com/rocboss/paopao-ce)
- [go-zero](https://github.com/zeromicro/go-zero)

### go-zero

#### model 生成

配置scripts下的gen-model-from-db.sh脚本,填入参数,调用脚本会生成model,这里默认启用cache

#### 后端服务+swagger

开发阶段:
使用脚本 api/backend.sh,swagger通过goctl及插件生成通过docker部署,后端服务使用go run直接运行

### 前端

## issue

### go

数据文件夹名称改为.data,为data时go mod tidy会扫描该文件夹,无权限访问时会报错

### fish config

- [fish add PATH](https://blog.csdn.net/yangxiang92/article/details/20057975)
- fish_config,open config web page,可视化编辑
