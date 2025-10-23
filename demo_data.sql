-- MySQL dump 10.13  Distrib 8.0.42, for Win64 (x86_64)
--
-- Host: localhost    Database: fantasyfutleague_db
-- ------------------------------------------------------
-- Server version	8.0.42

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `matches`
--

LOCK TABLES `matches` WRITE;
/*!40000 ALTER TABLE `matches` DISABLE KEYS */;
INSERT INTO `matches` VALUES (1,1,1,2,'2025-10-23','FINISHED',2,2,3),(2,2,1,2,'2025-10-23','FINISHED',1,3,2),(3,2,1,3,'2025-10-23','FINISHED',1,2,1),(4,2,1,4,'2025-10-23','FINISHED',1,2,1),(5,2,2,3,'2025-10-23','FINISHED',3,1,3),(6,2,2,4,'2025-10-23','FINISHED',4,0,1),(7,2,3,4,'2025-10-23','FINISHED',4,0,2);
/*!40000 ALTER TABLE `matches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `players`
--

LOCK TABLES `players` WRITE;
/*!40000 ALTER TABLE `players` DISABLE KEYS */;
INSERT INTO `players` VALUES (1,1,'Alisson','Brasil',22000000,7.17,'https://img.sofascore.com/api/v1/player/243609/image','Arquero'),(2,2,'Jan Oblak','Eslovenia',21000000,7,'https://img.sofascore.com/api/v1/player/69768/image','Arquero'),(3,4,'Emiliano Martínez','Argentina',19200000,7.1,'https://img.sofascore.com/api/v1/player/158263/image','Arquero'),(4,3,'Gianluigi Donnarumma','Italia',41000000,7.01,'https://img.sofascore.com/api/v1/player/824509/image','Arquero'),(5,3,'Achraf Hakimi','Marruecos',82000000,7.46,'https://img.sofascore.com/api/v1/player/814594/image','Lateral Derecho'),(6,1,'Jules Koundé','Francia',69000000,7.07,'https://img.sofascore.com/api/v1/player/827212/image','Lateral Derecho'),(7,2,'Nahuel Molina','Argentina',19400000,6.77,'https://img.sofascore.com/api/v1/player/831799/image','Lateral Derecho'),(8,NULL,'Diogo Dalot','Portugal',28000000,7.11,'https://img.sofascore.com/api/v1/player/843200/image','Lateral Derecho'),(9,NULL,'Pau Cubarsí','España',75000000,6.91,'https://img.sofascore.com/api/v1/player/1402913/image','Defensor Central'),(10,3,'Dean Huijsen','España',77000000,7.27,'https://img.sofascore.com/api/v1/player/1176744/image','Defensor Central'),(11,NULL,'Rúben Dias','Portugal',71000000,7.06,'https://img.sofascore.com/api/v1/player/318941/image','Defensor Central'),(12,3,'Antonio Rüdiger','Alemania',15500000,6.98,'https://img.sofascore.com/api/v1/player/142622/image','Defensor Central'),(13,1,'Marquinhos','Brasil',34000000,7.39,'https://img.sofascore.com/api/v1/player/155995/image','Defensor Central'),(14,2,'Leny Yoro','Francia',51000000,6.71,'https://img.sofascore.com/api/v1/player/1153315/image','Defensor Central'),(15,1,'Harry Maguire','Inglaterra',13900000,7.03,'https://img.sofascore.com/api/v1/player/149380/image','Defensor Central'),(16,2,'Cristian Romero','Argentina',46000000,7.18,'https://img.sofascore.com/api/v1/player/829932/image','Defensor Central'),(17,2,'Miloš Kerkez','Hungría',46000000,6.96,'https://img.sofascore.com/api/v1/player/1097425/image','Lateral Izquierdo'),(18,3,'Valentín Barco','Argentina',16500000,7.08,'https://img.sofascore.com/api/v1/player/1127057/image','Lateral Izquierdo'),(19,4,'Alejandro Balde','España',66000000,6.9,'https://img.sofascore.com/api/v1/player/997035/image','Lateral Izquierdo'),(20,1,'Álvaro Carreras','España',53000000,7.13,'https://img.sofascore.com/api/v1/player/1085081/image','Lateral Izquierdo'),(21,NULL,'Martín Zubimendi','España',62000000,7.06,'https://img.sofascore.com/api/v1/player/966837/image','Mediocampista Defensivo'),(22,NULL,'Aurélien Tchouaméni','Francia',78000000,7.13,'https://img.sofascore.com/api/v1/player/859025/image','Mediocampista Defensivo'),(23,2,'Frenkie de Jong','Países Bajos',42000000,7.01,'https://img.sofascore.com/api/v1/player/795222/image','Mediocampista Defensivo'),(24,1,'Declan Rice','Inglaterra',114000000,7.35,'https://img.sofascore.com/api/v1/player/856714/image','Mediocampista Defensivo'),(25,4,'Casemiro','Brasil',9100000,7.25,'https://img.sofascore.com/api/v1/player/122951/image','Mediocampista Defensivo'),(26,2,' Manuel Locatelli','Italia',29000000,7.27,'https://img.sofascore.com/api/v1/player/363860/image','Mediocampista Defensivo'),(27,NULL,'Enzo Fernández','Argentina',73000000,7.27,'https://img.sofascore.com/api/v1/player/974505/image','Mediocampista Central'),(28,4,'Alexis Mac Allister','Argentina',110000000,7.11,'https://img.sofascore.com/api/v1/player/895324/image','Mediocampista Central'),(29,NULL,'Federico Valverde','Uruguay',126000000,7.37,'https://img.sofascore.com/api/v1/player/831808/image','Mediocampista Central'),(30,1,'Adrien Rabiot','Francia',24000000,7.37,'https://img.sofascore.com/api/v1/player/250737/image','Mediocampista Central'),(31,3,'Bruno Guimarães','Brasil',82000000,7.2,'https://img.sofascore.com/api/v1/player/866469/image','Mediocampista Central'),(32,1,'Nicolò Barella','Italia',71000000,7.13,'https://img.sofascore.com/api/v1/player/363856/image','Mediocampista Central'),(33,3,'Teun Koopmeiners','Países Bajos',33000000,6.93,'https://img.sofascore.com/api/v1/player/803033/image','Mediocampista Central'),(34,3,'Giuliano Simeone','Argentina',38000000,6.98,'https://img.sofascore.com/api/v1/player/1099352/image','Mediocampista Por Derecha'),(35,NULL,'Michael Olise','Francia',125000000,7.76,'https://img.sofascore.com/api/v1/player/978838/image','Mediocampista Por Derecha'),(36,NULL,'Amad Diallo','Costa de Marfil',44000000,7.26,'https://img.sofascore.com/api/v1/player/971037/image','Mediocampista Por Derecha'),(37,3,'Filip Kostić','Serbia',4600000,7.09,'https://img.sofascore.com/api/v1/player/126588/image','Mediocampista Por Izquierda'),(38,NULL,'Nico Williams','España',76000000,7,'https://img.sofascore.com/api/v1/player/1085400/image','Mediocampista Por Izquierda'),(39,NULL,'Kingsley Coman','Francia',32000000,7.14,'https://img.sofascore.com/api/v1/player/280441/image','Mediocampista Por Izquierda'),(40,NULL,'Dominik Szoboszlai','Hungría',73000000,7.15,'https://img.sofascore.com/api/v1/player/869856/image','Mediocampista Ofensivo'),(41,NULL,'Dani Olmo','España',58000000,7.14,'https://img.sofascore.com/api/v1/player/789071/image','Mediocampista Ofensivo'),(42,2,'Bruno Fernandes','Portugal',46000000,7.73,'https://img.sofascore.com/api/v1/player/288205/image','Mediocampista Ofensivo'),(43,4,'Vinícius Júnior','Brasil',156000000,7.39,'https://img.sofascore.com/api/v1/player/868812/image','Extremo Izquierdo'),(44,2,'Raphinha','Brasil',81000000,7.84,'https://img.sofascore.com/api/v1/player/831005/image','Extremo Izquierdo'),(45,1,'Alejandro Garnacho','Argentina',42000000,6.94,'https://img.sofascore.com/api/v1/player/1135873/image','Extremo Izquierdo'),(46,NULL,'Gabriel Martinelli','Brasil',51000000,6.96,'https://img.sofascore.com/api/v1/player/922573/image','Extremo Izquierdo'),(47,NULL,'Khvicha Kvaratskhelia','Georgia',93000000,7.37,'https://img.sofascore.com/api/v1/player/889259/image','Extremo Izquierdo'),(48,NULL,'Luis Díaz','Colombia',76000000,7.26,'https://img.sofascore.com/api/v1/player/883537/image','Extremo Izquierdo'),(49,NULL,'Désiré Doué','Francia',93000000,7.37,'https://img.sofascore.com/api/v1/player/1154605/image','Extremo Derecho'),(50,NULL,'Lamine Yamal','España',215000000,7.81,'https://img.sofascore.com/api/v1/player/1402912/image','Extremo Derecho'),(51,4,'Mohamed Salah','Egipto',48000000,7.48,'https://img.sofascore.com/api/v1/player/159665/image','Extremo Derecho'),(52,1,'Bryan Mbeumo','Camerún',53000000,7.37,'https://img.sofascore.com/api/v1/player/927083/image','Extremo Derecho'),(53,NULL,'Rodrygo','Brasil',96000000,7.33,'https://img.sofascore.com/api/v1/player/910536/image','Extremo Derecho'),(54,2,'Brahim Díaz','Marruecos',42000000,6.88,'https://img.sofascore.com/api/v1/player/835485/image','Extremo Derecho'),(55,2,'Harry Kane','Inglaterra',81000000,7.86,'https://img.sofascore.com/api/v1/player/108579/image','Delantero Centro'),(56,NULL,'Kylian Mbappé','Francia',191000000,7.69,'https://img.sofascore.com/api/v1/player/826643/image','Delantero Centro'),(57,1,'Robert Lewandowski','Polonia',9400000,7.21,'https://img.sofascore.com/api/v1/player/41789/image','Delantero Centro'),(58,3,'Serhou Guirassy','Guinea',47000000,7.27,'https://img.sofascore.com/api/v1/player/328027/image','Delantero Centro'),(59,NULL,'Erling Haaland','Noruega',196000000,7.4,'https://img.sofascore.com/api/v1/player/839956/image','Delantero Centro'),(60,3,'Darwin Núñez','Uruguay',41000000,6.81,'https://img.sofascore.com/api/v1/player/924871/image','Delantero Centro'),(61,4,'Álvaro Morata','España',11400000,6.95,'https://img.sofascore.com/api/v1/player/125407/image','Delantero Centro'),(62,NULL,'Ederson','Brasil',16600000,6.93,'https://img.sofascore.com/api/v1/player/254491/image','Arquero'),(63,4,'Daniel Carvajal','España',7700000,7,'https://img.sofascore.com/api/v1/player/138572/image','Lateral Derecho'),(64,4,'Lisandro Martínez','Argentina',42000000,7.11,'https://img.sofascore.com/api/v1/player/859999/image','Defensor Central'),(65,4,'Francesco Acerbi','Italia',2900000,7.1,'https://img.sofascore.com/api/v1/player/126816/image','Defensor Central'),(66,NULL,'Theo Hernández','Francia',37000000,7.25,'https://img.sofascore.com/api/v1/player/788027/image','Lateral Izquierdo'),(67,NULL,'Granit Xhaka','Suiza',12700000,7.3,'https://img.sofascore.com/api/v1/player/117777/image','Mediocampista Defensivo'),(68,NULL,'Exequiel Palacios','Argentina',31000000,7.14,'https://img.sofascore.com/api/v1/player/822600/image','Mediocampista Central'),(69,NULL,'Isco','España',5200000,7.4,'https://img.sofascore.com/api/v1/player/103417/image','Mediocampista Ofensivo'),(70,4,'Luka Modrić','Croacia',4400000,7.4,'https://img.sofascore.com/api/v1/player/15466/image','Mediocampista Central'),(71,NULL,'Franco Mastantuono','Argentina',49000000,7.26,'https://img.sofascore.com/api/v1/player/1403559/image','Extremo Derecho'),(72,NULL,'Jérémy Doku','Bélgica',46000000,7.31,'https://img.sofascore.com/api/v1/player/934386/image','Extremo Izquierdo'),(73,NULL,'Rasmus Højlund','Dinamarca',49000000,6.82,'https://img.sofascore.com/api/v1/player/1086417/image','Delantero Centro');
/*!40000 ALTER TABLE `players` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `teams`
--

LOCK TABLES `teams` WRITE;
/*!40000 ALTER TABLE `teams` DISABLE KEYS */;
INSERT INTO `teams` VALUES (1,'user2team',2,'/static/uploads/20251023153229_badge1.png','433'),(2,'user3team',3,'/static/uploads/20251023153256_badge2.png','4231'),(3,'user4team',4,'/static/uploads/20251023153325_badge3.png','442'),(4,'user5team',5,'/static/uploads/20251023171207_badge4.png','433');
/*!40000 ALTER TABLE `teams` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `tournament_teams`
--

LOCK TABLES `tournament_teams` WRITE;
/*!40000 ALTER TABLE `tournament_teams` DISABLE KEYS */;
INSERT INTO `tournament_teams` VALUES (1,1,0),(2,1,0),(1,2,0),(2,2,0),(2,3,0),(2,4,0);
/*!40000 ALTER TABLE `tournament_teams` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `tournaments`
--

LOCK TABLES `tournaments` WRITE;
/*!40000 ALTER TABLE `tournaments` DISABLE KEYS */;
INSERT INTO `tournaments` VALUES (1,2,'Torneo Relámpago',10000000,'2025-10-23','2025-10-23',2,'FINISHED'),(2,4,'Liguilla Regional ',30000000,'2025-10-23','2025-10-24',1,'FINISHED'),(3,2,'Torneo Relámpago 2',10000000,'2025-10-24','2025-10-24',NULL,'NOT STARTED');
/*!40000 ALTER TABLE `tournaments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'user1','$2a$10$d8z7CEsWhmGTDvY..QL/i.8KThd/uaucdyG01IgTijiqTORtttzhu','correo1@gmail.com',520000000,'ADMIN'),(2,'user2','$2a$10$IEb2dMKbGIZ0lEjC6XQI5ehGW38CjDzGg7q0CzGMdsE.Q2zdfIE/S','correo2@gmail.com',44700000,'USER'),(3,'user3','$2a$10$LG9feVOlxsh3T9xw4Ai84.AesBFttf.arQfvbKALLJsI18Ira7cQO','correo3@gmail.com',25600000,'USER'),(4,'user4','$2a$10$y/U8QnEEFsElJ6gn9uVXGe5JOAfV.TSfF8L6irEDzCPYAW.iu749a','correo4@gmail.com',42400000,'USER'),(5,'user5','$2a$10$t0F0OOtK2YW8tHcxHr8nz.whQNeZgqmQQONKUgNV3lbXCPrY8lCmS','correo5@gmail.com',43300000,'USER');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-10-23 17:24:08
