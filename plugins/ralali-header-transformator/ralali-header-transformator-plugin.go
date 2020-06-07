package main

import (
	"context"
	"fmt"
	"net/http"
)

// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
var ClientRegisterer = registerer("ralaliHeaderTransformatorPlugin")

type registerer string

const pluginName = "ralaliHeaderTransformatorPlugin"

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(pluginName, r.registerClients)
}

func (r registerer) registerClients(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("changing header request")
		req.Header.Set("Coba-Set", "Value-Set")
		req.Header.Add("Coba-Add", "Value-Add")
		w.Header().Add("Coba-H-Add", "Value-H-Add")
		w.Header().Set("Coba-H-Set", "Value-H-Set")
	}), nil
}

func init() {
	fmt.Println("loading ralali-header-transformator-plugin..")
}

func main() {}
