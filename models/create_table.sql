DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `community_id` bigint(20) unsigned NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
    `introduction` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`) USING BTREE,
    UNIQUE KEY `idx_community_name` (`community_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `community` VALUES (1, 1, 'Go', 'Golang', '2024-05-01 13:00:00', '2024-05-01 13:00:00');
INSERT INTO `community` VALUES (2, 2, 'leecode', '刷題', '2024-05-02 13:00:00', '2024-05-02 13:00:00');
INSERT INTO `community` VALUES (3, 3, 'CS:GO', 'RUSH B', '2024-05-03 13:00:00', '2024-05-03 13:00:00');
INSERT INTO `community` VALUES (4, 4, 'LOL', '歡迎來到英雄聯盟', '2024-05-04 13:00:00', '2024-05-04 13:00:00');

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) unsigned NOT NULL COMMENT '帖子ID',
    `title` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '標題',
    `content` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '內容',
    `author_id` bigint(20) unsigned NOT NULL COMMENT '作者的用戶ID',
    `community_id` bigint(20) unsigned NOT NULL COMMENT '帖子所屬社區ID',
    `status` tinyint(4) NULL DEFAULT '1' COMMENT '帖子狀態，0-已刪除，1-正常',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_author_id` (`author_id`) USING BTREE,
    UNIQUE KEY `idx_community_id` (`community_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;