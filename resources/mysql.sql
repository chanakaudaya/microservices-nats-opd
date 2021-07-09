CREATE TABLE `patient_details` (
  `id` int(10) unsigned NOT NULL,
  `full_name` varchar(100) NOT NULL,
`address` varchar(255),
`sex` varchar(10),
`phone` int(15) unsigned,
`remarks` varchar(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE `test` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;