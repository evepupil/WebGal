package models
import "WebGal/dao"

type Todo struct {
	ID int `json:"id"`  //和前端做交互 使用的是json数据
	Title string `json:"title"`
	Status bool	`json:"status"`
}
/*
下面是Todo这个Model的crud操作
*/

func CreateATodo(todo *Todo)(err error)  {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}
func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil{
		return nil, err
	}
	return
}
func GetATodo(id string)(todo *Todo, err error)  {
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil{
		return nil, err
	}
	return
}
func UpdateATodo(todo *Todo)(err error)  {
	err = dao.DB.Save(todo).Error
	return
}
func DeleteTodo(id string)(err error)  {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}

