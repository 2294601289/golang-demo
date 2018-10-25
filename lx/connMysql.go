package main
import (
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "fmt"
//    "log"
)
func main(){

    //创建连接
    db, err := sql.Open("mysql", "mytestuser:mysql@tcp(127.0.0.1:3306)/mytestdata");
    check(err);

    //
    var cpt caoptest
    cpt.userno="0012"
    cpt.username="柯南12"
    cpt.password="0"
    cpt.birthday="1993-05-06"
    cpt.isvalid="1"
    cpt.spellno="KN"
    //增加
    //insertcp(db,cpt);
    //修改
    //updatecp(db,cpt);
    //删除
    deletecp(db,"0012");
    //查询
    selectcp(db);
    //按照ID查询
    //selectcpOne(db,"0011");

    //关闭连接
    db.Close();

}
//错误检查方法
func check(err error){
	if err != nil{
		fmt.Println("失败，错误：",err);
		panic(err);
    }
//    fmt.Println("正常！");
}
//结构体：参考数据库表定义变量
type caoptest struct{
    userno string
    username string
    password string
    birthday string
    isvalid string
    spellno string
}
//查询
func selectcp(db *sql.DB){
	rows, err :=db.Query("select * from caoptest");
    check(err);
    for rows.Next() {
        var cp caoptest;
        err = rows.Scan(&cp.userno,&cp.username,&cp.password,&cp.birthday,&cp.isvalid,&cp.spellno);
        check(err);
        fmt.Println(cp);
     }
     rows.Close();
}
//按照ID查询
func selectcpOne(db *sql.DB,userno string){
	rows, err :=db.Query("select * from caoptest where userno=?",userno);
    check(err);
    for rows.Next() {
        var cp caoptest;
        err = rows.Scan(&cp.userno,&cp.username,&cp.password,&cp.birthday,&cp.isvalid,&cp.spellno);
        check(err);
        fmt.Println(cp);
     }
     rows.Close();
}
//增加
func insertcp(db *sql.DB,cp caoptest){
    stmt,err:=db.Prepare("INSERT INTO caoptest(userno,username,password,birthday,isvalid,spellno) VALUES(?,?,?,STR_TO_DATE(?,'%Y-%m-%d'),?,?)");
    check(err);
    res,err:=stmt.Exec(cp.userno,cp.username,cp.password,cp.birthday,cp.isvalid,cp.spellno);
    check(err);
    id,err :=res.LastInsertId();
    check(err);
    fmt.Println(id);
    stmt.Close();
}
//删除
func deletecp(db *sql.DB , userno string){
    stmt,err:=db.Prepare("DELETE FROM caoptest where userno=?");
    check(err);
     res,err:=stmt.Exec(userno);
    check(err);
    affect,err :=res.RowsAffected();
    check(err);
    fmt.Println(affect);
    stmt.Close();
}
//修改
func updatecp(db *sql.DB , cp caoptest){
    stmt,err:=db.Prepare("UPDATE caoptest set username=?,password=?,birthday=STR_TO_DATE(?,'%Y-%m-%d'),isvalid=?,spellno=? where userno=?");
    check(err);
    res,err:=stmt.Exec(cp.username,cp.password,cp.birthday,cp.isvalid,cp.spellno,cp.userno);
    check(err);
    affect,err :=res.RowsAffected();
    check(err);
    fmt.Println(affect);
    stmt.Close();
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