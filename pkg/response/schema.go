package response

type ErrorResponse struct {
	Code    int      `json:"code" example:"404"`
	Status  string   `json:"status" example:"Not Found"`
	Data    struct{} `json:"data"`
	Message string   `json:"message" example:"Data tidak ditemukan"`
}

type ValidationErrorResponse struct {
	Code    int               `json:"code" example:"400"`
	Status  string            `json:"status" example:"Bad Request"`
	Data    struct{}          `json:"data"`
	Message map[string]string `json:"message" example:"{\"title\":\"Kolom title wajib diisi\"}"`
}

type SuccessResponse struct {
	Code    int      `json:"code" example:"200"`
	Status  string   `json:"status" example:"OK"`
	Data    struct{} `json:"data"`
	Message string   `json:"message" example:"Data berhasil diambil"`
}
