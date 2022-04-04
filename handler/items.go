package handler

import (
    "context"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "gitlab.com/idoko/bucketeer/db"
    "gitlab.com/idoko/bucketeer/models"
)

//var itemIDKey = "itemID"
func items(router chi.Router) {
    router.Get("/", GetAllValidItems)          // MUST REDIRECT TO MAIN PAGE
    router.Route("/item", func(router chi.Router) {
        //router.Use(ItemContext)
        router.Get("/", GetAllValidItems)
		router.Get("/", GetAllRejectedItems)
		router.Route("")
        router.Put("/", updateItem)
        router.Delete("/", deleteItem)
    })
}

/*
func ItemContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        itemId := chi.URLParam(r, "itemId")
        if itemId == "" {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
            return
        }
        id, err := strconv.Atoi(itemId)
        if err != nil {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
        }
        ctx := context.WithValue(r.Context(), itemIDKey, id)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}*/