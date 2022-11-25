# GoLearn-Gin
`Gin`框架学习使用,并实践常用包在Gin框架中的集成和使用。[Github地址: https://github.com/DivenZhong/GoLearn-Gin.git](https://github.com/DivenZhong/GoLearn-Gin.git)

用于diven zhong个人学习实践,转载请说明

# 项目流程

```
step-1:初始化配置文件、集成服务，并赋值到相应的全局变量
step-2:启动服务
       1.创建一个gin运行引擎，初始化RouterGroup(所有路由处理的实现方法以及一个分组方法),创建Context对象池sync.Pool，创建方法树
       2.注册公共中间件(日志和恢复等中间件)
       3.注册路由，将URL的方法链注册到相应的方法树中
       4.获取自定义http配置
       5.调用net/http的tcp监听和http服务，开始进入http,启动WEB服务
step-3：处理HTTP请求
       1.从对象池中获取一个Context
	   2.对Context进行必要的初始化
	   3.处理该HTTP请求
	     >获得该HTTP请求的方法树
		 >找到该URL的处理方法链
		 >通过上下文的Next方法实现中间件的调用以及方法的调用
	   4.完成对HTTP的请求后，释放该Context
```

# 目录介绍

```
├── api // 接口
├── config // 配置
├── core // 核心代码
├── global // 全局变量和常量
├── initialize // 初始化相关
├── logs // 日志目录
├── middleware // 中间件
├── main.go // 启动文件
├── model // 实体和表结构一一对应   
    └── request // 请求参数结构体
    └── response // 返回参数结构、返回函数封装
├── router // 路由
├── test // 单元测试目录
└── utils // 工具包
```

# 集成服务
```
viper 配置文件管理组件(参考:https://github.com/spf13/viper):
   1.设置默认值
   2.支持 JSON/YAML/envfile/Java properties 等多种格式的配置文件
   3.实时监控和重新读取配置文件
   4.写入运行时配置
   5.读取环境变量、命令行参数、远程配置系统、缓冲区等配置信息
   6.提供别名系统，以便在不破坏现有代码的情况下轻松重命名参数

gorm:
    1.对象关系映射
	2.数据库迁移
	3.通过Tag设置字段的等级权限、数据结构、格式、外键关联等
	 column //指定db列名
     type   //列数据类型
     size //指定列大小，例如：size:256
     primaryKey //指定列为主键
     unique //指定列为唯一
     default //指定列的默认值
     precision //指定列的精度
     scale //指定列大小
     not null //指定列为 NOT NULL
     autoIncrement //指定列为自动增长
     embedded //嵌套字段
	 embeddedPrefix //嵌入字段的列名前缀
	 autoCreateTime //创建时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoCreateTime:nano
     autoUpdateTime //更新时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:mill
     index //根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 索引 获取详情
     uniqueIndex //与 index 相同，但创建的是唯一索引
     check //创建检查约束，例如 check:age > 13，查看 约束 获取详情
     <- //设置字段写入的权限，<-:create 只创建、<-:update 只更新、<-:false 无写入权限、<- 创建和更新权限
     -> //设置字段读的权限，->:false 无读权限
     - //忽略该字段，-无读写权限
zap 实现高性能的日志组件：
	1.自己实现 json Encoder，指明类型避免了反射，每种类型直接提供了转换 []byte + append 的函数
    2.预分配byte slice字节片段
    3.通过 sync.Pool实现对象复用，避免竞争
	4.日志级别逻辑判断拦截，避免一切不必要的开销	
testify:单元测试
redis：内存缓存
elastic：日志收集
JWT ：Token令牌授权和认证
```

# 集成Grpc-go、Protobuf( https://github.com/DivenZhong/GoLearn-Grpc-Go)

```
#protobuf简介
Protocol Buffers(protobuf)：与编程语言无关，与程序运行平台无关的数据序列化协议以及接口定义语言(IDL: interface definition language)。

#Grpc-Go的核心概念
    1.编译器protoc：
	2.Go语言的protobuf插件:
	     编程语言的protobuf插件，搭配protoc编译器，根据.proto文件生成对应编程语言的代码，实现各自语言的protobuf协议。
		 在发送请求和接受响应的时候，完成对应的编码和解码工作，将你即将发送的数据编码成gRPC能够传输的形式，又或者将即将接收到的数据解码为编程语言能够理解的数据格式
	相应的指令:	 
	     protoc --go_out=. common.proto
         protoc --go-grpc_out=. common.proto
#接口开发流程步骤	
    1.定义.proto文件，包括消息体和rpc服务接口定义
    2.使用protoc命令来编译.proto文件，用于生成xx.pb.go和xx_grpc.pb.go文件
    3.在服务端实现rpc里定义的方法
    4.客户端调用rpc方法，获取响应结果	
	
#consul
      1.安装运行consul,使得127.0.0.1:8500端口可以正常访问
	  2.启动grpc服务端，完成consul服务注册，把grpc服务端服务IP、端口、服务名称等注册到consul上面
	  3.启动grpc客户端，根据服务名称去consul拿到grpc服务端的服务IP、端口等信息，并建立起连接
	  4.正常走grpc之间调用的流程

#etcd
      1.安装运行etcd,使得127.0.0.1:2379端口可以正常访问,借助etcdkeeper可视化etcd
	  2.启动grpc服务端，
	      1)注册etcd租约，
		  2)完成etcd服务注册，把grpc服务端服务IP、端口、服务名称等注册到etcd上面，
		  3)租约异步续租	  
		  4)启用监听，监听某个Key值是否改变，并完成后续流程
	  3.启动grpc客户端，根据服务名称去etcd拿到grpc服务端的服务IP、端口等信息，并建立起连接
	  4.正常走grpc之间调用的流程
```

# 集成Cobra(https://github.com/DivenZhong/GoLearn-Cobra)

```

#Cobra的核心概念   
   CoBra 命令行交互工具：
     命令（Command）：就是需要执行的操作；
     参数（Arg）：命令的参数，即要操作的对象；
     选项（Flag）：命令选项可以调整命令的行为。
	
#接口	
   cobra 提供了非常丰富的特性和定制化接口，例如：
         设置钩子函数，在命令执行前、后执行某些操作；	
	
#脚手架Scaffold示例
   CoBra 生成Scaffold示例：github路径下的生成的main.exe，实现这样一个功能，根据传入的年、月，打印这个月的日历。如果没有传入选项，使用当前的年、月
		  用法：main.exe date --year 2019 --month 12
```

# 中间件

```
   1.全局中间件：
                gin.Engine结构体的Use()方法
				router := New()
	            router.Use(Logger(), Recovery())
   2.分组使用中间件：
                user := router.Group("user", gin.Logger(),gin.Recovery())
                {
                }
   3.单个路由使用中间件：
                router.GET("/test",gin.Recovery(),func(c *gin.Context){
                       c.JSON(200,"test")
                })   
				
   自定义gin中间件有两种写法。
   1)定义一个方法接收一个*gin.Context类型的参数，和handler的写法是一样的。代码参考如下：
           func MyMiddleware(c *gin.Context){    
           }
           router.Use(MyMiddleware)			   
   2)定义一个返回值为HandlerFunc类型的函数，参考代码如下：
           func MyMiddleware(){
             return func(c *gin.Context){
            }
           }
           router.Use(MyMiddleware())
```

