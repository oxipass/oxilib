package bslib

const sqlCreateTableFields = `
	CREATE TABLE fields (
   		field_id 		    CHAR PRIMARY KEY NOT NULL,
   		field_type_id 		CHAR NOT NULL,
		field_value         VARCHAR NOT NULL,
		creation_timestamp  DATETIME NOT NULL,
   		update_timestamp    DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (field_type_id) REFERENCES fields_definitions(field_type_id)
	)
`
