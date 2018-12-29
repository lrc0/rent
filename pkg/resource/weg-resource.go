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

//WegResource .
type WegResource struct {
}

//WebService .
func (w WegResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/wegdosages").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"wegdosage"}

	ws.Route(ws.POST("/").To(w.addWegDosage).
		Doc("add dosage").
		Reads(types.WegDosage{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil))

	ws.Route(ws.GET("/").To(w.findAllWegDosage).
		// docs
		Doc("get all weg dosage").
		Param(ws.QueryParameter("page_index", "index")).
		Param(ws.QueryParameter("page_size", "size")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(types.WegDosage{}))

	ws.Route(ws.GET("/{room_name}").To(w.findWegDosage).
		// docs
		Doc("get weg dosage by room name").
		Param(ws.PathParameter("room_name", "房间名")).
		Param(ws.QueryParameter("page_index", "index")).
		Param(ws.QueryParameter("page_size", "size")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(types.WegDosage{}))

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

func (w *WegResource) addWegDosage(request *restful.Request, response *restful.Response) {
	var weg types.WegDosage
	err := request.ReadEntity(&weg)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	err = mapper.AddWegDosage(&weg)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
}

func (w *WegResource) findWegDosage(request *restful.Request, response *restful.Response) {
	roomName := request.PathParameter("room_name")
	size, _ := strconv.Atoi(request.QueryParameter("page_size"))
	index, _ := strconv.Atoi(request.QueryParameter("page_index"))
	var pa = new(page.Pageable)
	pa.PageIndex = index
	pa.PageSize = size

	rs, err := mapper.FindWegDosage(roomName, pa)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, rs)
}

func (w *WegResource) findAllWegDosage(request *restful.Request, response *restful.Response) {
	size, _ := strconv.Atoi(request.QueryParameter("page_size"))
	index, _ := strconv.Atoi(request.QueryParameter("page_index"))
	var pa = new(page.Pageable)
	pa.PageIndex = index
	pa.PageSize = size

	rs, err := mapper.FindAllWegDosage(pa)
	if err != nil {
		util.WriteBadRequestError(response, err.Error())
		return
	}
	util.WriteSuccessEntity(response, rs)
}
