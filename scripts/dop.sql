/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost
 Source Database       : dop

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : utf-8

 Date: 01/11/2018 11:30:35 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `account`
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '编号',
  `username` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '名称',
  `password` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `email` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '邮件',
  `phone` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '电话号码',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '账号状态',
  `deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_unique` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户表';

-- ----------------------------
--  Records of `account`
-- ----------------------------
BEGIN;
INSERT INTO `account` VALUES ('5b4a990536e947b584cfd13df072d352', 'admin', '7d91f135c305c902730b75f8e4a2c7f2', 'xingshanghe@gmail.com', '18010636836', '0', '0001-01-01 00:00:00'), ('81fe5028e7da410cbbe959f6efb0d9fc', 'test2', 'b7779b7cacb751643768cc5369df6793', 'xing@qq.com', '18010636836', '0', '0001-01-01 00:00:00'), ('ace860fc82ac4fb4b2d4bba420879f43', 'test1', 'b7779b7cacb751643768cc5369df6793', 'test1@qq.com', '13551381506', '0', '0001-01-01 00:00:00'), ('dba3d603cc874f97bd7a1b0e50c0f12c', 'guodegang', 'b7779b7cacb751643768cc5369df6793', 'gdg@qq.com', '13551180088', '0', '0001-01-01 00:00:00');
COMMIT;

