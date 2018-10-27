package main

// go get -u github.com/iris-contrib/middleware/...

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"database/sql"
    _"github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
)

func main() {
	//创建连接
	db, err := sql.Open("mysql", "mytestuser:mysql@tcp(127.0.0.1:3306)/mytestdata");
	check(err);


	//--
	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	v1 := app.Party("/v1", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
	{
		v1.Get("/query", func(ctx iris.Context) {
			var cpt string 
			cpt = selectcp(db);
			ctx.WriteString(cpt)
		})
		v1.Get("/about", func(ctx iris.Context) {
			var cpt string 
			cpt = selectcp(db);
			ctx.WriteString(cpt)
		})
		v1.Post("/send", func(ctx iris.Context) {
			
			ctx.WriteString("updatedsend")
		})
		v1.Put("/send", func(ctx iris.Context) {
			ctx.WriteString("updated")
		})
		v1.Delete("/send", func(ctx iris.Context) {
			ctx.WriteString("deleted")
		})
	}

	app.Run(iris.Addr("localhost:8085"))
}

//错误检查方法
func check(err error){
	if err != nil{
		fmt.Println("失败，错误：",err);
		panic(err);
    }
//    fmt.Println("正常！");
}
//结构体：参考数据库表定义变量(必须大写开头)
type caoptest struct{
    Userno 		string `json:"userno"`
    Username 	string `json:"username"`
    Password 	string `json:"password"`
    Birthday 	string `json:"birthday"`
    Isvalid 	string `json:"isvalid"`
    Spellno 	string `json:"spellno"`
}
type cptjson struct{
    Cptj []caoptest `json:"cptj"`
}
//查询
func selectcp(db *sql.DB) string{
	rows, err :=db.Query("select * from caoptest");
    check(err);
    var cptjn cptjson;
    for rows.Next() {
        var cp caoptest;
        err = rows.Scan(&cp.Userno,&cp.Username,&cp.Password,&cp.Birthday,&cp.Isvalid,&cp.Spellno);
        check(err);
        cptjn.Cptj=append(cptjn.Cptj,caoptest{Userno:cp.Userno,Username:cp.Username,Password:cp.Password,Birthday:cp.Birthday,Isvalid:cp.Isvalid,Spellno:cp.Spellno});
        check(err);
     }
     rows.Close();
     j,err:=json.Marshal(cptjn);
     check(err);
	 fmt.Print(string(j));
	 return string(j);
}
///表
// CREATE TABLE caoptest(
// 	userno varchar(20) not null COMMENT '用户代码',
// 	username varchar(15) not null COMMENT '用户名',
// 	password varchar(20) not null COMMENT '用户密码',
// 	birthday datetime COMMENT '出生日期',
// 	isvalid char(1) not null default '1' COMMENT '是否有效',
// 	spellno varchar(15) COMMENT '简拼'
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';
// alter table caoptest add primary key PK_caoptest (userno);