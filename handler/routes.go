package handler

import (
	"net/http"

	"github.com/larien/potato/utils/router"
)

var Routes = router.Routes{
	{
		Pattern: "/potato",
		Version: "v1",
		Method:  http.MethodGet,
		Handler: V1GetPotato,
	},
}
