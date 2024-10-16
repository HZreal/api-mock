create table if not exists tb_api
(
    id                    int auto_increment
    primary key,
    name                  varchar(256)     null comment '名称',
    uri                   varchar(128)     null comment '路径',
    args                  varchar(128)     null comment '路径参数',
    method                varchar(8)       null,
    params                json             null comment '参数',
    content_type          varchar(128)     null,
    uri_args              varchar(256)     null comment '带 args 路径',
    request_body          text             null,
    body_type             tinyint unsigned null,
    response_content_type varchar(64)      null comment '响应类型'
    )
    comment '接口表';

create table if not exists tb_rule
(
    id    int auto_increment
    primary key,
    name  varchar(16)                  null comment '名称',
    rules json                         null comment '规则对',
    type  tinyint unsigned default '0' not null comment '规则类型: 0 - 基本类型；1 - 业务类型'
    );

create table if not exists tb_task
(
    id            int auto_increment
    primary key,
    name          varchar(16)                  null comment '名称',
    total_count   int unsigned     default '0' not null comment '总数量',
    success_count int unsigned     default '0' not null comment '成功数量',
    current       int unsigned                 null comment '当前',
    settings      json                         null comment '设置',
    progress      tinyint unsigned             null comment '进度 0 - 100',
    status        tinyint unsigned default '0' not null comment '状态：0 - 待处理；1 - 进行中；2 - 已完成；3 - 失败'
    )
    comment '任务表';

create table if not exists tb_task_record
(
    id       int auto_increment
    primary key,
    task_id  int                          not null comment '任务 id',
    api_id   int                          not null comment '接口 id',
    params   json                         null comment '参数',
    response json                         null comment '结果',
    status   tinyint unsigned default '0' not null comment '状态：0 - 待处理；1 - 进行中；2 - 已完成；3 - 失败',
    cost     mediumint unsigned           null comment '时间消耗 ms'
);

