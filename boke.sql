-- 创建库
create database if not exists blog;

-- 选择数据库
use blog;

-- 文章分类表
create table if not exists `article_class`(
  `id` int unsigned auto_increment primary key,
  `name` varchar(20) unique NOT NULL COMMENT '分类名',
  `sort` tinyint(2) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0-隐藏 1-显示 2-header'
)engine=innodb default charset=utf8 comment '文章分类表';

-- 文章表
CREATE TABLE IF NOT EXISTS `article` (
  `article_id` int unsigned NOT NULL AUTO_INCREMENT,
  `class_name` varchar(20) comment '分类名',
  `headline` varchar(45) DEFAULT '' comment '标题',
  `summary` varchar(255) DEFAULT '' comment '摘要',
  `content` text DEFAULT '',
  `created_at` int DEFAULT 0,
  `updated_at` int DEFAULT 0,
  `comm` smallint(5) unsigned NOT NULL DEFAULT '0' comment '评论数量',
  `uv` int DEFAULT 0 COMMENT '访客量',
  `pv` int DEFAULT 0 COMMENT '点击量',
  `status` tinyint(1) DEFAULT 0 COMMENT '0-隐藏 1-显示 2-未提交',
  PRIMARY KEY (`article_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='文章表';

--
-- 表的结构 `comment`
--
CREATE TABLE IF NOT EXISTS `article_comment` (
  `comment_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned,
  `nick` varchar(20) COLLATE utf8_unicode_ci DEFAULT '',
  `email` varchar(100) comment '邮箱',
  `home_link` varchar(100) DEFAULT '' comment '用户主页',
  `content` varchar(1000) COLLATE utf8_unicode_ci DEFAULT '',
  `superior` varchar(20) DEFAULT ''  COLLATE utf8_unicode_ci comment '父级(回复)昵称',
  `ip` varchar(20) unsigned DEFAULT '0',
  `created_at` int(10) unsigned DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '0-隐藏 1-显示',
  PRIMARY KEY (`comment_id`),
  KEY `article_id` (`article_id`),
  KEY `ip` (`ip`),
  KEY `nick` (`nick`),
  KEY `superior` (`superior`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;
--
-- 邮件发送表
--
CREATE TABLE IF NOT EXISTS `article_comment_send` (
  `comment_email_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `comment_id` int(10) NOT NULL comment '评论id',
  `email` char(50) NOT NULL comment '发送邮箱',
  `content` char(50) NOT NULL comment '发送内容',
  `state` tinyint(1) DEFAULT 0 comment '0-未发送,1-发送成功,2-发送失败',
  `error_num` int(1) DEFAULT 0 comment '失败次数',
  `error_msg` char(100) DEFAULT "" comment '失败原因',
  `start_time` int(10) unsigned DEFAULT '0' comment '发送时间',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`comment_email_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;

