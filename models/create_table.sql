CREATE TABLE `user` (
    `id` bigint(20) NOT NULL,
    `uswe_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`uswe_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;