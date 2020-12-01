package error_def


var (
	CODE_SUCCESS  int32= 0
	//数据转化类型错误码， 范围 -100 ~ -199
	CODE_ERR_MARSHAL_FAIL     int32 = -100
	CODE_ERR_UNMARSHAL_FAIL   int32 = -101
	CODE_ERR_SET_REDIS_FAIL   int32 = -103
	CODE_ERR_GET_REDIS_FAIL   int32 = -104
	CODE_ERR_SET_DB_FAIL      int32 = -105
	CODE_ERR_GET_DB_FAIL      int32 = -106

	// 业务类型错误码 （根据类型业务区分）-200 ~ -999
	CODE_ERR_PARAM_EMPTY        int32 = -201
	CODE_ERR_PARAM_INVALID      int32 = -202

	CODE_ERR_UNKNOWN            int32 = -999
)