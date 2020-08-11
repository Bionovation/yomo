package routes

import (
	"github.com/kataras/iris/v12/context"
	"yomo/server/config"
	"yomo/utils"
)

// get workspace list
func GetWorkspaceList(ctx context.Context)  {
	res := utils.NewRes(ctx)
	conf := config.Get()
	res.DoneData(conf.Workspaces)
}