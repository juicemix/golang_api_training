package common

type Error struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

type ErrorReturn struct {
	Err Error `json:"response"`
}
