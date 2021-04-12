package database

func Reset() {

	exec(`DROP TABLE
		Message,
		Invite,
		ChannelPermissions,
		Channel,
		Role,
		ServerMember,
		Server,
		Account;`)

	exec(`CREATE TABLE Account (
		id INT NOT NULL AUTO_INCREMENT,
		email VARCHAR(50) NOT NULL,
		username VARCHAR(50) NOT NULL,
		password BINARY(32) NOT NULL,
		salt BINARY(32) NOT NULL,
		avatar VARCHAR(255) NOT NULL,
		PRIMARY KEY (id)
		);`)

	exec(`CREATE TABLE Server (
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(25),
		description VARCHAR(100),
		image VARCHAR(255) NOT NULL,
		PRIMARY KEY (id)
		)`)

	exec(`CREATE TABLE ServerMember (
		server_id INT NOT NULL,
		account_id Int NOT NULL,
		role_id INT NOT NULL,
		alias VARCHAR(50) NOT NULL,
		PRIMARY KEY (server_id, account_id),
		FOREIGN KEY (server_id) REFERENCES Server (id) ON DELETE CASCADE,
		FOREIGN KEY (account_id) REFERENCES Account (id) ON DELETE CASCADE
		);`)

	exec(`CREATE TABLE Role (
		id INT NOT NULL AUTO_INCREMENT,
		ranking INT NOT NULL,
		server_id INT NOT NULL,
		name VARCHAR(50) NOT NULL,
		permissions TINYINT UNSIGNED NOT NULL, 
		PRIMARY KEY (id),
		FOREIGN KEY (server_id) REFERENCES Server (id) ON DELETE CASCADE
		);`)

	exec(`CREATE TABLE Channel (
		id INT NOT NULL AUTO_INCREMENT, 
		server_id INT NOT NULL, 
		name VARCHAR(25) NOT NULL, 
		PRIMARY KEY (id), 
		FOREIGN KEY (server_id) REFERENCES Server (id) ON DELETE CASCADE
		);`)

	exec(`CREATE TABLE ChannelPermissions (
		role_id INT NOT NULL, 
		channel_id INT NOT NULL, 
		permissions TINYINT UNSIGNED NOT NULL,
		PRIMARY KEY (role_id, channel_id), 
		FOREIGN KEY (role_id) REFERENCES Role (id) ON DELETE CASCADE, 
		FOREIGN KEY (channel_id) REFERENCES Channel (id) ON DELETE CASCADE
		);`)

	exec(`Create TABLE Invite (
		id INT NOT NULL AUTO_INCREMENT, 
		server_id Int NOT NULL, 
		code VARCHAR(10), 
		PRIMARY KEY (id), 
		FOREIGN KEY (server_id) REFERENCES Server (id) ON DELETE CASCADE
		);`)

	exec(`Create TABLE Message (
		id INT NOT NULL AUTO_INCREMENT, 
		channel_id Int NOT NULL, 
		account_id Int NOT NULL, 
		text VARCHAR(255), 
		media VARCHAR(255), 
		parent_id Int,
		is_edited boolean Default false,
		time_posted DATETIME NOT NULL, 
		PRIMARY KEY (id), 
		FOREIGN KEY (channel_id) REFERENCES Channel (id) ON DELETE CASCADE, 
		FOREIGN KEY (account_id) REFERENCES Account (id) ON DELETE CASCADE
		);`)

}
