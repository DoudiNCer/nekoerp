/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : 127.0.0.1:3306
 Source Schema         : nekoerp

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 25/04/2024 16:33:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for di
-- ----------------------------
DROP TABLE IF EXISTS `di`;
CREATE TABLE `di` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` int NOT NULL COMMENT '货物种类',
  `count` int NOT NULL COMMENT '入库数量',
  `created_at` datetime NOT NULL COMMENT '入库时间',
  `operator` int NOT NULL COMMENT '操作员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of di
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` char(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '货物名称',
  `price` decimal(10,2) DEFAULT NULL COMMENT '货物单价',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of goods
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '权限名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`) VALUES (1, '管理员');
INSERT INTO `role` (`id`, `name`) VALUES (2, '员工');
COMMIT;

-- ----------------------------
-- Table structure for storage
-- ----------------------------
DROP TABLE IF EXISTS `storage`;
CREATE TABLE `storage` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` int unsigned NOT NULL COMMENT '货物种类',
  `count` int unsigned NOT NULL COMMENT '数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of storage
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tiao
-- ----------------------------
DROP TABLE IF EXISTS `tiao`;
CREATE TABLE `tiao` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` int NOT NULL COMMENT '货物种类',
  `count` int NOT NULL COMMENT '出库数量',
  `created_at` datetime NOT NULL COMMENT '出库时间',
  `operator` int NOT NULL COMMENT '操作员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of tiao
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户名',
  `password` char(129) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户密码',
  `role` int DEFAULT NULL COMMENT '用户权限',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `blocked` tinyint(1) DEFAULT '0' COMMENT '是否禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `account`, `password`, `role`, `created_at`, `updated_at`, `blocked`) VALUES (1, 'admin', '123456', 1, '2024-04-25 15:27:03', '2024-04-25 15:27:06', 0);
INSERT INTO `user` (`id`, `account`, `password`, `role`, `created_at`, `updated_at`, `blocked`) VALUES (2, 'test', '123456', 2, '2024-04-25 15:27:27', '2024-04-25 15:27:29', 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
