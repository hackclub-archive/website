package hackedu

import (
	"net/http"
	"time"

	"github.com/crhym3/go-endpoints/endpoints"

	"appengine/datastore"
)

type School struct {
	Key       *datastore.Key `json:"id" datastore:"-"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	Name      string         `json:"name,omitempty"`
	Location  Location       `json:"location,omitempty"`
}

type Location struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

// SchoolsList is a response type of SchoolService.List method
type SchoolsList struct {
	Items []*School `json:"schools"`
}

// Request type for SchoolService.List
type SchoolsListReq struct {
	Limit int `json:"limit" endpoints="d=10"`
}

// SchoolService can create, list, and delete schools from the datastore
type SchoolService struct {
}

func (ss *SchoolService) List(r *http.Request, req *SchoolsListReq, resp *SchoolsList) error {
	if req.Limit <= 0 {
		req.Limit = 10
	}

	c := endpoints.NewContext(r)
	q := datastore.NewQuery("school").Order("-CreatedAt").Limit(req.Limit)
	schools := make([]*School, 0, req.Limit)
	keys, err := q.GetAll(c, &schools)
	if err != nil {
		return err
	}

	for i, k := range keys {
		schools[i].Key = k
	}
	resp.Items = schools
	return nil
}
