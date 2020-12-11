package response

type JsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (jsonRes *JsonResponse) returnJsonResponse(isSuccess bool, message string) {

}
