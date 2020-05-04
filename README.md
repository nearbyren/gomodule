learning 学习目录

conf：用于存储配置文件   
middleware：应用中间件    
models：应用数据库模型  
pkg：第三方包    
routers 路由逻辑处理  
runtime：应用运行时数据

###阅读go.mod文件
#####新增 replace 配置项 因为开始以下模块没推送到远程，无法下载下来，因此需要使用replace，将其指定读取本地的模块路径，这样就可以解决本地模块读取问题
replace (
    github.com/nearbyren/gomodule//pkg/setting => ~/go-application/gomodule/pkg/setting
    github.com/nearbyren/gomodule//conf    	  => ~/go-application/gomodule/pkg/conf
    github.com/nearbyren/gomodule//middleware  => ~/go-application/gomodule/middleware
    github.com/nearbyren/gomodule//models 	  => ~/go-application/gomodule/models
    github.com/nearbyren/gomodule//routers 	  => ~/go-application/gomodule/routers
)