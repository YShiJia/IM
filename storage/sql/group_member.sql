CREATE TABLE `im_group_member` (
                                   `id` int(11) NOT NULL COMMENT '成员关系id',
                                   `group_id` int(11) NOT NULL COMMENT '群聊id',
                                   `member_id` int(11) NOT NULL COMMENT '成员id',
                                   `created_at` timestamp NOT NULL COMMENT '创建时间',
                                   `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `group_id_index` (`group_id`),
                                   UNIQUE KEY `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='群聊成员表';