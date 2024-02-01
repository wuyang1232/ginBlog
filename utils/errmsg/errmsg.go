package errmsg

// 状态码
const (
	SUCCESS = 200
	ERROR   = 500

	//code = 1000 ... 用户模块的错误

	//code = 2000 ... 文章模块的错误

	//code = 3000 ... 分类模块的错误

)

// 声明一个字典
var codemsg = map[int]string{}

// 输出错误信息的函数
func GetErrMsg(code int) string {

	return codemsg[code]
}
