package api

type (
	StorePostRequest struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)
