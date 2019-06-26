## swagger 使用

### 安装 swag
- 1、go get

```
$ go get -u github.com/swaggo/swag/cmd/swag
```

若 $GOPATH/bin 没有加入$PATH中，你需要执行将其可执行文件移动到$GOBIN下

```
mv $GOPATH/bin/swag /usr/local/go/bin
```

- 2、gopm get

该包有引用golang.org上的包，若无科学上网，你可以使用 gopm 进行安装

```
gopm get -g -v github.com/swaggo/swag/cmd/swag

cd $GOPATH/src/github.com/swaggo/swag/cmd/swag

go install
```

**同理将其可执行文件移动到$GOBIN下**


### 验证是否安装成功

```
$ swag -v
swag version v1.5.1
```

## 使用
### 初始化目录结构（与生成api 文档的命令相同）

进入到项目的根目录，与main.go相同的目录时，执行下面的命令：

```
$ swag init 
2019/06/25 16:24:56 Generate swagger docs....
2019/06/25 16:24:56 Generate general API Info
2019/06/25 16:24:56 create docs.go at  docs/docs.go
2019/06/25 16:24:56 create swagger.json at  docs/swagger.json
2019/06/25 16:24:56 create swagger.yaml at  docs/swagger.yaml

```
执行完毕后会在项目根目录下生成docs

```
docs/
├── docs.go
├── swagger.json
└── swagger.yaml
```
### 引包

```
"gin-web-demo/docs" #本项目，上一步生成的
"github.com/swaggo/gin-swagger"
"github.com/swaggo/gin-swagger/swaggerFiles"
```     

### 集成gin ，使用

```
// programatically set swagger info
docs.SwaggerInfo.Title = "PLD:ONLINE API"
docs.SwaggerInfo.Description = "This is a pld server online restfull api ."
docs.SwaggerInfo.Version = "1.0"
docs.SwaggerInfo.Host = "localhost:8080"
docs.SwaggerInfo.BasePath ="/pld"
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

### 在具体方法上，编写 API注释
Swagger 中需要将相应的注释或注解编写到方法上，再利用生成器自动生成说明文件

gin-swagger 给出的范例：

**GET**
```
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c *gin.Context) {
```

**POST**
```
// @Summary user login with username and pwd
// @Description user login  with username and pwd
// @Accept  json
// @Produce  json
// @Param user body vo.LoginJSON true "user body" #指定数据类型，在方法内使用 ctx.ShouldBindJSON(&obj)
// @Success 200 {string} json
// @Failure 401 {string} json
// @Failure 500 {string} json
// @Router /v1/pwlogin [post]
func V1PasswordLoginfunc(c *gin.Context) {
    var json vo.LoginJSON
	if err:=c.ShouldBindJSON(&json);err==nil {
	
	}
	
```





### 简要的列举一些 [swagger-api](https://github.com/swaggo/swag)
#### API Operation
|annotation	|description|
|:-----|:-----|
|description|	A verbose explanation of the operation behavior.|
|id	|A unique string used to identify the operation. Must be unique among all API operations.|
|tags|	A list of tags to each API operation that separated by commas.|
|summary|	A short summary of what the operation does.|
|accept	|A list of MIME types the APIs can consume. Value MUST be as described under Mime Types.|
|produce|	A list of MIME types the APIs can produce. Value MUST be as described under Mime Types.|
|param|	Parameters that separated by spaces. param name,param type,data type,is mandatory?,comment attribute(optional)|
|security|	Security to each API operation.|
|success|	Success response that separated by spaces. return code,{param type},data type,comment|
|failure|	Failure response that separated by spaces. return code,{param type},data type,comment|
|header|	Header in response that separated by spaces. return code,{param type},data type,comment|
|router	|Path definition that separated by spaces. path,[httpMethod]|
 

#### Mime Types


#### Param Type
- query
- path
- header
- body
- formData

#### Data Type
- string (string)
- integer (int, uint, uint32, uint64)
- number (float32)
- boolean (bool)
- user defined struct