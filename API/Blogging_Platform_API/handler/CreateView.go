package handler

import (
	Blogging_Platform_API "Blogging_Platform_API/database"
	"Blogging_Platform_API/utility"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CreateViewAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db, err := sql.Open("mysql", utility.DNS())
	if err != nil {
		panic(err)
	}
	queries := Blogging_Platform_API.New(db)
	if r.Method == http.MethodGet {
		fmt.Println("Retrieving Posts ...")

		if term := r.URL.Query().Get("term"); term != "" {
			fmt.Println("Retrieving Filtered ...")
			params := Blogging_Platform_API.TermBlogSearchParams{
				CONCAT:   term,
				CONCAT_2: term,
				CONCAT_3: term,
			}
			filteredBlogs, err := queries.TermBlogSearch(ctx, params)
			if err != nil {
				panic(err)
			}
			if len(filteredBlogs) == 0 {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			jsonified, err := json.Marshal(filteredBlogs)
			if err != nil {
				panic(err)
			}

			fmt.Println("Filtered Retrieved ...")

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonified)
			return
		}

		i, err := queries.ViewAllBlog(ctx)
		if err != nil {
			panic(err)
		}

		jsonified, err := json.Marshal(i)
		if err != nil {
			panic(err)
		}

		fmt.Println("Posts retrieved ...")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonified)
	}
	if r.Method == http.MethodPost {
		fmt.Println("Creating Post ...")
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		content := r.FormValue("content")
		category := r.FormValue("category")
		tags := r.FormValue("tags")

		tagsRawMessage := json.RawMessage([]byte(tags))

		args := Blogging_Platform_API.CreateBlogParams{
			Title:     r.FormValue("title"),
			Content:   content,
			Category:  category,
			Tags:      tagsRawMessage,
			Createdat: time.Now(),
			Updatedat: time.Now(),
		}

		res, err := queries.CreateBlog(ctx, args)
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		i, err := queries.ViewSingleBlog(ctx, id)
		if err != nil {
			panic(err)
		}

		jsonified, err := json.Marshal(i)
		if err != nil {
			panic(err)
		}

		fmt.Println("Post created ...")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonified)
	}
}
