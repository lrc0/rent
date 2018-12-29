package resource

import (
	"strconv"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/xorm-page/page"

	"rentmanagement/pkg/mapper"
	"rentmanagement/pkg/types"
	"rentmanagement/pkg/util"
)

// HouseResource .
type HouseResource struct {
}

// WebService creates a new service that can handle REST requests for User resources.
func (h HouseResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"house"}

	ws.Route(ws.GET("").To(h.findAllHouseInfo).
		// docs
		Doc("get all house info").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil))

	ws.Route(ws.GET("/{name}").To(h.findHouseInfo).
		// docs
		Doc("get a house info").
		Param(ws.PathParameter("name", "房东姓名")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.POST("").To(h.addHouseInfo).
		// docs
		Doc("add a house info").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(types.HouseInfo{})) // from the request

	ws.Route(ws.DELETE("/{uuid}").To(h.deleteHouseInfo).
		// docs
		Doc("delete a house info").
		Param(ws.PathParameter("uuid", "uuid")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	return ws
}

func (h HouseResource) addHouseInfo(request *restful.Request, response *restful.Response) {
	var house types.HouseInfo
	err := request.ReadEntity(&house)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	err = mapper.AddHouseInfo(&house)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, types.Message{Message: "success"})
}

func (h HouseResource) findAllHouseInfo(request *restful.Request, response *restful.Response) {
	size, _ := strconv.Atoi(request.QueryParameter("page_size"))
	index, _ := strconv.Atoi(request.QueryParameter("page_index"))
	var pa = new(page.Pageable)
	pa.PageIndex = index
	pa.PageSize = size

	rs, err := mapper.FindAllHouseInfo(pa)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, rs)
}

func (h HouseResource) findHouseInfo(request *restful.Request, response *restful.Response) {
	uuid := request.PathParameter("uuid")
	size, _ := strconv.Atoi(request.QueryParameter("page_size"))
	index, _ := strconv.Atoi(request.QueryParameter("page_index"))
	var pa = new(page.Pageable)
	pa.PageIndex = index
	pa.PageSize = size

	rs, err := mapper.FindHouseInfo(uuid, pa)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, rs)
}

func (h HouseResource) deleteHouseInfo(request *restful.Request, response *restful.Response) {
	uuid := request.PathParameter("uuid")

	err := mapper.DeleteHouseInfo(uuid)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, types.Message{Message: "success"})
}
