package resource

import (
	"strconv"
	"time"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/xorm-page/page"

	"rentmanagement/pkg/mapper"
	"rentmanagement/pkg/types"
	"rentmanagement/pkg/util"
)

//TenantResource .
type TenantResource struct {
}

//WebService .
func (t TenantResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/tenants").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"tenant"}

	ws.Route(ws.POST("/").To(t.addTenant).
		Doc("add tenant").
		Reads(types.Tenant{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil))

	ws.Route(ws.GET("/").To(t.findAllTenant).
		// docs
		Doc("get all tenant").
		Param(ws.QueryParameter("page_index", "index")).
		Param(ws.QueryParameter("page_size", "size")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(types.Tenant{}))

	ws.Route(ws.GET("/{room_name}").To(t.findTenant).
		// docs
		Doc("get Tenant by room name").
		Param(ws.PathParameter("room_name", "房间名")).
		Param(ws.QueryParameter("page_index", "index")).
		Param(ws.QueryParameter("page_size", "size")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(types.Tenant{}))

	// ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
	// 	// docs
	// 	Doc("update a user").
	// 	Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(types.User{})) // from the request

	// ws.Route(ws.PUT("").To(u.createUser).
	// 	// docs
	// 	Doc("create a user").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(types.User{})) // from the request

	// ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
	// 	// docs
	// 	Doc("delete a user").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	return ws
}

func (t *TenantResource) addTenant(request *restful.Request, response *restful.Response) {
	var tenant types.Tenant
	err := request.ReadEntity(&tenant)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	tenant.CreateTime = time.Now()
	tenant.UpdateTime = time.Now()
	err = mapper.AddTenant(&tenant)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, types.Message{Message: "success to add tenant"})
}

func (t *TenantResource) findAllTenant(request *restful.Request, response *restful.Response) {
	size, _ := strconv.Atoi(request.QueryParameter("page_size"))
	index, _ := strconv.Atoi(request.QueryParameter("page_index"))
	var pa = new(page.Pageable)
	pa.PageIndex = index
	pa.PageSize = size

	rs, err := mapper.FindAllTenant(pa)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, rs)
}

func (t *TenantResource) findTenant(request *restful.Request, response *restful.Response) {
	roomName := request.PathParameter("room_name")
	size, _ := strconv.Atoi(request.QueryParameter("page_size"))
	index, _ := strconv.Atoi(request.QueryParameter("page_index"))
	var pa = new(page.Pageable)
	pa.PageIndex = index
	pa.PageSize = size

	rs, err := mapper.FindTenant(roomName, pa)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, rs)
}
