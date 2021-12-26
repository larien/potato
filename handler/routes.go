package handler

import (
	"net/http"

	"github.com/larien/potato/utils/router"
)

var Routes = router.Routes{
	{
		Name:    "GetPotato",
		Path:    "/potato",
		Method:  http.MethodGet,
		Handler: GetPotato,
	},
	{
		Name:    "GetPotatos",
		Path:    "/potatos",
		Method:  http.MethodGet,
		Handler: GetPotatos,
	},
	{
		Name:    "GetPotatoByID",
		Path:    "/potatos/{id}",
		Method:  http.MethodGet,
		Handler: GetPotatoByID,
	},
	{
		Name:    "CreatePotato",
		Path:    "/potatos",
		Method:  http.MethodPost,
		Handler: CreatePotato,
	},
	{
		Name:    "DeletePotato",
		Path:    "/potatos/{id}",
		Method:  http.MethodDelete,
		Handler: DeletePotato,
	},
}
