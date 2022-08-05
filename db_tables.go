package oxilib

// sqlCreateTableFields is the SQL query to create a table 'fields'.
const sqlCreateTableFields = `
	CREATE TABLE IF NOT EXISTS fields (
   		field_id 		    INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id             INTEGER NOT NULL,
		field_icon          VARCHAR NOT NULL,
   		field_name 			VARCHAR NOT NULL,
		value_type          VARCHAR NOT NULL,
		field_value         VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (item_id) REFERENCES items(item_id)
	)
`

// sqlCreateTableItems is the SQL query to create a table 'items'.
const sqlCreateTableItems = `
	CREATE TABLE IF NOT EXISTS items (
   		item_id 		    INTEGER PRIMARY KEY AUTOINCREMENT,
   		name 			    VARCHAR NOT NULL,
		icon             	VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

// sqlCreateTableSettings is the SQL query to create a table 'settings'.
const sqlCreateTableSettings = `
	CREATE TABLE IF NOT EXISTS settings (
   		database_id 		VARCHAR PRIMARY KEY NOT NULL,
   		keyword 			VARCHAR NOT NULL,
		crypt_id            VARCHAR NOT NULL,
		language            VARCHAR NOT NULL,
		database_version 	INTEGER NOT NULL,
   		update_timestamp    DATETIME NOT NULL,
		sync_timestamp		DATETIME NOT NULL
	)
`

// sqlCreateTableTags is the SQL query to create a table 'tags'.
const sqlCreateTableTags = `
	CREATE TABLE IF NOT EXISTS tags (
   		tag_id 		   	    INTEGER PRIMARY KEY AUTOINCREMENT,
		extid 				VARCHAR NOT NULL,
		name         		VARCHAR NOT NULL,
		color        		VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

// sqlCreateTableTemplateItems is the SQL query to create a table 'template_items'.
const sqlCreateTableTemplateItems = `
	CREATE TABLE IF NOT EXISTS template_items (
   		item_id 		    VARCHAR NOT NULL PRIMARY KEY,
   		name 			    VARCHAR NOT NULL,
		icon             	VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

// sqlCreateTableTemplateFields is the SQL query to create a table 'template_fields'.
const sqlCreateTableTemplateFields = `
	CREATE TABLE IF NOT EXISTS template_fields (
   		field_id 		    VARCHAR NOT NULL PRIMARY KEY,
		item_id             VARCHAR NOT NULL,
		field_icon          VARCHAR NOT NULL,
   		field_name 			VARCHAR NOT NULL,
		value_type          VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (item_id) REFERENCES templates_items(item_id)
	)
`

// sqlCreateTableItemsTags is the SQL query to create a table 'items_tags'.
const sqlCreateTableItemsTags = `
	CREATE TABLE IF NOT EXISTS items_tags (
        it_id				INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id 		    INTEGER NOT NULL,
		tag_id				INTEGER NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (item_id) REFERENCES items(item_id),
		FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
	)
`

func (sdb *storageDB) createTable(tableDefinition string) error {
	_, err := sdb.sTX.Exec(tableDefinition)
	if err != nil {
		return formError(BSERR00002DbTableCreationFailed, tableDefinition, err.Error())
	}
	return nil
}
