--
-- Table structure for table `urls`
--

DROP TABLE IF EXISTS `urls`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `urls`
(
    `id`           bigint                                  NOT NULL AUTO_INCREMENT,
    `user_id`      int                                     NOT NULL,
    `original_url` text COLLATE utf8mb4_unicode_ci         NOT NULL,
    `short_code`   varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `is_custom`    tinyint(1) NOT NULL DEFAULT '0',
    `views`        int                                     NOT NULL DEFAULT '0',
    `expired_at`   bigint                                  NOT NULL,
    `created_at`   bigint                                           DEFAULT NULL,
    `updated_at`   bigint                                           DEFAULT NULL,
    `deleted_at`   bigint                                           DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `short_code` (`short_code`),
    KEY            `idx_short_code` (`short_code`),
    KEY            `idx_expired_at_short_code` (`short_code`,`expired_at`),
    KEY            `idx_user_id` (`user_id`),
    KEY            `idx_expired_at_user_id` (`user_id`,`expired_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `urls`
--

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users`
(
    `id`            int                                     NOT NULL AUTO_INCREMENT,
    `email`         varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password_hash` text COLLATE utf8mb4_unicode_ci         NOT NULL,
    `created_at`    bigint DEFAULT NULL,
    `updated_at`    bigint DEFAULT NULL,
    `deleted_at`    bigint DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    KEY             `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--


--
-- Dumping routines for database 'urlshortener'
--
