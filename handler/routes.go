package handler

import (
	"net/http"

	"github.com/larien/potato/utils/router"
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
	},
	{
		Name:    "UpdatePotato",
		Path:    "/potatoes/{id}",
		Method:  http.MethodPut,
		Handler: UpdatePotato,
	},
	{
		Name:    "DeletePotato",
		Path:    "/potatoes/{id}",
		Method:  http.MethodDelete,
		Handler: DeletePotato,
	},
}
