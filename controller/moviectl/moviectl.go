package moviectl

import (
	"context"
	"net/http"
	"time"

	"github.com/gilcrest/go-api-basic/app"
	"github.com/gilcrest/go-api-basic/controller"
	"github.com/gilcrest/go-api-basic/datastore/movieds"
	"github.com/gilcrest/go-api-basic/domain/errs"
	"github.com/gilcrest/go-api-basic/domain/movie"
	"github.com/rs/xid"
)

// MovieController is used as the base controller for the Movie logic
type MovieController struct {
	App *app.Application
	SRF controller.StandardResponseFields
}

// MovieRequest is the request struct
type MovieRequest struct {
	Title    string `json:"Title"`
	Year     int    `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"ReleaseDate"`
	RunTime  int    `json:"RunTime"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
}

// MovieResponse is the response struct for a single Movie
type MovieResponse struct {
	ExtlID          xid.ID `json:"ExtlID"`
	Title           string `json:"Title"`
	Year            int    `json:"Year"`
	Rated           string `json:"Rated"`
	Released        string `json:"ReleaseDate"`
	RunTime         int    `json:"RunTime"`
	Director        string `json:"Director"`
	Writer          string `json:"Writer"`
	CreateTimestamp string `json:"CreateTimestamp"`
}

// ListMovieResponse is the response struct for multiple Movies
type ListMovieResponse struct {
	controller.StandardResponseFields
	Data []*MovieResponse `json:"data"`
}

// SingleMovieResponse is the response struct for multiple Movies
type SingleMovieResponse struct {
	controller.StandardResponseFields
	Data *MovieResponse `json:"data"`
}

// provideMovieResponse is an initializer for MovieResponse
func (ctl *MovieController) provideMovieResponse(m *movie.Movie) *MovieResponse {
	return &MovieResponse{
		ExtlID:          m.ExtlID,
		Title:           m.Title,
		Year:            m.Year,
		Rated:           m.Rated,
		Released:        m.Released.Format(time.RFC3339),
		RunTime:         m.RunTime,
		Director:        m.Director,
		Writer:          m.Writer,
		CreateTimestamp: m.CreateTimestamp.Format(time.RFC3339),
	}
}

// NewMovieController initializes MovieController
func NewMovieController(app *app.Application, srf controller.StandardResponseFields) *MovieController {
	return &MovieController{App: app, SRF: srf}
}

// Add adds a movie to the catalog.
func (ctl *MovieController) Add(ctx context.Context, r *MovieRequest) (*MovieResponse, error) {
	const op errs.Op = "controller/moviectl/MovieController.Add"

	err := ctl.App.DS.BeginTx(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}

	mds, err := movieds.NewMovieDS(ctl.App)
	if err != nil {
		return nil, errs.E(op, err)
	}

	m, err := provideMovie(r)
	if err != nil {
		return nil, errs.E(op, err)
	}

	err = m.Add(ctx)
	if err != nil {
		return nil, errs.E(err)
	}

	err = mds.Store(ctx, m)
	if err != nil {
		return nil, errs.E(op, errs.Database, ctl.App.DS.RollbackTx(err))
	}

	resp := ctl.provideMovieResponse(m)

	if err := ctl.App.DS.CommitTx(); err != nil {
		return nil, errs.E(op, errs.Database, err)
	}

	return resp, nil
}

// Update updates the movie given the id sent in
func (ctl *MovieController) Update(ctx context.Context, id string, r *MovieRequest) (*MovieResponse, error) {
	const op errs.Op = "controller/moviectl/MovieController.Update"

	err := ctl.App.DS.BeginTx(ctx)
	if err != nil {
		return nil, errs.E(op, err)
	}

	mds, err := movieds.NewMovieDS(ctl.App)
	if err != nil {
		return nil, errs.E(op, err)
	}

	m, err := provideMovie(r)
	if err != nil {
		return nil, errs.E(op, err)
	}

	err = m.Add(ctx)
	if err != nil {
		return nil, errs.E(err)
	}

	err = mds.Store(ctx, m)
	if err != nil {
		return nil, errs.E(op, errs.Database, ctl.App.DS.RollbackTx(err))
	}

	resp := ctl.provideMovieResponse(m)

	if err := ctl.App.DS.CommitTx(); err != nil {
		return nil, errs.E(op, errs.Database, err)
	}

	return resp, nil
}

// FindByID finds a movie given its' unique ID
func (ctl *MovieController) FindByID(ctx context.Context, id string) (*SingleMovieResponse, error) {
	const op errs.Op = "controller/moviectl/FindByID"

	mds, err := movieds.NewMovieDS(ctl.App)
	if err != nil {
		return nil, errs.E(op, err)
	}

	i, err := xid.FromString(id)
	if err != nil {
		return nil, errs.E(op, errs.Validation, "Invalid id in URL path")
	}

	m, err := mds.FindByID(ctx, i)
	if err != nil {
		return nil, errs.E(op, errs.Database, ctl.App.DS.RollbackTx(err))
	}

	mr := ctl.provideMovieResponse(m)

	response := ctl.NewSingleMovieResponse(mr)

	return response, nil
}

// FindAll finds the entire set of Movies
func (ctl *MovieController) FindAll(ctx context.Context, r *http.Request) (*ListMovieResponse, error) {
	const op errs.Op = "controller/moviectl/FindByID"

	mds, err := movieds.NewMovieDS(ctl.App)
	if err != nil {
		return nil, errs.E(op, err)
	}

	ms, err := mds.FindAll(ctx)
	if err != nil {
		return nil, errs.E(op, errs.Database, ctl.App.DS.RollbackTx(err))
	}

	response := ctl.NewListMovieResponse(ms, r)

	return response, nil
}

// NewListMovieResponse is an initializer for ListMovieResponse
func (ctl *MovieController) NewListMovieResponse(ms []*movie.Movie, r *http.Request) *ListMovieResponse {
	const op errs.Op = "controller/moviectl/NewListMovieResponse"

	var s []*MovieResponse

	for _, m := range ms {
		mr := ctl.provideMovieResponse(m)
		s = append(s, mr)
	}

	return &ListMovieResponse{StandardResponseFields: ctl.SRF, Data: s}
}

// NewSingleMovieResponse is an initializer for SingleMovieResponse
func (ctl *MovieController) NewSingleMovieResponse(mr *MovieResponse) *SingleMovieResponse {
	const op errs.Op = "controller/moviectl/NewSingleMovieResponse"

	return &SingleMovieResponse{StandardResponseFields: ctl.SRF, Data: mr}
}

// dateFormat is the expected date format for any date fields
// in the request
const dateFormat string = "Jan 02 2006"

// NewMovie is an initializer for the Movie struct
func provideMovie(am *MovieRequest) (*movie.Movie, error) {
	const op errs.Op = "controller/moviectl/NewMovie"

	t, err := time.Parse(dateFormat, am.Released)
	if err != nil {
		return nil, errs.E(op,
			errs.Validation,
			errs.Code("invalid_date_format"),
			errs.Parameter("ReleaseDate"),
			err)
	}

	return &movie.Movie{
		Title:    am.Title,
		Year:     am.Year,
		Rated:    am.Rated,
		Released: t,
		RunTime:  am.RunTime,
		Director: am.Director,
		Writer:   am.Writer,
	}, nil
}
