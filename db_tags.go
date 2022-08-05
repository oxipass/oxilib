package oxilib

import (
	"database/sql"
	"sort"
	"time"
)

const sqlInsertTag = `
	INSERT INTO tags (extid, name,color,created,updated,deleted) 
		VALUES (?, ?,?,?,?,0)
`

func (sdb *storageDB) dbInsertTag(tagName string, color string, extid string) (tagId int64, err error) {
	if sdb.sTX == nil {
		return 0, formError(BSERR00003DbTransactionFailed, "dbInsertTag")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlInsertTag)
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertTag")
	}
	res, errStmt := stmt.Exec(extid,
		tagName,
		color,
		creationTime,
		creationTime)

	if errStmt != nil {
		return 0, formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertTag")
	}
	tagId, err = res.LastInsertId()
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertTag")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertTag")
	}

	return tagId, nil
}

const sqlAssignTagToItem = `
	INSERT INTO items_tags (item_id, tag_id, created, updated, deleted) 
		VALUES (?,?,?,?,0)
`

func (sdb *storageDB) dbAssignTag(tagId int64, itemId int64) (itId int64, err error) {
	if sdb.sTX == nil {
		return 0, formError(BSERR00003DbTransactionFailed, "dbAssignTag")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlAssignTagToItem)
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbAssignTag")
	}
	res, errStmt := stmt.Exec(itemId, tagId,
		creationTime, creationTime)

	if errStmt != nil {
		return 0, formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbAssignTag")
	}
	itId, err = res.LastInsertId()
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbAssignTag")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, formError(BSERR00006DbInsertFailed, errClose.Error(), "dbAssignTag")
	}

	return itId, nil
}

// sqlListItemTags - List all non-deleted items
const sqlListItemTags = `
	SELECT tags.tag_id, tags.name, tags.color, tags.extid, tags.created, tags.updated, tags.deleted
		FROM tags 
		INNER JOIN items_tags it on tags.tag_id = it.tag_id
		WHERE tags.deleted='0' 
		  AND it.deleted='0'
		  AND it.item_id=?
`

// sqlListTags - List all available tags (excluding deleted)
const sqlListTags = `
	SELECT tag_id, name, color, extid, created, updated, deleted
		FROM tags 
		WHERE tags.deleted='0' 
`

// dbSelectItemTags - select tags assigned to requested the item
func (sdb *storageDB) dbSelectItemTags(itemId int64) (tags []OxiTag, err error) {
	var rows *sql.Rows

	if itemId == -1 {
		rows, err = sdb.sDB.Query(sqlListTags)
	} else {
		rows, err = sdb.sDB.Query(sqlListItemTags, itemId)
	}
	if err != nil {
		return tags, formError(BSERR00021FieldsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = formError(BSERR00021FieldsReadFailed, err.Error(), errClose.Error())
		}
	}()

	var tag OxiTag

	for rows.Next() {
		err = rows.Scan(&tag.ID,
			&tag.Name,
			&tag.Color,
			&tag.ExtId,
			&tag.Created,
			&tag.Updated,
			&tag.Deleted)
		if err != nil {
			return tags, err
		}

		tags = append(tags, tag)
	}
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	}) // Sort by name before returning

	return tags, nil
}

// dbSelectTags - select all available tags
func (sdb *storageDB) dbSelectTags() (tags []OxiTag, err error) {
	return sdb.dbSelectItemTags(-1)
}
