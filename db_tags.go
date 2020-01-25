package bslib

const sqlCreateTableTags = `
	CREATE TABLE IF NOT EXISTS tags (
   		tag_id 		   	    INT PRIMARY KEY NOT NULL,
		name         		VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`
