package resource

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"

	"rentmanagement/pkg/message"
	"rentmanagement/pkg/types"
	"rentmanagement/pkg/util"
)

//MailResource .
type MailResource struct {
}

//WebService .
func (m MailResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/mails").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"mail"}

	ws.Route(ws.POST("/").To(m.sendMail).
		Doc("send mail").
		Param(ws.QueryParameter("receiver", "收件人")).
		Param(ws.QueryParameter("theme", "邮件主题")).
		Reads(EmailBody{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil))
	return ws
}

//EmailBody .
type EmailBody struct {
	Message string `json:"message"`
}

func (m *MailResource) sendMail(request *restful.Request, response *restful.Response) {
	receiver := request.QueryParameter("receiver")
	theme := request.QueryParameter("theme")

	var eb EmailBody
	err := request.ReadEntity(&eb)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}

	mess := eb.Message
	err = message.SendMail(receiver, mess, theme)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, types.Message{Message: "success"})
}
