/*
SQLyog Enterprise v12.14 (64 bit)
MySQL - 5.7.19 : Database - mxshop_order_srv
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`mxshop_order_srv` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `mxshop_order_srv`;

/*Table structure for table `ordergoods` */

DROP TABLE IF EXISTS `ordergoods`;

CREATE TABLE `ordergoods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `order` int(11) DEFAULT NULL,
  `goods` int(11) DEFAULT NULL,
  `goods_name` varchar(100) DEFAULT NULL,
  `goods_image` varchar(200) DEFAULT NULL,
  `goods_price` float DEFAULT NULL,
  `nums` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ordergoods_order` (`order`),
  KEY `idx_ordergoods_goods` (`goods`),
  KEY `idx_ordergoods_goods_name` (`goods_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `ordergoods` */

/*Table structure for table `orderinfo` */

DROP TABLE IF EXISTS `orderinfo`;

CREATE TABLE `orderinfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `user` int(11) DEFAULT NULL,
  `order_sn` varchar(30) DEFAULT NULL,
  `pay_type` varchar(20) DEFAULT NULL COMMENT 'alipay(支付宝)， wechat(微信)',
  `status` varchar(20) DEFAULT NULL COMMENT 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)',
  `trade_no` varchar(100) DEFAULT NULL COMMENT '交易号',
  `order_mount` float DEFAULT NULL,
  `pay_time` datetime DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `signer_name` varchar(20) DEFAULT NULL,
  `singer_mobile` varchar(11) DEFAULT NULL,
  `post` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_orderinfo_user` (`user`),
  KEY `idx_orderinfo_order_sn` (`order_sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `orderinfo` */

/*Table structure for table `shoppingcart` */

DROP TABLE IF EXISTS `shoppingcart`;

CREATE TABLE `shoppingcart` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `user` int(11) DEFAULT NULL,
  `goods` int(11) DEFAULT NULL,
  `nums` int(11) DEFAULT NULL,
  `checked` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_shoppingcart_user` (`user`),
  KEY `idx_shoppingcart_goods` (`goods`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `shoppingcart` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
