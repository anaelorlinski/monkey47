package monkey47

import (
	"net/http"

	health "github.com/hellofresh/health-go/v5"
)

func HealthProbe() {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    Component(),
		Version: Version(),
	}))
	http.Handle("/status", h.Handler()) // follow monkey convention
	http.ListenAndServe(":3000", nil)
}
