CREATE TABLE `im_user` (
                           `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'userid',
                           `social_id` VARCHAR(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '社交id',
                           `nickname` VARCHAR(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
                           `password` VARCHAR(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
                           `email` VARCHAR(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
                           `status` INT(11) NOT NULL COMMENT '状态(0:正常;1:冻结;2:删除)',
                           `gender` INT(11) DEFAULT NULL COMMENT '性别',
                           `created_at` TIMESTAMP NOT NULL COMMENT '创建日期',
                           `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除日期',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `nick_id` (`social_id`) COMMENT '社交id唯一'
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='im系统user表';