/*
SQLyog Enterprise v12.14 (64 bit)
MySQL - 5.7.19 : Database - mxshop_user_srv
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`mxshop_user_srv` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `mxshop_user_srv`;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `mobile` varchar(11) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nick_name` varchar(20) DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `gender` varchar(6) DEFAULT 'male' COMMENT 'female表示女,male表示男',
  `role` int(11) DEFAULT '1' COMMENT '1表示普通用户,2表示管理员',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_mobile` (`mobile`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4;

/*Data for the table `user` */

insert  into `user`(`id`,`add_time`,`update_time`,`deleted_at`,`is_deleted`,`mobile`,`password`,`nick_name`,`birthday`,`gender`,`role`) values 
(31,'2025-03-17 13:09:29.173','2025-03-17 13:09:29.173',NULL,0,'18782222220','$pbkdf2-sha512$21PX9fEHuUDR7eET$79f030836ef80a6cd8a28a2a53d50fba20bb474848e9c1abbfe76c4ef77dd022','bobby0',NULL,'male',1),
(32,'2025-03-17 13:09:29.175','2025-03-17 13:09:29.175',NULL,0,'18782222221','$pbkdf2-sha512$8B5WtSHMbPO3WLEp$426d8570c53e9c1d7e90b775f7876b1c83e63e45b591ca2290f7291e3fbeeda3','bobby1',NULL,'male',1),
(33,'2025-03-17 13:09:29.176','2025-03-17 13:09:29.176',NULL,0,'18782222222','$pbkdf2-sha512$cKNRyeKL6Xanh7aj$b3e360b37c80058b9ce0d433a03388183bc13ac36ac62de0acd7cdcdb3789029','bobby2',NULL,'male',1),
(34,'2025-03-17 13:09:29.178','2025-03-17 13:09:29.178',NULL,0,'18782222223','$pbkdf2-sha512$ogBDM00ysw5qlSpg$f736d71daa2aff419167cfd7ae20378bce2c7d5d7b114f796377b4ea8909eae3','bobby3',NULL,'male',1),
(35,'2025-03-17 13:09:29.179','2025-03-17 13:09:29.179',NULL,0,'18782222224','$pbkdf2-sha512$Ie0AAn1ygA873zXc$fa01fe051ffb64a228d8b9fa32331d50f98c3864a76733948a3d6d5b7fc1f252','bobby4',NULL,'male',1),
(36,'2025-03-17 13:09:29.180','2025-03-17 13:09:29.180',NULL,0,'18782222225','$pbkdf2-sha512$fQ0ZAiXeOEnbZOPE$c99ec00d24318daebd7cb0d73d9e780ce96b15f489a319539b479b72c1fadead','bobby5',NULL,'male',1),
(37,'2025-03-17 13:09:29.181','2025-03-17 13:09:29.181',NULL,0,'18782222226','$pbkdf2-sha512$gQ4LgUhy1m5YzyKW$0b099b6d3ab11f74e4d11801374ec6c79c2da18436340157a2f5a1ea0b22bf55','bobby6',NULL,'male',2),
(38,'2025-03-17 13:09:29.182','2025-03-17 13:09:29.182',NULL,0,'18782222227','$pbkdf2-sha512$6F7VdEjPgPLog0qr$aa8af82ab1c1701fa035cfb6581e0c8ed42f316e250714feeeaa59cc93f7af9a','bobby7',NULL,'male',1),
(39,'2025-03-17 13:09:29.183','2025-03-17 13:09:29.183',NULL,0,'18782222228','$pbkdf2-sha512$W3WVW1S2bfnVt9ar$ddaa94571a22d89d56ddb3c6c86636eac0b4ed9d1c19a4d7e01cb7b6b554be26','bobby8',NULL,'male',1),
(40,'2025-03-17 13:09:29.184','2025-03-17 13:09:29.184',NULL,0,'18782222229','$pbkdf2-sha512$CcFkhoxaLNfAcdTc$3de81984bf491c420f21f72c29586cca62ae71f7d58ff2a99d6185fd58290e70','bobby9',NULL,'male',1);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
