package resource

import (
	// "strconv"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	// "github.com/xorm-page/page"

	// "rentmanagement/pkg/mapper"
	"rentmanagement/pkg/types"
	"rentmanagement/pkg/util"
)

// HealthResource .
type HealthResource struct {
}

// WebService creates a new service that can handle REST requests for User resources.
func (h HealthResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/health").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"health"}

	ws.Route(ws.GET("").To(h.checkHealth).
		// docs
		Doc("get all house info").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil))
	return ws
}

func (h HealthResource) checkHealth(request *restful.Request, response *restful.Response) {
	util.WriteSuccessEntity(response, types.Message{Message: "OK"})
}
