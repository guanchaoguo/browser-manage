/*
SQLyog  v12.2.6 (64 bit)
MySQL - 5.7.18-log : Database - browser
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`browser` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `browser`;

/*Table structure for table `access_statis_day` */

DROP TABLE IF EXISTS `access_statis_day`;

CREATE TABLE `access_statis_day` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `domain` varchar(45) NOT NULL COMMENT '域名',
  `ymd` date NOT NULL COMMENT '当天',
  `num` int(11) DEFAULT '0' COMMENT '访问量',
  PRIMARY KEY (`id`),
  KEY `idx1` (`ymd`,`domain`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='访问次数(天)统计表';

/*Table structure for table `access_statis_hour` */

DROP TABLE IF EXISTS `access_statis_hour`;

CREATE TABLE `access_statis_hour` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `domain` varchar(45) NOT NULL COMMENT '域名',
  `ymd` date NOT NULL COMMENT '当天',
  `hour0` int(11) DEFAULT '0' COMMENT '00:00~00:59访问量',
  `hour1` int(11) DEFAULT '0' COMMENT '01:00~01:59访问量',
  `hour2` int(11) DEFAULT '0' COMMENT '01:00~01:59访问量',
  `hour3` int(11) DEFAULT '0',
  `hour4` int(11) DEFAULT '0',
  `hour5` int(11) DEFAULT '0',
  `hour6` int(11) DEFAULT '0',
  `hour7` int(11) DEFAULT '0',
  `hour8` int(11) DEFAULT '0',
  `hour9` int(11) DEFAULT '0',
  `hour10` int(11) DEFAULT '0',
  `hour11` int(11) DEFAULT '0',
  `hour12` int(11) DEFAULT '0',
  `hour13` int(11) DEFAULT '0',
  `hour14` int(11) DEFAULT '0',
  `hour15` int(11) DEFAULT '0',
  `hour16` int(11) DEFAULT '0',
  `hour17` int(11) DEFAULT '0',
  `hour18` int(11) DEFAULT '0',
  `hour19` int(11) DEFAULT '0',
  `hour20` int(11) DEFAULT '0',
  `hour21` int(11) DEFAULT '0',
  `hour22` int(11) DEFAULT '0',
  `hour23` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx1` (`ymd`,`domain`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='访问次数(小时)统计表';

/*Table structure for table `admin_user` */

DROP TABLE IF EXISTS `admin_user`;

CREATE TABLE `admin_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '后台用户登录名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '用户密码（加密后的密码，算法为：md5(sha1(value + 盐值))）',
  `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '8-20位密码加密盐值',
  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '用户真实姓名',
  `last_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '账号最后修改时间',
  `last_login_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `last_login_ip` char(15) DEFAULT '' COMMENT '最后登录IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COMMENT='后台管理员账户表';

/*Table structure for table `area` */

DROP TABLE IF EXISTS `area`;

CREATE TABLE `area` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '区域ID',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父区域ID',
  `name` varchar(45) NOT NULL COMMENT '名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='地区信息表';

/*Table structure for table `area_statis` */

DROP TABLE IF EXISTS `area_statis`;

CREATE TABLE `area_statis` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `area_id` int(10) unsigned NOT NULL COMMENT '区域ID',
  `area_parent_id` int(11) NOT NULL COMMENT '父区域ID',
  `ymd` date NOT NULL COMMENT '当天',
  `num` int(11) DEFAULT '0' COMMENT '数量',
  PRIMARY KEY (`id`),
  KEY `idx1` (`ymd`,`area_id`),
  KEY `idx2` (`ymd`,`area_parent_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='地区分布统计表';

/*Table structure for table `client_ver` */

DROP TABLE IF EXISTS `client_ver`;

CREATE TABLE `client_ver` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ver` varchar(15) DEFAULT NULL COMMENT '版本号',
  `desc` text COMMENT '更新描述',
  `url` varchar(255) DEFAULT NULL COMMENT '安装包下载url',
  `md5` varchar(32) DEFAULT NULL COMMENT '安装包MD5值',
  `size` int(11) DEFAULT NULL COMMENT '安装包大小',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='客户端版本信息表';

/*Table structure for table `file_upload` */

DROP TABLE IF EXISTS `file_upload`;

CREATE TABLE `file_upload` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `file_name` varchar(255) NOT NULL DEFAULT '' COMMENT '上传的文件路径和名称',
  `succeed_number` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '导入成功的数量',
  `failure_number` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '导入失败的次数',
  `create_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '导入的时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8 COMMENT='数据导入记录表';

/*Table structure for table `installed_statis` */

DROP TABLE IF EXISTS `installed_statis`;

CREATE TABLE `installed_statis` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mac` varchar(15) NOT NULL COMMENT 'MAC地址',
  `is_uninstall` tinyint(4) DEFAULT '0' COMMENT '是否卸载',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx1` (`mac`),
  KEY `idx2` (`create_time`,`is_uninstall`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='装机量统计表';

/*Table structure for table `menus` */

DROP TABLE IF EXISTS `menus`;

CREATE TABLE `menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL COMMENT '所属父级菜单ID,0为顶级菜单',
  `title_cn` varchar(45) DEFAULT NULL COMMENT '名称',
  `title_en` varchar(45) DEFAULT '' COMMENT '英文菜单名称',
  `class` tinyint(1) DEFAULT '0' COMMENT '菜单类型分类，0为分类菜单，1为页面菜单，默认为分类菜单',
  `desc` varchar(255) DEFAULT NULL COMMENT '菜单描述',
  `link_url` varchar(255) DEFAULT NULL COMMENT '链接地址',
  `icon` varchar(255) DEFAULT '' COMMENT '菜单图标URL',
  `state` tinyint(3) DEFAULT '1' COMMENT '状态：1启用，0禁用',
  `sort_id` int(11) DEFAULT '1' COMMENT '排序字段',
  `menu_code` varchar(45) DEFAULT NULL COMMENT '菜单字符串',
  `update_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=166 DEFAULT CHARSET=utf8 COMMENT='后台菜单表';

/*Table structure for table `proxy_info` */

DROP TABLE IF EXISTS `proxy_info`;

CREATE TABLE `proxy_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `addr` varchar(45) NOT NULL DEFAULT '' COMMENT '地址',
  `port` int(11) NOT NULL DEFAULT '0' COMMENT '端口',
  `user_name` varchar(45) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(45) NOT NULL DEFAULT '' COMMENT '密码',
  `last_update_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '记录最后修改时间',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=21 DEFAULT CHARSET=utf8 COMMENT='代理服务器信息表';

/*Table structure for table `run_statis` */

DROP TABLE IF EXISTS `run_statis`;

CREATE TABLE `run_statis` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ymd` date NOT NULL COMMENT '当天',
  `lt_1m` int(11) DEFAULT '0' COMMENT '<1m用户数量',
  `lt_20m` int(11) DEFAULT '0' COMMENT '1m~20m用户数量',
  `lt_40m` int(11) DEFAULT '0' COMMENT '20m~40m用户数量',
  `lt_60m` int(11) DEFAULT '0',
  `lt_2h` int(11) DEFAULT '0',
  `lt_3h` int(11) DEFAULT '0',
  `lt_6h` int(11) DEFAULT '0',
  `lt_10h` int(11) DEFAULT '0',
  `lt_20h` int(11) DEFAULT '0',
  `gt_20h` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx1` (`ymd`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='运行时长统计表';

/*Table structure for table `start_statis_day` */

DROP TABLE IF EXISTS `start_statis_day`;

CREATE TABLE `start_statis_day` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ymd` date NOT NULL COMMENT '当天',
  `num` int(11) DEFAULT '0' COMMENT '当天启动量',
  PRIMARY KEY (`id`),
  KEY `idx1` (`ymd`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='启动次数(天)统计表';

/*Table structure for table `start_statis_hour` */

DROP TABLE IF EXISTS `start_statis_hour`;

CREATE TABLE `start_statis_hour` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ymd` date NOT NULL COMMENT '当天',
  `hour0` int(11) DEFAULT '0' COMMENT '00:00-00:59启动量',
  `hour1` int(11) DEFAULT '0' COMMENT '01:00-01:59启动量',
  `hour2` int(11) DEFAULT '0' COMMENT '02:00-02:59启动量',
  `hour3` int(11) DEFAULT '0',
  `hour4` int(11) DEFAULT '0',
  `hour5` int(11) DEFAULT '0',
  `hour6` int(11) DEFAULT '0',
  `hour7` int(11) DEFAULT '0',
  `hour8` int(11) DEFAULT '0',
  `hour9` int(11) DEFAULT '0',
  `hour10` int(11) DEFAULT '0',
  `hour11` int(11) DEFAULT '0',
  `hour12` int(11) DEFAULT '0',
  `hour13` int(11) DEFAULT '0',
  `hour14` int(11) DEFAULT '0',
  `hour15` int(11) DEFAULT '0',
  `hour16` int(11) DEFAULT '0',
  `hour17` int(11) DEFAULT '0',
  `hour18` int(11) DEFAULT '0',
  `hour19` int(11) DEFAULT '0',
  `hour20` int(11) DEFAULT '0',
  `hour21` int(11) DEFAULT '0',
  `hour22` int(11) DEFAULT '0',
  `hour23` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx1` (`ymd`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='启动次数(小时)统计表';

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mac` varchar(15) NOT NULL COMMENT 'MAC地址',
  `token` varchar(64) DEFAULT NULL COMMENT '令牌',
  `aeskey` varchar(64) DEFAULT NULL COMMENT 'AES密码',
  `expiry_time` bigint(11) DEFAULT NULL COMMENT '令牌超时时间',
  `pre_token` varchar(64) DEFAULT NULL COMMENT '前一个令牌',
  `pre_aeskey` varchar(64) DEFAULT NULL COMMENT '前一个AES密码',
  `is_uninstall` tinyint(4) DEFAULT '0' COMMENT '是否卸载浏览器',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx1` (`mac`)
) ENGINE=MyISAM AUTO_INCREMENT=100015 DEFAULT CHARSET=utf8 COMMENT='浏览器用户信息表';

/*Table structure for table `whitelist` */

DROP TABLE IF EXISTS `whitelist`;

CREATE TABLE `whitelist` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `domain` varchar(45) NOT NULL DEFAULT '' COMMENT '域名',
  `hall_name` varchar(45) NOT NULL DEFAULT '' COMMENT '所属人名称',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '启用状态，2为锁定，1为启用状态，默认为1',
  `channel` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '隧道，1：走IP(默认)，2：走代理',
  `info_state` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '记录状态，1为正常，2为更新，3为删除',
  `ips` text COMMENT 'IP字段，用英文半角分号隔开',
  `lock_remark` varchar(255) NOT NULL DEFAULT '' COMMENT '锁定原因',
  `create_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `last_update_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后修改时间',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx1` (`domain`)
) ENGINE=MyISAM AUTO_INCREMENT=449 DEFAULT CHARSET=utf8 COMMENT=' 域名白名单信息表';

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
