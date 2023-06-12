package common

type Object struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var WorldObjects = map[int]*Object{
	1: {
		Id:   1,
		Name: "餐厅",
	},
	2: {
		Id:   2,
		Name: "花园",
	},
	3: {
		Id:   3,
		Name: "广场",
	},
	4: {
		Id:   4,
		Name: "便利店",
	},
	5: {
		Id:   5,
		Name: "大学",
	},
	6: {
		Id:   6,
		Name: "咖啡店",
	},
	7: {
		Id:   7,
		Name: "篮球场",
	},
	8: {
		Id:   8,
		Name: "育儿园",
	},
	9: {
		Id:   9,
		Name: "办公楼",
	},
	93: {
		Id:   93,
		Name: "林婷的家",
	},
	94: {
		Id:   94,
		Name: "李明的家",
	},
	95: {
		Id:   95,
		Name: "陈刚的家",
	},
	96: {
		Id:   96,
		Name: "刘洋的家",
	},
	97: {
		Id:   97,
		Name: "黄浩的家",
	},
	98: {
		Id:   98,
		Name: "黄伟的家",
	},
	99: {
		Id:   99,
		Name: "张伟的家",
	},
}
