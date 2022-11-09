# zero-paopao

- [zero-paopao](#zero-paopao)
  - [reference](#reference)
  - [开发](#开发)
    - [后端](#后端)
      - [工具准备](#工具准备)
      - [model 生成](#model-生成)
      - [后端服务+swagger](#后端服务swagger)
    - [前端](#前端)
      - [环境准备](#环境准备)
      - [构建部署](#构建部署)
      - [参考](#参考)
  - [issue](#issue)
    - [go](#go)
    - [fish config](#fish-config)

paopao-ce implement by go-zero & vue3 & typescript

## reference

- [paopao-ce](https://github.com/rocboss/paopao-ce)
- [go-zero](https://github.com/zeromicro/go-zero)

## 开发

### 后端

#### 工具准备

1. goctl安装,参考<https://go-zero.dev/cn/docs/goctl/goctl>
    go1.16以上

    ```sh
    GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
    ```

2. goctl-swagger

   ```sh
   GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
   ```

3. mysql+redis+zinc部署(docker)

    ```sh
    docker-compose up -d
    ```

#### model 生成

配置scripts下的gen-model-from-db.sh脚本,填入参数,调用脚本会生成model,这里默认启用cache

#### 后端服务+swagger

开发阶段:
使用脚本 api/backend.sh,swagger通过goctl及插件生成通过docker部署,后端服务使用go run直接运行

### 前端

#### 环境准备

1. nodejs安装,按Microsoft文档,在wsl上安装[doc](https://learn.microsoft.com/zh-cn/windows/dev-environment/javascript/nodejs-on-wsl)

   ```sh
   curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/master/install.sh | bash

   nvm -v

   #设置国内镜像
   nvm npm_mirror https://npmmirror.com/mirrors/npm/
   #安装yarn
   npm install -g yarn --registry=https://registry.npmmirror.com
   #设置镜像源 开源镜像 https://npmmirror.com/
   yarn config set registry https://registry.npmmirror.com -g
   yarn config set sass_binary_site "https://npmmirror.com/mirrors/node-sass/" -g

   #可选,安装cnpm
   npm install -g cnpm --registry=https://registry.npmmirror.com
   #设置镜像
   npm config set registry  https://registry.npmmirror.com

   # 下载项目依赖
   yarn

   # 构建项目
   yarn build
   ```

#### 构建部署

使用yarn构建过后可以使用docker部署(直接使用yarn build结果)

```sh
bash front.sh
```

#### 参考

- npm镜像 <https://blog.csdn.net/weixin_45182409/article/details/117981169>

## issue

### go

数据文件夹名称改为.data,为data时go mod tidy会扫描该文件夹,无权限访问时会报错

### fish config

- [fish add PATH](https://blog.csdn.net/yangxiang92/article/details/20057975)
  在配置文件里添加,set PATH xxxx $PATH
- fish_config,open config web page,可视化编辑
