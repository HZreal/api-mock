create table tb_api
(
    id           int auto_increment
        primary key,
    name         varchar(16)  null comment '名称',
    uri          varchar(128) null comment '路径',
    args         varchar(128) null comment '路径参数',
    method       varchar(8)   null,
    params       json         null comment '参数',
    content_type varchar(64)  null,
    uri_args     varchar(256) null comment '带 args 路径',
    body         varchar(512) null
)
    comment '接口表';
