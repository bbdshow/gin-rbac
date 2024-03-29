package service

import (
	"github.com/bbdshow/bkit/tests"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"testing"
)

func TestService_ImportActionConfig(t *testing.T) {
	if err := svc.ImportActionConfig(ctx, &model.ImportActionConfigReq{
		AppId:  "",
		Name:   "时区",
		Path:   "/v1/date/update",
		Method: "POST",
	}); err != nil {
		t.Fatal(err)
	}
}

func TestService_CreateMenuConfig(t *testing.T) {
	if err := svc.CreateMenuConfig(ctx, &model.CreateMenuConfigReq{
		AppId:    "",
		Name:     "时间名称设置",
		Memo:     "",
		ParentId: 4,
		Sequence: 1,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestService_UpdateMenuAction(t *testing.T) {
	if err := svc.UpdateMenuConfigAction(ctx, &model.UpdateMenuConfigActionReq{
		MenuId:   1,
		ActionId: []int64{1, 2},
	}); err != nil {
		t.Fatal(err)
	}
}

func TestService_menuTreeDirs(t *testing.T) {
	menus, err := svc.d.FindMenuConfig(ctx, &model.FindMenuConfigReq{
		AppId:    "000000",
		ParentId: -1, // all menus
	})
	if err != nil {
		t.Fatal(err)
	}
	dirs, err := svc.menuTreeDirs(menus)
	if err != nil {
		t.Fatal(err)
	}
	tests.PrintBeautifyJSON(dirs)
}

func TestService_GetMenuActions(t *testing.T) {
	out := &model.GetMenuActionsResp{}
	if err := svc.GetMenuActions(ctx, &model.GetMenuActionsReq{
		MenuId: 2,
	}, out); err != nil {
		t.Fatal(err)
	}
	tests.PrintBeautifyJSON(out)
}
