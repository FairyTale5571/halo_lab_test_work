package router

import (
	"net/http"

	"github.com/fairytale5571/privat_test/pkg/cache"
	"github.com/fairytale5571/privat_test/pkg/database"
	"github.com/fairytale5571/privat_test/pkg/logger"
	"github.com/fairytale5571/privat_test/pkg/models"
	"github.com/fairytale5571/privat_test/pkg/storage/redis"
	"github.com/gin-gonic/gin"
)

type Router struct {
	rout  *gin.Engine
	log   *logger.Wrapper
	rdb   *redis.Redis
	db    *database.DB
	cache *cache.Config
	cfg   models.Config
}

func New(cfg models.Config, db *database.DB, rdb *redis.Redis) *Router {
	return &Router{
		cfg:   cfg,
		rout:  gin.Default(),
		log:   logger.New("server"),
		rdb:   rdb,
		db:    db,
		cache: cache.SetupCache(rdb),
	}
}

func (r *Router) Start() {
	r.log.Infof("router started\n")
	r.mainRouter()
	if err := r.rout.Run(":" + r.cfg.Port); err != nil {
		r.log.Fatalf("cant open gin engine: %v", err)
	}
}

func (r *Router) mainRouter() {
	r.rout.GET("/film/:title", r.getFilm)
}

func (r *Router) getFilm(c *gin.Context) {
	title := c.Param("title")
	if film, err := r.cache.Get("filmTitle:" + title); film != "" && err == nil {
		c.JSON(http.StatusNotModified, film)
		r.log.Infof("get film from cache: %v", film)
		return
	}

	film, err := r.GetFilm(title)
	if film.IsEmpty() {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err != nil {
		r.log.Errorf("error get film: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, film)
	if err := r.cache.Set("filmTitle:"+title, film.String()); err != nil {
		r.log.Errorf("error cache film: %v", err)
		return
	}
}

func (r *Router) GetFilm(title string) (models.Film, error) {
	var film models.Film
	rows, err := r.db.Query(`SELECT 
		film_id, title, description, release_year, language_id, 
		rental_duration, rental_rate, length, replacement_cost, rating, last_update, special_features, fulltext
		FROM film WHERE title = $1 LIMIT 1`, title)
	defer rows.Close() //nolint directives: not need defer check error
	if err != nil {
		r.log.Errorf("GetFilm() error get film: %v\n", err)
		return film, err
	}
	for rows.Next() {
		if err := rows.Scan(
			&film.FilmID,
			&film.Title,
			&film.Description,
			&film.ReleaseYear,
			&film.LanguageID,
			&film.RentalDuration,
			&film.RentalRate,
			&film.Length,
			&film.ReplacementCost,
			&film.Rating,
			&film.LastUpdate,
			&film.SpecialFeatures,
			&film.FullText,
		); err != nil {
			return film, err
		}
	}
	return film, nil
}
