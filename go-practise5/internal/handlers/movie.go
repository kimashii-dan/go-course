package handlers

import (
	"encoding/json"
	"go-practise5/internal/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

type MovieResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Year       int    `json:"year"`
	ActorCount int64  `json:"actor_count"`
}

func GetMovies(writer http.ResponseWriter, request *http.Request) {
	queryURL := request.URL.Query()
	start := time.Now()

	yearMinStr := queryURL.Get("year_min")
	yearMaxStr := queryURL.Get("year_max")
	limitStr := queryURL.Get("limit")
	offsetStr := queryURL.Get("offset")

	var yearMin, yearMax, limit, offset int

	if yearMinStr != "" {
		yearMin, _ = strconv.Atoi(yearMinStr)
	}
	if yearMaxStr != "" {
		yearMax, _ = strconv.Atoi(yearMaxStr)
	}
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}
	if offsetStr != "" {
		offset, _ = strconv.Atoi(offsetStr)
	}

	var movies []MovieResponse
	query := db.DB.Table("movies AS m").
		Select("m.id, m.title, m.year, COUNT(a.id) AS actor_count").
		Joins("LEFT JOIN actors a ON a.movie_id = m.id").
		Group("m.id").
		Order("m.year DESC")

	if yearMin != 0 {
		query = query.Where("m.year >= ?", yearMin)
	}
	if yearMax != 0 {
		query = query.Where("m.year <= ?", yearMax)
	}
	if limit != 0 {
		query = query.Limit(limit)
	}
	if offset != 0 {
		query = query.Offset(offset)
	}

	if err := query.Scan(&movies).Error; err != nil {
		http.Error(writer, "Database query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	duration := time.Since(start)
	log.Printf("Query took %v", duration)
	writer.Header().Set("X-Query-Time", duration.String())
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(movies)
}
