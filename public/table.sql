create table tb_api
(
    id           int auto_increment
        primary key,
    name         varchar(16)  null,
    uri          varchar(64)  null,
    args         varchar(128) null,
    method       varchar(8)   null,
    params       json         null,
    content_type varchar(64)  null,
    uri_args     varchar(64)  null
)
    comment '接口表';