-- ----------------------------
--  Table structure for `account_detail`
-- ----------------------------
DROP TABLE IF EXISTS `account_detail`;
CREATE TABLE `account_detail` (
  `id` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '编号',
  `account_id` char(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0' COMMENT '用户id',
  `nickname` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '昵称',
  `gender` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '性别',
  `age` int(4) DEFAULT NULL COMMENT '年龄',
  `address` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '地址',
  `birthday` date DEFAULT NULL,
  `created` int(11) NOT NULL DEFAULT '0',
  `updated` int(11) NOT NULL DEFAULT '0',
  `deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户细节表';

-- ----------------------------
--  Records of `account_detail`
-- ----------------------------
BEGIN;
INSERT INTO `account_detail` VALUES ('7b325709125549aa984813a0634cdfb8', '81fe5028e7da410cbbe959f6efb0d9fc', '测试2', 'male', '18', '测试地址2', '2017-12-27', '1514338524', '1515509589', null), ('a8d0f52884144193bf5f73c6e02dddd5', 'ace860fc82ac4fb4b2d4bba420879f43', '测试1', 'female', '18', '测试地址11', '2017-12-07', '1514310133', '1515509593', '0001-01-01 00:00:00'), ('d12f2417dfad4f29abaf8627e2cbe24e', '5b4a990536e947b584cfd13df072d352', '管理员', 'female', '26', '四川省成都市', '2017-12-27', '1514308308', '1515564061', null), ('ee601913a1ad4eb9a10daa5a55e5ae04', 'dba3d603cc874f97bd7a1b0e50c0f12c', '郭德纲', 'female', '18', '北京德云社', '2017-12-27', '1514358351', '1515509591', null);
COMMIT;

-- ----------------------------
--  Table structure for `menu`
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '编号',
  `title` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '菜单名称',
  `icon` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '字体图标',
  `link` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '菜单地址',
  `parent_id` char(32) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '父id',
  `is_group` int(11) NOT NULL DEFAULT '0' COMMENT '是否分组',
  `is_sub` int(11) NOT NULL DEFAULT '1' COMMENT '是否子菜单',
  `is_side` int(11) NOT NULL DEFAULT '1',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态',
  `sort` int(1) NOT NULL DEFAULT '0',
  `updated` int(11) NOT NULL DEFAULT '0',
  `created` int(11) NOT NULL DEFAULT '0',
  `deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='菜单表';

-- ----------------------------
--  Records of `menu`
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES ('0a6af3dd936c4c25a9835b52ddd5d2d0', '菜单管理', 'icon-safe', '/settings/system/menus', 'afab4e4d69cb4443a907fc856f418dcf', '0', '0', '1', '0', '0', '1515400749', '1515400749', null), ('13d0fb1152054bd38782134eaa771cf8', 'test', 'el-icon-edit', '/', '', '0', '0', '0', '0', '0', '1515506420', '1515506420', null), ('146951887110406d9af92dadd45fa2ae', '个人设置', 'icon-user', '/settings/personal/profile', '7f35a6c6fa92434b98a21512f4a8a658', '0', '0', '1', '0', '8', '1515438199', '1515431578', null), ('27ca5ec0b65546acafe0a28ace8636bd', '个人设置', '', '', '146951887110406d9af92dadd45fa2ae', '1', '1', '1', '0', '0', '1515431685', '1515431685', null), ('345f520608b7481690b7046d77d5b4e0', '管理控制台', 'icon-home2', '/', '', '0', '0', '1', '0', '0', '1515400837', '1515400837', null), ('4f569a92c62e4611b7333f5bd07647ad', '修改密码', 'icon-key', '/settings/personal/password', '27ca5ec0b65546acafe0a28ace8636bd', '0', '0', '1', '0', '0', '1515431726', '1515431726', null), ('5397bf6cf9b04b2c856fa80bbd15134e', '个人资料设置', 'icon-vcard', '/settings/personal/profile', '27ca5ec0b65546acafe0a28ace8636bd', '0', '0', '1', '0', '0', '1515431757', '1515431757', null), ('7f35a6c6fa92434b98a21512f4a8a658', '设置', 'icon-cog3', '/settings', '', '0', '0', '1', '0', '9', '1515438052', '1515435272', null), ('814f789ed9714d068c297cde227656d4', '接口管理', 'icon-flip-vertical4', '/settings/system/apis', 'afab4e4d69cb4443a907fc856f418dcf', '0', '0', '1', '0', '1', '1515400780', '1515400780', null), ('879a38ef29214211b8e64e2a2d692aaa', '账号管理', 'icon-vcard', '/users/accounts', 'bc20b16c2e944223bf0789f12a887587', '0', '0', '1', '0', '0', '1515399880', '1515399880', null), ('afab4e4d69cb4443a907fc856f418dcf', '系统设置', '', '', 'e64f0a08e859429ca0d57d5c2bcfc3fc', '1', '1', '1', '0', '0', '1515400713', '1515400701', null), ('b42f7e56df4b4301bf5a7e6d3c3027e1', '用户管理', 'icon-users4', '/users/accounts', '', '0', '0', '1', '0', '0', '1515464081', '1515398578', null), ('bc20b16c2e944223bf0789f12a887587', '用户管理', '', '', 'b42f7e56df4b4301bf5a7e6d3c3027e1', '1', '1', '1', '0', '0', '1515464126', '1515398768', null), ('c3aa207798a94b4a8ec65e412dd9931d', '角色管理', 'icon-eye-plus', '/users/roles', 'bc20b16c2e944223bf0789f12a887587', '0', '0', '1', '0', '1', '1515399927', '1515399927', null), ('e64f0a08e859429ca0d57d5c2bcfc3fc', '系统设置', 'icon-wrench2', '/settings/system/menus', '7f35a6c6fa92434b98a21512f4a8a658', '0', '0', '1', '0', '0', '1515435300', '1515400665', null);
COMMIT;

-- ----------------------------
--  Table structure for `menu_rule`
-- ----------------------------
DROP TABLE IF EXISTS `menu_rule`;
CREATE TABLE `menu_rule` (
  `id` varchar(255) NOT NULL,
  `p_type` varchar(255) DEFAULT NULL,
  `v0` varchar(255) DEFAULT NULL,
  `v1` varchar(255) DEFAULT NULL,
  `v2` varchar(255) DEFAULT NULL,
  `v3` varchar(255) DEFAULT NULL,
  `v4` varchar(255) DEFAULT NULL,
  `v5` varchar(255) DEFAULT NULL,
  `created` bigint(20) DEFAULT NULL,
  `updated` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `menu_rule`
-- ----------------------------
BEGIN;
INSERT INTO `menu_rule` VALUES ('021fe1f1868b4712ad10155bcf3de6bf', 'p', 'a1feda2116114e4683019a319fed15da', '7f35a6c6fa92434b98a21512f4a8a658', '', '', '', '', '1515506722', '1515506722'), ('135961ac61464c1b8200bc1f19c989a1', 'p', 'a1feda2116114e4683019a319fed15da', 'afab4e4d69cb4443a907fc856f418dcf', '', '', '', '', '1515506722', '1515506722'), ('20b6b8a7aef645c4bdc305bfb3b16de2', 'p', 'a1feda2116114e4683019a319fed15da', '4f569a92c62e4611b7333f5bd07647ad', '', '', '', '', '1515506722', '1515506722'), ('49515f6cfe00455cb672fa80f202ba48', 'p', 'a1feda2116114e4683019a319fed15da', '0a6af3dd936c4c25a9835b52ddd5d2d0', '', '', '', '', '1515506722', '1515506722'), ('4e49aeb92c454e76b2e9dafb6f3929ff', 'p', 'a1feda2116114e4683019a319fed15da', '27ca5ec0b65546acafe0a28ace8636bd', '', '', '', '', '1515506722', '1515506722'), ('517bc63b62694c299dcafb0edd8512a6', 'p', 'a1feda2116114e4683019a319fed15da', '5397bf6cf9b04b2c856fa80bbd15134e', '', '', '', '', '1515506722', '1515506722'), ('5760741daa7c4ae3a8ff34b2d4e3ae54', 'p', 'a1feda2116114e4683019a319fed15da', '879a38ef29214211b8e64e2a2d692aaa', '', '', '', '', '1515506722', '1515506722'), ('6664d534b99847608b0c2995b90f7d9d', 'p', 'a1feda2116114e4683019a319fed15da', '814f789ed9714d068c297cde227656d4', '', '', '', '', '1515506722', '1515506722'), ('6c8899388a234cf296c2dc161175c5fb', 'p', 'a1feda2116114e4683019a319fed15da', 'e64f0a08e859429ca0d57d5c2bcfc3fc', '', '', '', '', '1515506722', '1515506722'), ('7beed47626e343c99cf9e43cb4c108e1', 'p', '21eba6d0e099449392f98a32802cfcb0', '345f520608b7481690b7046d77d5b4e0', '', '', '', '', '1515403585', '1515403585'), ('a2ef32ca7d7f4a9ba33bd85e4da627e7', 'p', 'a1feda2116114e4683019a319fed15da', 'b42f7e56df4b4301bf5a7e6d3c3027e1', '', '', '', '', '1515506722', '1515506722'), ('b178050022ec42b7a8fc2f2e072e5aa3', 'p', 'a1feda2116114e4683019a319fed15da', 'bc20b16c2e944223bf0789f12a887587', '', '', '', '', '1515506722', '1515506722'), ('c7fd0f05b4d844b6b83d78fd0ef0dce4', 'p', 'a1feda2116114e4683019a319fed15da', '345f520608b7481690b7046d77d5b4e0', '', '', '', '', '1515506722', '1515506722'), ('ce66d948bf404a27be5b8d256e45402c', 'g', '5b4a990536e947b584cfd13df072d352', 'a1feda2116114e4683019a319fed15da', '', '', '', '', '1515404202', '1515404202'), ('d0e5c27c3c91403882544b7b9fd2b57c', 'p', 'a1feda2116114e4683019a319fed15da', '146951887110406d9af92dadd45fa2ae', '', '', '', '', '1515506722', '1515506722'), ('ec06d7e3554a46ce914aea5635eb482f', 'p', 'a1feda2116114e4683019a319fed15da', 'c3aa207798a94b4a8ec65e412dd9931d', '', '', '', '', '1515506722', '1515506722');
COMMIT;

-- ----------------------------
--  Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '编号',
  `name` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色名称',
  `code` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色编码',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态',
  `updated` int(11) NOT NULL DEFAULT '0',
  `created` int(11) NOT NULL DEFAULT '0',
  `deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='角色表';

-- ----------------------------
--  Records of `role`
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES ('21eba6d0e099449392f98a32802cfcb0', '终端用户', 'customer', '0', '1514444065', '1514444065', null), ('590765a8b37140cb95fdd392070581ff', 'tester2', 'tester2', '0', '1514441583', '1514440693', '2017-12-28 15:37:33'), ('7baa76458f3c4e98a771282fce59fef4', '开发人员', 'developer', '0', '1514431085', '1514431085', null), ('8aad7831d3b64f05862f03f2bddf883c', 'tester', 'tester', '0', '1514440637', '1514440637', '2017-12-28 14:10:28'), ('a1feda2116114e4683019a319fed15da', '超级管理员', 'administrator', '0', '1514430799', '1514430799', null), ('d7c350d27b604bd6b9cfd32a0444e78e', '111111', '1111', '0', '1514431563', '1514431563', '2017-12-28 13:49:13'), ('ded21d51bf9142a7ad6b0bb640252189', '测试人员', 'tester', '0', '1514431062', '1514431062', null);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
