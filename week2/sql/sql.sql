
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE `go_up_test` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';

USE `go_up_test`;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `age` int NULL DEFAULT NULL COMMENT '年龄',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


INSERT INTO `user` VALUES (1, 'User001', 18);
INSERT INTO `user` VALUES (2, 'User002', 30);

SET FOREIGN_KEY_CHECKS = 1;

