package handler

import (
	"net/http"

	"github.com/osamikoyo/portfiler/internal/service"
	"github.com/osamikoyo/portfiler/pkg/loger"
)

type Handler struct{
	logger loger.Logger
	service *service.PortfileService
}

func New() (Handler, error) {
	service, err := service.New()
	return Handler{service: &service}, err
}

type handler func (w http.ResponseWriter, r *http.Request) error

func (h Handler) RegisterRouter(mux *http.ServeMux) {
	h.logger.Info().Msg("Registing routes! :3")

	mux.Handle("/review/add", h.errorRoute(h.addReviewHandler))
	mux.Handle("/portfolio/add", h.errorRoute(h.addPortfolioHandler))
	mux.Handle("portfolio/get", h.errorRoute(h.getPortfolioHandler))
}