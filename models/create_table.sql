DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
         `id` int(11) NOT NULL AUTO_INCREMENT,
         `community_id` int(10) unsigned NOT NULL,
         `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
         `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
         `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
         `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
         PRIMARY KEY (`id`),
         UNIQUE KEY `idx_community_id` (`community_id`),
         UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community` VALUES ('1', '1', 'Go', 'Golang', '2023-01-14 08:00:00', '2023-01-14 08:00:00');
INSERT INTO `community` VALUES ('2', '2', 'leetcode', '刷题冲冲冲！！！', '2023-01-14 08:00:00', '2023-01-14 08:00:00');
INSERT INTO `community` VALUES ('3', '3', 'CS:GO', 'Rush B！！！', '2023-01-14 08:00:00', '2023-01-14 08:00:00');
INSERT INTO `community` VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2023-01-14 08:00:00', '2023-01-14 08:00:00');
INSERT INTO `community` VALUES ('5', '5', 'daily', '记录daily日常!', '2023-01-14 08:00:00', '2023-01-14 08:00:00');