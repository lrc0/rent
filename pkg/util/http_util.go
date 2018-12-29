package util

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

// 主机错误码
const (
	NodeBadRequestCode          = 1001400            // 主机请求参数错误码
	NodeUnauthorizedCode        = 1001401            // 主机未授权错误码
	NodeNotFoundCode            = 1001404            // 主机资源不存在错误码
	NodeInternalServerErrorCode = 1001500            // 系统错误错误码
	ContentType                 = "application/json" // 返回数据格式

)

// ErrorResponse 错误返回实体
type ErrorResponse struct {
	ErrorCode int    `json:"err_code,omitempty"` //错误编号
	ErrorMsg  string `json:"err_msg,omitempty"`  //错误消息
	ErrorDesc string `json:"err_desc,omitempty"` //错误具体描述
}

// NewErrorResponse 创建一个错误返回实体
func NewErrorResponse(errorCode int, errorMsg, errorDesc string) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode: errorCode,
		ErrorDesc: errorDesc,
		ErrorMsg:  errorMsg}
}

// WriteEntity 返回json报文
func WriteEntity(response *restful.Response, httpStatus int, entity interface{}) {
	response.WriteHeaderAndJson(httpStatus, entity, ContentType)
	return
}

// WriteSuccessEntity 返回请求成功json报文
func WriteSuccessEntity(response *restful.Response, entity interface{}) {
	WriteEntity(response, http.StatusOK, entity)
}

// WriteCreatedEntity 返回创建成功json报文
func WriteCreatedEntity(response *restful.Response, entity interface{}) {
	WriteEntity(response, http.StatusCreated, entity)
}

// WriteError 返回json错误
func WriteError(response *restful.Response, httpStatus int, msg interface{}) {
	response.WriteHeaderAndJson(httpStatus, msg, ContentType)
	return
}

// WriteBadRequestError 返回参数错误
func WriteBadRequestError(response *restful.Response, errDesc string) {

	WriteError(response, http.StatusBadRequest, NewErrorResponse(NodeBadRequestCode, "参数错误", errDesc))
}

// WriteUnauthorizedError 返回未授权错误
func WriteUnauthorizedError(response *restful.Response, errDesc string) {

	WriteError(response, http.StatusBadRequest, NewErrorResponse(NodeUnauthorizedCode, "授权错误", errDesc))
}

// WriteNotFoundError 返回资源不存在错误
func WriteNotFoundError(response *restful.Response, errDesc string) {

	WriteError(response, http.StatusNotFound, NewErrorResponse(NodeNotFoundCode, "资源不存在", errDesc))
}

// WriteInternalServerError 系统错误
func WriteInternalServerError(response *restful.Response) {
	WriteError(response, http.StatusInternalServerError, NewErrorResponse(NodeInternalServerErrorCode, "系统错误", "请联系管理员"))
}
