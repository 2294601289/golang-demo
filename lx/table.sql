DROP TABLE caoptest cascade;

CREATE TABLE caoptest(
	userno varchar(20) not null COMMENT '用户代码',
	username varchar(15) not null COMMENT '用户名',
	password varchar(20) not null COMMENT '用户密码',
	birthday datetime COMMENT '出生日期',
	isvalid char(1) not null default '1' COMMENT '是否有效',
	spellno varchar(15) COMMENT '简拼'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

alter table caoptest add primary key PK_caoptest (userno);