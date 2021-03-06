package riminder

import "github.com/Xalrandion/go-riminder-api/response"

// source class contains methods to interact with the sources on riminder's api.
type source struct {
	client *clientw
}

func newSource(riminder *Riminder) *source {
	s := new(source)

	s.client = riminder.clientw
	return s
}

// List gets the list of your sources.
func (s *source) List() ([]riminderResponse.SourceListElem, error) {

	resp := riminderResponse.SourceListContainer{}
	err := s.client.Get("sources", map[string]string{}, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Get gets a specific source.
func (s *source) Get(options map[string]interface{}) (riminderResponse.SourceGetElem, error) {

	query := map[string]string{
		"source_id": options["source_id"].(string),
	}

	resp := riminderResponse.SourceGetContainer{}
	err := s.client.Get("source", query, &resp)
	if err != nil {
		return riminderResponse.SourceGetElem{}, err
	}
	return resp.Data, nil
}
