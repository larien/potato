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
}
