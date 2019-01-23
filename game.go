package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Game -add-tags json -w

// GameService handles all the API
// calls for the IGDB Game endpoint.
type GameService service

// Game contains information on an IGDB entry for a particular video game.
//
// For more information, visit: https://api-docs.igdb.com/#game
type Game struct {
	AgeRatings            []int        `json:"age_ratings"`
	AggregatedRating      float64      `json:"aggregated_rating"`
	AggregatedRatingCount int          `json:"aggregated_rating_count"`
	AlternativeNames      []int        `json:"alternative_names"`
	Artworks              []int        `json:"artworks"`
	Bundles               []int        `json:"bundles"`
	Category              GameCategory `json:"category"`
	Collection            int          `json:"collection"`
	Cover                 int          `json:"cover"`
	CreatedAt             int          `json:"created_at"`
	DLCS                  []int        `json:"dlcs"`
	Expansions            []int        `json:"expansions"`
	ExternalGames         []int        `json:"external_games"`
	FirstReleaseDate      int          `json:"first_release_date"`
	Follows               int          `json:"follows"`
	Franchise             int          `json:"franchise"`
	Franchises            []int        `json:"franchises"`
	GameEngines           []int        `json:"game_engines"`
	GameModes             []int        `json:"game_modes"`
	Genres                []int        `json:"genres"`
	Hypes                 int          `json:"hypes"`
	InvolvedCompanies     []int        `json:"involved_companies"`
	Keywords              []int        `json:"keywords"`
	MultiplayerModes      []int        `json:"multiplayer_modes"`
	Name                  string       `json:"name"`
	ParentGame            int          `json:"parent_game"`
	Platforms             []int        `json:"platforms"`
	PlayerPerspectives    []int        `json:"player_perspectives"`
	Popularity            float64      `json:"popularity"`
	PulseCount            int          `json:"pulse_count"`
	Rating                float64      `json:"rating"`
	RatingCount           int          `json:"rating_count"`
	ReleaseDates          []int        `json:"release_dates"`
	Screenshots           []int        `json:"screenshots"`
	SimilarGames          []int        `json:"similar_games"`
	Slug                  string       `json:"slug"`
	StandaloneExpansions  []int        `json:"standalone_expansions"`
	Status                GameStatus   `json:"status"`
	Storyline             string       `json:"storyline"`
	Summmary              string       `json:"summmary"`
	Tags                  []Tag        `json:"tags"`
	Themes                []int        `json:"themes"`
	TimeToBeat            int          `json:"time_to_beat"`
	TotalRating           float64      `json:"total_rating"`
	TotalRatingCount      int          `json:"total_rating_count"`
	UpdatedAt             int          `json:"updated_at"`
	URL                   string       `json:"url"`
	VersionParent         int          `json:"version_parent"`
	VersionTitle          int          `json:"version_title"`
	Videos                []int        `json:"videos"`
	Websites              []int        `json:"websites"`
}

// AltName contains information on an
// alternative name for an IGDB object.
type AltName struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// CompletionTime contains the time to complete
// a particular video game. This time is measured
// in seconds.
type CompletionTime struct {
	Hastly     int `json:"hastly"`
	Normally   int `json:"normally"`
	Completely int `json:"completely"`
}

// ESRB contains the rating and synopsis
// for a particular video game given by
// the Entertainment Software Rating Board.
type ESRB struct {
	Rating   ESRBCode `json:"rating"`
	Synopsis string   `json:"synopsis"`
}

// External contains information for
// connecting external service IDs to
// the IGDB for a particular object.
type External struct {
	Steam string `json:"steam"`
}

// PEGI contains the rating and synopsis
// for a particular video game given by
// the Pan European Game Information organization.
type PEGI struct {
	Rating   PEGICode `json:"rating"`
	Synopsis string   `json:"synopsis"`
}

// YoutubeVideo contains the name and
// ID of a  Youtube video.
type YoutubeVideo struct {
	Name string `json:"name"`
	ID   string `json:"video_id"` // Youtube slug
}

// Website contains address and category
// information on a website referenced
// in the IGDB.
type Website struct {
	Category WebsiteCategory `json:"category"`
	URL      URL             `json:"url"`
}

// Get returns a single Game identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Games, an error is returned.
func (gs *GameService) Get(id int, opts ...FuncOption) (*Game, error) {
	var g []*Game

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(GameEndpoint, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Game with ID %v", id)
	}

	return g[0], nil
}

// List returns a list of Games identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of Games based solely on the provided
// options. Any ID that does not match a Game is ignored. If none of the IDs
// match a Game, an error is returned.
func (gs *GameService) List(ids []int, opts ...FuncOption) ([]*Game, error) {
	var g []*Game

	err := gs.client.get(GameEndpoint, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot list Games with IDs %v", ids)
	}

	return g, nil
}

// Search returns a list of Games found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Games are found using the provided query, an error is returned.
//TODO: remember that Search also has its own endpoint
func (gs *GameService) Search(qry string, opts ...FuncOption) ([]*Game, error) {
	var g []*Game

	opts = append(opts, setSearch(qry))
	err := gs.client.get(GameEndpoint, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot search for Game with query %s", qry)
	}

	return g, nil
}

// Count returns the number of Games available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Games to count.
func (gs *GameService) Count(opts ...FuncOption) (int, error) {
	ct, err := gs.client.getEndpointCount(GameEndpoint, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Games")
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Game object.
func (gs *GameService) ListFields() ([]string, error) {
	fl, err := gs.client.getEndpointFieldList(GameEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "cannot list Game fields")
	}

	return fl, nil
}
