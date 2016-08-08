package routes

import "github.com/golang/glog"
import "github.com/kataras/iris"

import "github.com/sizethree/meritoss.api/api"
import "github.com/sizethree/meritoss.api/api/dal"


// CreateClient
//
// request handler for POST /clients
// 
// will attempt to load in json data as a `ClientFacade` and use the `dal.CreateClient`
// function to persist that information to the clients table
func CreateClient(context *iris.Context) {
	runtime, ok := context.Get("runtime").(*api.Runtime)

	if !ok {
		glog.Error("bad runtime")
		context.Panic()
		context.StopExecution()
		return
	}

	// prepare our facade for iris to load into
	var target dal.ClientFacade

	if e := context.ReadJSON(&target); e != nil {
		runtime.Errors = append(runtime.Errors, e)
		return
	}

	// attempt to create the new client
	client, e := dal.CreateClient(&runtime.DB, &target);

	if e != nil {
		runtime.Errors = append(runtime.Errors, e)
		return
	}

	runtime.Results = append(runtime.Results, client)
	runtime.Meta.Total = 1

	glog.Infof("created client %d\n", client.ID)
	context.Next()
}