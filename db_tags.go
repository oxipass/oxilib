package bslib

const sqlCreateTableTags = `
	CREATE TABLE IF NOT EXISTS tags (
   		tag_id 		   	    INTEGER PRIMARY KEY AUTOINCREMENT,
		name         		VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

const sqlCreateTableItemsTags = `
	CREATE TABLE IF NOT EXISTS items_tags (
        it_id				INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id 		    INT NOT NULL,
		tag_id				INT NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (item_id) REFERENCES items(item_id),
		FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
	)
`
