package menu

type NodeSon struct {
	NodeSonText string
	NodeSonPriv string
}

type Node struct {
	NodeText  string
	NodeUrl   string
	NodePriv  string
	PrivChild []NodeSon
}

type Menu struct {
	MenuText string
	MenuPriv string
	MenuIcon string
	Nodes    []Node
}

var MenuList []Menu

func GetMenu() []Menu {
	MenuList = []Menu{
		{
			MenuText: "设置",
			MenuPriv: "setting",
			MenuIcon: "fa fa-cog",
			Nodes: []Node{
				{
					NodeText: "管理员管理",
					NodeUrl:  "/admin/setting/adminuser/index",
					NodePriv: "/admin/setting/adminuser/index",
					PrivChild: []NodeSon{
						{
							NodeSonText: "添加管理员",
							NodeSonPriv: "/admin/setting/adminuser/add",
						},
						{
							NodeSonText: "编辑管理员",
							NodeSonPriv: "/admin/setting/adminuser/edit",
						},
						{
							NodeSonText: "保存管理员",
							NodeSonPriv: "/admin/setting/adminuser/save",
						},
					},
				},
				{
					NodeText: "角色管理",
					NodeUrl:  "/admin/setting/admingroup/index",
					NodePriv: "/admin/setting/admingroup/index",
					PrivChild: []NodeSon{
						{
							NodeSonText: "添加角色",
							NodeSonPriv: "/admin/setting/admingroup/add",
						},
						{
							NodeSonText: "编辑角色",
							NodeSonPriv: "/admin/setting/admingroup/edit",
						},
						{
							NodeSonText: "保存角色",
							NodeSonPriv: "/admin/setting/admingroup/save",
						},
					},
				},
				{
					NodeText: "系统日志",
					NodeUrl:  "/admin/setting/system/index",
					NodePriv: "/admin/setting/system/index",
					PrivChild: []NodeSon{
						{
							NodeSonText: "日志列表",
							NodeSonPriv: "/admin/setting/system/index",
						},
						{
							NodeSonText: "获取目录",
							NodeSonPriv: "/admin/setting/system/getdir",
						},
						{
							NodeSonText: "读取日志",
							NodeSonPriv: "/admin/setting/system/view",
						},
					},
				},
			},
		},
		{
			MenuText: "示例",
			MenuPriv: "demo",
			MenuIcon: "fa fa-cog",
			Nodes: []Node{
				{
					NodeText: "附件管理",
					NodeUrl:  "/admin/demo/show",
					NodePriv: "/admin/demo/show",
					PrivChild: []NodeSon{
						{
							NodeSonText: "上传文件",
							NodeSonPriv: "/admin/demo/upload",
						},
					},
				},
			},
		},
	}
	return MenuList
}
