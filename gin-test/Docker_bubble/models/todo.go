package models

import (
	"bubble/dao"
)

// model
type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 添加数据
func AddTodo(todo *Todo) (err error) {
	tx := dao.Db.Create(todo)
	return tx.Error
}

// 更新数据
func UpdateTodo(id string, todo *Todo) (err error) {
	tx := dao.Db.Model(&Todo{}).Debug().Where("id = ?", id).Save(todo)
	return tx.Error
}

// 查询全部数据
func FindTodoAll() (todoList *[]Todo, err error) {
	tx := dao.Db.Debug().Find(&todoList)
	err = tx.Error
	return
}

// 查询单条数据
func FindTodo(id string) (todo Todo, err error) {
	tx := dao.Db.Debug().First(&todo, id)
	err = tx.Error
	return
}

// 删除指定数据
func DeleteTodo(id string) (err error) {
	tx := dao.Db.Debug().Where("id = ?", id).Delete(&Todo{})
	err = tx.Error
	return
}
