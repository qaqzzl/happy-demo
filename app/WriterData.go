package app



type WriterData struct {
	Status	string		`json:"status"`
	Code 	int			`json:"code"`
	Msg		string		`json:"msg"`
	Data	interface{}	`json:"data"`
}

var apiCode [6000]string

func Error(code int, data interface{}) WriterData {
	return WriterData {
		Status: "ERR",
		Msg:	apiCode[code],
		Code:	code,
		Data:	data,
	}
}

func Success(code int, data interface{}) WriterData {
	return WriterData {
		Status: "OK",
		Msg:	apiCode[code],
		Code:	code,
		Data:	data,
	}
}


func init() {
	//OK
	apiCode[2000] = "ok"

	//ERROR
	apiCode[5000] = "服务器错误"
}