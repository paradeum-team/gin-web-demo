module gin-web-demo

go 1.12

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/flyleft/gprofile v0.0.0-20190121091042-4c613f874133
	github.com/gin-gonic/gin v1.4.0
	github.com/go-resty/resty v0.0.0-20190619084753-e284be3e6edc
	github.com/kataras/golog v0.0.0-20180321173939-03be10146386
	github.com/kataras/pio v0.0.0-20190103105442-ea782b38602d // indirect
	github.com/olebedev/config v0.0.0-20190528211619-364964f3a8e4 // indirect
	github.com/swaggo/gin-swagger v1.1.0
	github.com/swaggo/swag v1.5.0
	golang.org/x/tools v0.0.0-20190322203728-c1a832b0ad89
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

replace github.com/gin-gonic/gin v1.4.0 => github.com/paradeum-team/gin v1.4.4

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
