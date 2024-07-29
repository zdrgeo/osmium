package view

type DeleteViewHandler struct {
	repository ViewRepository
}

func NewDeleteViewHandler(repository ViewRepository) *DeleteViewHandler {
	return &DeleteViewHandler{repository: repository}
}

func (handler *DeleteViewHandler) DeleteView(analysisName, name string) {
	handler.repository.Remove(analysisName, name)
}
