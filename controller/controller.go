package controller
import (
	"WebGal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
/*
正常流程：  url-> controller -> logic -> model 模型层的crud
*/

func IndexHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CreateATodo(c *gin.Context) {
	// 1.请求中取数据
	var todo models.Todo
	c.BindJSON(&todo)

	//2.存储数据 3.返回响应
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todo)
		/*统一响应格式c.JSON(http.StatusOK, gin.H{
							"code": 2000,
							"msg": "success",
							"data": todo,
		})*/
	}
}
func GetTodoList(c *gin.Context) {
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")  //获取请求参数
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
		return
	}

	// 根据id查询待修改实体
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todo)
	}
}
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error":"无效的id"})
		return
	}
	if err := models.DeleteTodo(id); err!=nil{
		c.JSON(http.StatusOK,gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

