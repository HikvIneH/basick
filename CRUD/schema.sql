CREATE TABLE `user_review` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `orderid` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `productid` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `userid` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rating` float NOT NULL,
  `usersreview` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `createdat` timestamp NOT NULL DEFAULT current_timestamp(),
  `updatedat` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci