-- 创建数据库
CREATE DATABASE IF NOT EXISTS ihome DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

-- 创建信息用户表
DROP TABLE IF EXISTS `ih_user`;

CREATE TABLE `ih_user` (
`user_id` integer NOT NULL AUTO_INCREMENT COMMENT '用户标号',
`name` varchar(32) NOT NULL default "" COMMENT '用户昵称',
`password_hash` varchar(128) NOT NULL default "" COMMENT '加密的密码',
`mobile` varchar(11) NOT NULL unique COMMENT '手机号',
`real_name` varchar(32) NOT NULL default "" COMMENT '真实姓名',
`id_card` varchar(20) NOT NULL default "" COMMENT '身份证号',
`avatar_url` varchar(128) NOT NULL default "" COMMENT '用户头像',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户信息表';

-- 创建信息城区表
DROP TABLE IF EXISTS `ih_area`;
CREATE TABLE `ih_area` (
`area_id` integer AUTO_INCREMENT NOT NULL COMMENT '区域编号',
`aname` varchar(32) NOT NULL default "" COMMENT '区域名字',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`area_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '城区信息表';

--创建信息房屋表
DROP TABLE IF EXISTS `ih_house`;
CREATE TABLE `ih_house`(
`house_id` integer NOT NULL AUTO_INCREMENT COMMENT '房屋编号',
`title` varchar(64) NOT NULL default "" COMMENT '标题',
`price` integer NOT NULL DEFAULT 0 COMMENT '价格 单位:分',
`address` varchar(512) NOT NULL default "" COMMENT '地址',
`room_count` integer NOT NULL default 1 COMMENT '房间数量',
`acreage` integer NOT NULL default 0 COMMENT '房间面积',
`unit` varchar(32) NOT NULL default "" COMMENT '房间类型,如几室几厅',
`capacity` integer NOT NULL default 1 COMMENT '房间容纳的人数',
`beds` varchar(64) NOT NULL default "" COMMENT '房间床铺配置',
`deposit` integer NOT NULL default 0 COMMENT '房屋押金',
`min_days` integer NOT NULL default 1 COMMENT '最少入住天数',
`max_days` integer NOT NULL default 0 COMMENT '最多入住天数零为不限制',
`order_count` integer NOT NULL default 0 COMMENT '预定完成的订单数',
`index_image_url` varchar(256) NOT NULL default "" COMMENT '房屋的主图片地址',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`house_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '房屋信息表';

-- 创建设备信息表
DROP TABLE IF EXISTS `ih_facility`;
CREATE TABLE `ih_facility`(
`fid` integer NOT NULL AUTO_INCREMENT COMMENT '设置编号',
`name` varchar(32) NOT NULL default "" COMMENT '设备名字',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`fid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '设备信息表';

-- 创建房屋图片表
DROP TABLE IF EXISTS `ih_house_image`;
CREATE TABLE `ih_house_image`(
`house_image_id` integer NOT NULL AUTO_INCREMENT COMMENT '图片',
`url` varchar(256) NOT NULL default "" COMMENT '图片地址',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`house_id` integer NOT NULL,
PRIMARY KEY (`house_image_id`),
FOREIGN KEY(house_id) references ih_house(house_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '房屋图片表';

--创建订单表
DROP TABLE IF EXISTS `ih_order`;
CREATE TABLE `ih_order` (
`order_id` integer NOT NULL AUTO_INCREMENT COMMENT '订单编号',
`user_id`  integer COMMENT '关联的用户外键',
`house_id` integer COMMENT '关联的房屋外键',
`begin_date` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '预订的起始时间',
`end_date` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '预订的结束时间',
`days` integer COMMENT '预订的总天数',
`house_price` integer COMMENT '房屋的单价',
`amount` integer COMMENT '订单的总金额',
`status` Enum("WAIT_ACCEPT","WAIT_PAYMENT","PAID","WAIT_COMMENT","COMPLETE","CANCELED","REJECTED")
default "WAIT_ACCEPT" COMMENT '订单的状态',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`order_id`),
FOREIGN KEY(user_id) references ih_user(user_id),
FOREIGN KEY(house_id) references ih_house(house_id),
KEY `idx_order_state` (`status`) USING BTREE  -- 索引
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '订单信息表';

-- 使用 ENUM 枚举这些情况
-- "WAIT_ACCEPT",  # 待接单,
-- "WAIT_PAYMENT",  # 待支付
-- "PAID",  # 已支付
-- "WAIT_COMMENT",  # 待评价
-- "COMPLETE",  # 已完成
-- "CANCELED",  # 已取消
-- "REJECTED"  # 已拒单