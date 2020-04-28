package err

// ErrCode 错误码
type ErrCode int

// 业务错误码
const (
	Success        ErrCode = 1000                //成功
	Exists         ErrCode = 10001 + iota - 1 //业务数据已存在
	CertErr                                   //身份认证失败，用户名/密码错误
	Invalid                                   //业务数据无效
	Inconsistent                              //数据校验不一致，被篡改
	UploadFailed                              //文件上传失败
	NoVliedParames                            //未验证的参数
	NotFound       ErrCode = 10404            //业务数据不存在
	AddFailed              = 10900 + iota - 8 //添加数据失败
	UpdateFailed                              //修改数据失败
	DelFailed                                 //删除数据失败
	Failed         ErrCode = 10999            //失败
)

var ErrMsg = map[ErrCode]string{
	Success:        "成功",
	Exists:         "业务数据已存在",
	CertErr:        "身份认证失败，用户名/密码错误",
	Invalid:        "业务数据无效",
	Inconsistent:   "数据校验不一致",
	UploadFailed:   "文件上传失败",
	NoVliedParames: "未验证的参数",
	NotFound:       "业务数据不存在",
	Failed:         "失败",
	AddFailed:      "添加失败",
	UpdateFailed:   "修改失败",
	DelFailed:      "删除失败",
}
