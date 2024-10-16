CREATE TABLE `im_group` (
                            `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '群聊id',
                            `social_id` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '群聊社交id',
                            `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '群聊名称',
                            `create_user_id` int(11) NOT NULL COMMENT '创建者id',
                            `created_at` timestamp NOT NULL COMMENT '创建时间',
                            `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `social_id_index` (`social_id`),
                            UNIQUE KEY `create_user_id_index` (`create_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='群聊数据列表';