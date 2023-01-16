CREATE TABLE `order_info` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `order_num` char(22) NOT NULL DEFAULT '' COMMENT '订单号',
                              `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
                              `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
                              `goods_prices` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品单价',
                              `goods_num` int(11) NOT NULL DEFAULT '0' COMMENT '商品数量',
                              `goods_all_prices` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品应有总价=商品单价*商品数量',
                              `goods_all_real_prices` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品实际总价',
                              `order_prices` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单应有价格',
                              `order_real_prices` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单实际价格',
                              `pay_id` varchar(255) NOT NULL DEFAULT '' COMMENT '支付号',
                              `business_id` int(11) NOT NULL DEFAULT '0' COMMENT '商家ID',
                              `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '订单状态 1：未支付 11：支付中 21：支付成功 31：支付失败 41：订单超时 51：退单 ',
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              KEY `goods_id` (`goods_id`) USING BTREE,
                              KEY `user_id` (`user_id`) USING BTREE,
                              KEY `order_num` (`order_num`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='订单表';

