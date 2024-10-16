CREATE TABLE `im_friend` (
                             `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '好友关系列表',
                             `user_id` int(11) NOT NULL COMMENT '用户id',
                             `friend_id` int(11) NOT NULL COMMENT '好友id',
                             `created_at` timestamp NOT NULL COMMENT '创建时间',
                             `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                             PRIMARY KEY (`id`),
                             KEY `user_id_index` (`user_id`),
                             KEY `friend_id_index` (`friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='好友关系列表';