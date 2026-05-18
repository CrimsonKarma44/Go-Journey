package handler

import (
	Blogging_Platform_API "Blogging_Platform_API/database"
	"Blogging_Platform_API/utility"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func UpdateDeleteGetPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db, err := sql.Open("mysql", utility.DNS())
	if err != nil {
		panic(err)
	}
	queries := Blogging_Platform_API.New(db)

	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		panic(err)
	}

	if r.Method == http.MethodPut {
		fmt.Println("Updating Post ...")
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		content := r.FormValue("content")
		category := r.FormValue("category")
		tags := r.FormValue("tags")

		tagsRawMessage := json.RawMessage([]byte(tags))

		args := Blogging_Platform_API.UpdateBlogParams{
			Title:     r.FormValue("title"),
			Content:   content,
			Category:  category,
			Tags:      tagsRawMessage,
			Updatedat: time.Now(),
			ID:        int64(id),
		}

		err := queries.UpdateBlog(ctx, args)
		if err != nil {
			panic(err)
		}

		i, err := queries.ViewSingleBlog(ctx, int64(id))
		if err != nil {
			panic(err)
		}

		jsonified, err := json.Marshal(i)
		if err != nil {
			panic(err)
		}

		fmt.Println("Post Updated ...")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonified)
	}
	if r.Method == http.MethodDelete {
		fmt.Println("Deleting Post ...")
		err := queries.DeleteBlog(ctx, int64(id))
		if err != nil {
			panic(err)
		}

		fmt.Println("Post Deleted ...")

		w.WriteHeader(http.StatusNoContent)
	}
	if r.Method == http.MethodGet {
		fmt.Println("Retrieving Post ...")
		i, err := queries.ViewSingleBlog(ctx, int64(id))
		if err != nil {
			panic(err)
		}

		jsonified, err := json.Marshal(i)
		if err != nil {
			panic(err)
		}

		fmt.Println("Post retrieved ...")

		w.WriteHeader(http.StatusCreated)
		w.Write(jsonified)
	}
}
