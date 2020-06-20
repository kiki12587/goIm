/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : 127.0.0.1:3306
 Source Schema         : goim

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 19/06/2020 21:13:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for goim_users
-- ----------------------------
DROP TABLE IF EXISTS `goim_users`;
CREATE TABLE `goim_users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `groupid` int(10) UNSIGNED NOT NULL DEFAULT 1,
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一邮箱',
  `password` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` tinyint(4) NOT NULL,
  `joindate` int(10) UNSIGNED NOT NULL,
  `joinip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `register_type` tinyint(3) NOT NULL,
  `openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `usersig` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'im用户签名',
  `expire` int(10) NULL DEFAULT 0 COMMENT '签名过期时间',
  `nick` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户昵称',
  `avatar` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户头像',
  `gender` tinyint(1) NULL DEFAULT 0 COMMENT '用户性别 0=\"未知\"  1=\"男\"  2=\"女\"',
  `self_signature` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '个性签名',
  `mobile` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE COMMENT '改成邮箱了'
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goim_users
-- ----------------------------
INSERT INTO `goim_users` VALUES (1, 1, '康康', 'MfRbwVwos9/7eIbYauATbQ==', 1, 1589300777, '127.0.0.1', '2014-01-07 09:32:12%!(EXTRA string=:成功注册)', 1, 'bqtcsaf1t0ijo27jd0dg', NULL, NULL, '', '', 0, '', NULL);
INSERT INTO `goim_users` VALUES (3, 1, '康康抗', 'MfRbwVwos9/7eIbYauATbQ==', 1, 1589300943, '127.0.0.1', '2014-01-07 09:32:12:成功注册', 1, 'bqtctjv1t0ikci14hsg0', NULL, NULL, '', '', 0, '', NULL);
INSERT INTO `goim_users` VALUES (4, 1, 'kk', 'MfRbwVwos9/7eIbYauATbQ==', 1, 1589474551, '127.0.0.1', '2014-01-07 09:32:12:成功注册', 1, 'bqun9tv1t0ij2e0q7nog', NULL, NULL, '', '', 0, '', NULL);
INSERT INTO `goim_users` VALUES (5, 1, '呃呃', 'PIpTXMTeQPfPlh5UtOhmHw==', 1, 1589809479, '127.0.0.1', '2014-01-07 09:32:12:成功注册', 1, 'br192hv1t0iles42io1g', NULL, NULL, '', '', 0, '', NULL);
INSERT INTO `goim_users` VALUES (6, 1, '1569482471@qq.com', 'jh6wlX4ueuAUKIy0oYGicQ==', 1, 1589985675, '127.0.0.1', '2014-01-07 09:32:12:成功注册', 1, 'br2k32v1t0im9p6l3h40', 'eJwszsFqhDAQxvF3mWuLTBI1TaCHWFopFnpoodCbkqlOrTarQWSXffcF9frn48d3gc*3j4TWwBOBlZmRiHi-RfY0Rv5hmsBCM8leyUVE5MGE-E91KcK*m31fh8AerEgRVS7z9BBmbsHCNxXdsytO*gu9aF4NRXenq5bc77q8jE-l4OpzNT6U7--u8SAjDwRWZEagVNrovS7bE5kgXG8BAAD--4T5NRM_', 1593615797, '', '', 0, '', NULL);

SET FOREIGN_KEY_CHECKS = 1;
