package handlers

type Head struct {
}

func (h *Head) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
func (h *Head) PathInfo() (string, string) {
	return "GET", "/{repo}/HEAD"
}
