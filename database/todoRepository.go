package database

import (
	"strings"

	"github.com/someDevDude/todo-server/models"
	"github.com/someDevDude/todo-server/util"
)

//QueryTodos queires todos
func QueryTodos(params models.ListParams) []models.TodoFull {
	var results []models.TodoFull

	queryString := "SELECT * FROM todo"
	separator := " WHERE"
	args := make(map[string]interface{})

	if params.Q != "" {
		queryString += separator + " LOWER(title) LIKE LOWER(CONCAT('%', :query, '%'))"
		separator = " AND"
		args["query"] = params.Q
	}

	if params.StartIndex != 0 {
		queryString += separator + " OFFSET :startIndex"
		separator = " AND"
		args["startIndex"] = params.StartIndex
	}

	if params.Sort != "" {
		queryString += " ORDER BY :sortColumn "
		sortParams := strings.SplitAfter(params.Sort, ":")
		args["sortColumn"] = sortParams[0][:len(sortParams[0])-1] + "  " + sortParams[1]
	}

	if params.MaxResults != 0 {
		queryString += " LIMIT :maxResults"
		args["maxResults"] = params.MaxResults
	}

	rows, err := DB.Query(queryString)
	util.CheckErr(err, func(err error) {
		util.Errorf("Error querying todos in database\n%s", err.Error())
		return
	})

	for rows.Next() {
		var r models.TodoFull

		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Done)

		results = append(results, r)
	}

	return results
}

//CreateTodo creates a todo
func CreateTodo(todo models.TodoFull) {
	stmt, err := DB.Prepare("INSERT todo SET title = ?, dewscription = ?, done = 0")
	util.CheckErr(err, func(err error) {
		util.Errorf("Error preparing insert statement for todos\n%s", err.Error())
		return
	})

	_, err = stmt.Exec(todo.Title, todo.Description)
	util.CheckErr(err, func(err error) {
		util.Errorf("Error inseting todo into database\n%s", err.Error())
		return
	})
}
