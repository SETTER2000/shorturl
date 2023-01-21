package handlers

import (
	"github.com/go-chi/chi/v5"
)

// Handler единая точка входа для приложения
// этот интерфейс стоит над всеми пакетами системы (user, catalog и т.п.),
// через него осуществляется связь во всем приложении и с внешним миром
// по сути это клей пакетов, маршрутизатор запросов
type Handler interface {
	Register(router *chi.Mux)
}
