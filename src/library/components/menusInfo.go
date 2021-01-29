package components

import (
	"models"
	"strconv"
	"strings"
)

/*
func UserMenusInfo(menus []MenusAuth)(rootNode MenusAuth) {

	//revMenus := reverse(menus)

	menusLen := len(menus)

	for i := 0; i < menusLen; i++ {
		for j := 0; j < menusLen; j++ {
			if menus[j].Id == menus[i].ParentId {
				menus[j].Child = append(menus[j].Child, menus[i])
			}
		}
	}

	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId == 0 {
			rootNode = menus[i]
		}
	}
	return rootNode
}

func reverse(m []MenusAuth) []MenusAuth {
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}
	return m
}
*/


const(
	IconHome = "fa-home"
	IconProject = "el-icon-menu"
	IconDeploy = "fa-table"
	IconResource = "el-icon-menu"
	IconCrontab = "el-icon-menu"
	IconFast = "el-icon-share"
	IconUser = "el-icon-setting"
)



//
func FindRootMenu(menus []MenusAuth)(rootNode MenusAuth){

	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId == 0 {
			rootNode = menus[i]
		}
	}
	return rootNode
}


// 递归
func FindChild(menus *[]MenusAuth, m *MenusAuth)(Node *MenusAuth){
	menusLen := len(*menus)
	if menusLen == 0 {
		return nil
	}
	//m.Title = m.
	m.Icon = MenuIcon(m)
	for i := 0; i < menusLen; i++ {
		if m.Id == (*menus)[i].ParentId {
			child := FindChild(menus,&(*menus)[i])
			if child == nil{
				return
			}
			m.Child = append(m.Child,(*menus)[i])
		}
	}
	respMenus := make([]MenusAuth,0)
	copy(respMenus,m.Child)
	return m
}


func FindDeployList(menus *[]MenusAuth, u models.User) {
	for i := 0; i < len(*menus); i++ {
		if  strings.Contains((*menus)[i].Path,"/deploy/list/") {
			(*menus)[i].Path += strconv.Itoa(u.Id)
		}
	}
}

func MenuIcon(m *MenusAuth) (icon string){
	switch m.Path {
	case "/home":
		icon = IconHome
	case "/project":
		icon = IconProject
	case "/deploy":
		icon = IconDeploy
	case "/resource":
		icon = IconResource
	case "/crontab":
		icon = IconCrontab
	case "/fast":
		icon = IconFast
	case "/user":
		icon = IconUser
	default:
		icon = ""
	}
	return
}