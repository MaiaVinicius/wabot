-- MySQL dump 10.13  Distrib 5.7.24, for macos10.14 (x86_64)
--
-- Host: 127.0.0.1    Database: wabot
-- ------------------------------------------------------
-- Server version	5.7.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `wabot_blacklist`
--

DROP TABLE IF EXISTS `wabot_blacklist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_blacklist` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` int(11) DEFAULT NULL,
  `phone` char(13) DEFAULT NULL,
  `reason` varchar(25) DEFAULT NULL,
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_blacklist`
--

LOCK TABLES `wabot_blacklist` WRITE;
/*!40000 ALTER TABLE `wabot_blacklist` DISABLE KEYS */;
/*!40000 ALTER TABLE `wabot_blacklist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_blacklist_terms`
--

DROP TABLE IF EXISTS `wabot_blacklist_terms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_blacklist_terms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `term` varchar(155) DEFAULT NULL,
  `must_match` bit(1) DEFAULT b'1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_blacklist_terms`
--

LOCK TABLES `wabot_blacklist_terms` WRITE;
/*!40000 ALTER TABLE `wabot_blacklist_terms` DISABLE KEYS */;
INSERT INTO `wabot_blacklist_terms` VALUES (1,'fuder',_binary ''),(2,'porra',_binary ''),(3,'tomar no',_binary ''),(4,'spam',_binary ''),(5,'quem é',_binary ''),(6,'não manda',_binary ''),(7,'saco',_binary ''),(8,'engano',_binary ''),(9,' cu ',_binary ''),(10,'foda',_binary ''),(11,'puta',_binary ''),(12,'caralho',_binary ''),(13,'telefone não é',_binary ''),(14,'celular não é',_binary ''),(15,'errado',_binary '');
/*!40000 ALTER TABLE `wabot_blacklist_terms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_contacts`
--

DROP TABLE IF EXISTS `wabot_contacts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_contacts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `saved` bit(1) DEFAULT b'1',
  `sender_id` int(11) DEFAULT NULL,
  `phone` char(13) DEFAULT NULL,
  `name` varchar(155) DEFAULT NULL,
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_contacts`
--

LOCK TABLES `wabot_contacts` WRITE;
/*!40000 ALTER TABLE `wabot_contacts` DISABLE KEYS */;
/*!40000 ALTER TABLE `wabot_contacts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_log`
--

DROP TABLE IF EXISTS `wabot_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `project_id` int(11) DEFAULT NULL,
  `log_type_id` int(11) DEFAULT NULL,
  `message` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=707 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `wabot_log_type`
--

DROP TABLE IF EXISTS `wabot_log_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_log_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(155) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_log_type`
--

LOCK TABLES `wabot_log_type` WRITE;
/*!40000 ALTER TABLE `wabot_log_type` DISABLE KEYS */;
INSERT INTO `wabot_log_type` VALUES (1,'Mensagem informativa'),(2,'Erro');
/*!40000 ALTER TABLE `wabot_log_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_project`
--

DROP TABLE IF EXISTS `wabot_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(155) DEFAULT NULL,
  `active` bit(1) DEFAULT b'1',
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `base_price` double DEFAULT '0.37',
  `status_id` int(11) DEFAULT '1',
  `license_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_project`
--

LOCK TABLES `wabot_project` WRITE;
/*!40000 ALTER TABLE `wabot_project` DISABLE KEYS */;
INSERT INTO `wabot_project` VALUES (2,'Teste',_binary '\0','2018-09-15 13:51:20',0,1,105);
/*!40000 ALTER TABLE `wabot_project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_project_status`
--

DROP TABLE IF EXISTS `wabot_project_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_project_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(155) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_project_status`
--

LOCK TABLES `wabot_project_status` WRITE;
/*!40000 ALTER TABLE `wabot_project_status` DISABLE KEYS */;
INSERT INTO `wabot_project_status` VALUES (1,'Em produção'),(2,'Cancelado'),(3,'Pausado'),(4,'Em configuração');
/*!40000 ALTER TABLE `wabot_project_status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_queue`
--

DROP TABLE IF EXISTS `wabot_queue`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_queue` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sender_id` int(11) DEFAULT NULL,
  `phone` char(13) DEFAULT NULL,
  `message` blob,
  `active` bit(1) DEFAULT b'1',
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `send_date` date DEFAULT NULL,
  `send_time` time DEFAULT NULL,
  `price` double DEFAULT '0.3',
  `license_id` int(11) DEFAULT NULL,
  `appointment_id` int(11) DEFAULT NULL,
  `event_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `wabot_queue_sender_id_index` (`sender_id`),
  KEY `wabot_queue_license_id_index` (`license_id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `wabot_response`
--

DROP TABLE IF EXISTS `wabot_response`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_response` (
  `id` varchar(55) NOT NULL,
  `project_id` int(11) DEFAULT NULL,
  `sent_message_id` int(11) DEFAULT NULL,
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `phone` char(13) DEFAULT NULL,
  `message` blob,
  `status_id` int(11) DEFAULT NULL COMMENT '1- recebido\n2- enviado\n3- enviado e nao lido\n4- enviado e lido',
  `from_me` bit(1) DEFAULT NULL,
  `license_id` int(11) DEFAULT NULL,
  `appointment_id` int(11) DEFAULT NULL,
  `event_id` int(11) DEFAULT NULL,
  `sync` bit(1) DEFAULT b'0',
  `sync_at` timestamp NULL DEFAULT NULL,
  `auto_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  UNIQUE KEY `wabot_response_auto_id_uindex` (`auto_id`),
  UNIQUE KEY `wabot_response_pk` (`auto_id`),
  UNIQUE KEY `wabot_response_id_uindex` (`id`),
  KEY `wabot_response_license_id_index` (`license_id`),
  KEY `wabot_response_project_id_index` (`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=68542 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `wabot_sender`
--

DROP TABLE IF EXISTS `wabot_sender`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_sender` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` int(11) DEFAULT NULL,
  `description` varchar(155) DEFAULT NULL,
  `phone` char(13) DEFAULT NULL,
  `active` bit(1) DEFAULT b'1',
  `status_id` int(11) DEFAULT '1',
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `wabot_sender_project_id_index` (`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_sender`
--

LOCK TABLES `wabot_sender` WRITE;
/*!40000 ALTER TABLE `wabot_sender` DISABLE KEYS */;
INSERT INTO `wabot_sender` VALUES (1,2,'SENDER EXEMPLO','5521999999999',_binary '',2,'2019-01-12 18:44:19');
/*!40000 ALTER TABLE `wabot_sender` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_sender_status`
--

DROP TABLE IF EXISTS `wabot_sender_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_sender_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(155) DEFAULT NULL,
  `active` bit(1) DEFAULT b'1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_sender_status`
--

LOCK TABLES `wabot_sender_status` WRITE;
/*!40000 ALTER TABLE `wabot_sender_status` DISABLE KEYS */;
INSERT INTO `wabot_sender_status` VALUES (1,'Ativo',_binary ''),(2,'Bloqueado',_binary '\0'),(3,'Pendente de autenticação',_binary '\0');
/*!40000 ALTER TABLE `wabot_sender_status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wabot_sent`
--

DROP TABLE IF EXISTS `wabot_sent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wabot_sent` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sender_id` int(11) DEFAULT NULL,
  `phone` char(13) DEFAULT NULL,
  `message` text,
  `status` varchar(3) DEFAULT '200',
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `price` double DEFAULT NULL,
  `metadata` blob,
  `license_id` int(11) DEFAULT NULL,
  `appointment_id` int(11) DEFAULT NULL,
  `event_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `wabot_sent_license_id_index` (`license_id`),
  KEY `wabot_sent_sender_id_index` (`sender_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1234 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wabot_sent`
--

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-01-26 14:26:10
