package riminder

import "github.com/Xalrandion/go-riminder-api/riminder/response"

type filter struct {
	client *clientw
}

func newfilter(riminder *Riminder) *filter {
	s := new(filter)

	s.client = riminder.clientw
	return s
}

func (s *filter) List() ([]response.FilterListElem, error) {

	resp := response.FilterListContainer{}
	tmpResp, err := s.client.Get("filters", map[string]string{}, resp)
	if tmpResp == nil || err != nil {
		return nil, err
	}
	return tmpResp.(*response.FilterListContainer).Data, nil
}

func (s *filter) Get(options map[string]interface{}) (response.FilterGetElem, error) {

	query := map[string]string{}
	AddIfNotEmptyStrMap(&query, options, "filter_id")
	AddIfNotEmptyStrMap(&query, options, "filter_reference")

	resp := response.FilterGetContainer{}
	tmpResp, err := s.client.Get("filter", query, resp)
	if tmpResp == nil || err != nil {
		return response.FilterGetElem{}, err
	}
	return tmpResp.(*response.FilterGetContainer).Data, nil
}
