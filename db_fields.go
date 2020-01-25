package bslib

const sqlCreateTableFields = `
	CREATE TABLE IF NOT EXISTS fields (
   		field_id 		    CHAR PRIMARY KEY NOT NULL,
   		field_def_id 		CHAR NOT NULL,
		field_value         VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (field_def_id) REFERENCES fields_definitions(field_type_id)
	)
`
