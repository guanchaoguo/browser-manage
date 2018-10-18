package models

type Menus struct {
	Id         int32  `json:"id" xorm:"autoincr pk id"`
	ParentId   int32  `json:"parent_id"`
	TitleCn    string `json:"title_cn"`
	TitleEn    string `json:"title_en"`
	Class      int32  `json:"class"`
	Desc       string `json:"desc"`
	LinkUrl    string `json:"link_url"`
	Icon       string `json:"icon"`
	State      int32  `json:"state"`
	SortId     int32  `json:"sort_id"`
	MenuCode   string `json:"menu_code"`
	UpdateDate string `json:"update_date"`
}

func (Menus) TableName() string {
	return "menus"
}

/*
	根据ID删除操作
*/
func (Menus) DeleteById(id int32, menus *Menus) bool {
	aff, err := engine.Id(id).Delete(menus)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	根据条件获取单条数据
*/
func (Menus) GetOne(menus *Menus) *Menus {
	has, err := engine.Get(menus)
	if !has || err != nil {
		return nil
	}
	return menus
}

/*
	添加单条数据
*/
func (Menus) Add(menus *Menus) int32 {
	aff, err := engine.Insert(menus)
	if aff == 0 || err != nil {
		return 0
	}
	return int32(menus.Id)
}

/*
	根据ID进行修改操作
*/
func (Menus) UpdateById(id int32, menus *Menus) bool {
	aff, err := engine.Id(id).Update(menus)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	获取列表
*/
func (Menus) List(list []Menus, condition *Menus) []Menus {
	err := engine.OrderBy("sort_id").Find(&list, condition)
	if err != nil {
		return nil
	}
	return list
}
