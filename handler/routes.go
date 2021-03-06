package handler

import (
	"net/http"

	"github.com/larien/potato/internal/router"
)

var Routes = router.Routes{
	{
		Name:    "GetPotatoes",
		Path:    "/potatoes",
		Method:  http.MethodGet,
		Handler: GetPotatoes,
	},
	{
		Name:    "GetPotatoByID",
		Path:    "/potatoes/{id}",
		Method:  http.MethodGet,
		Handler: GetPotatoByID,
	},
	{
		Name:    "CreatePotato",
		Path:    "/potatoes",
		Method:  http.MethodPost,
		Handler: CreatePotato,
		IsAdmin: true,
	},
	{
		Name:    "UpdatePotato",
		Path:    "/potatoes/{id}",
		Method:  http.MethodPut,
		Handler: UpdatePotato,
		IsAdmin: true,
	},
	{
		Name:    "DeletePotato",
		Path:    "/potatoes/{id}",
		Method:  http.MethodDelete,
		Handler: DeletePotato,
		IsAdmin: true,
	},
}
