package solarsystem

const (
	SQLWEATHER = "CREATE TABLE IF NOT EXISTS weather_days(" +
		"	`day` int NOT NULL," +
		"	`condition` varchar(20) NOT NULL," +
		"	PRIMARY KEY (`day`)" +
		")"
	SQLINSERTSOLARSYSTEM = "INSERT INTO weather_days("+
		"	`day`, " +
		"	`condition`" +
		")" +
		"VALUES (?,?)"

	SQLEXISTS = "SELECT EXISTS(SELECT * FROM weather_days)"

	SQLQUERY = "SELECT `condition` FROM weather_days WHERE `day` = %d"
)