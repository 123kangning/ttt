create table users
(
    id       int primary key auto_increment,
    username varchar(50) not null unique,
    password char(60)    not null
);
create table scenes
(
    id          int primary key auto_increment,
    name        varchar(255) not null,
    description text         not null
);
create table images
(
    id   int primary key auto_increment,
    sid  int          not null comment '景点id',
    name varchar(255) not null comment '图片名称',
    url  varchar(255) not null comment '图片地址'
);
create table journeys
(
    id      int primary key auto_increment,
    uid     int          not null comment '制定行程的用户',
    name    varchar(255) not null comment '行程名称',
    content text         not null comment '行程内容'
);
create table comments
(
    id      int primary key auto_increment,
    sid     int  not null comment '被评论的景点id',
    uid     int  not null comment '发表评论的用户',
    content text not null comment '评论内容'
);
