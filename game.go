package igdb

import (
	"encoding/json"
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

type GameWrapper struct {
	Game
}

//go:generate gomodifytags -file $GOFILE -struct Game -add-tags json -w

// Game contains information on an IGDB entry for a particular video game.
// For more information visit: https://api-docs.igdb.com/#game
type Game struct {
	ID                    int                        `json:"id"`
	AgeRatings            []AgeRatingWrapper         `json:"age_ratings"`
	AggregatedRating      float64                    `json:"aggregated_rating"`
	AggregatedRatingCount int                        `json:"aggregated_rating_count"`
	AlternativeNames      []AlternativeNameWrapper   `json:"alternative_names"`
	Artworks              []ArtworkWrapper           `json:"artworks"`
	Bundles               []GameWrapper              `json:"bundles"`
	Category              GameCategory               `json:"category"`
	Collection            CollectionWrapper          `json:"collection"`
	Cover                 CoverWrapper               `json:"cover"`
	CreatedAt             int                        `json:"created_at"`
	DLCS                  []GameWrapper              `json:"dlcs"`
	ExpandedGames         []GameWrapper              `json:"expanded_games"`
	Expansions            []GameWrapper              `json:"expansions"`
	ExternalGames         []ExternalGameWrapper      `json:"external_games"`
	FirstReleaseDate      int                        `json:"first_release_date"`
	Follows               int                        `json:"follows"`
	Forks                 []GameWrapper              `json:"forks"`
	Franchise             FranchiseWrapper           `json:"franchise"`
	Franchises            []FranchiseWrapper         `json:"franchises"`
	GameEngines           []GameEngineWrapper        `json:"game_engines"`
	GameModes             []GameModeWrapper          `json:"game_modes"`
	Genres                []GenreWrapper             `json:"genres"`
	Hypes                 int                        `json:"hypes"`
	InvolvedCompanies     []InvolvedCompanyWrapper   `json:"involved_companies"`
	Keywords              []KeywordWrapper           `json:"keywords"`
	MultiplayerModes      []MultiplayerModeWrapper   `json:"multiplayer_modes"`
	Name                  string                     `json:"name"`
	ParentGame            *GameWrapper               `json:"parent_game"`
	Platforms             []PlatformWrapper          `json:"platforms"`
	PlayerPerspectives    []PlayerPerspectiveWrapper `json:"player_perspectives"`
	Ports                 []GameWrapper              `json:"ports"`
	Rating                float64                    `json:"rating"`
	RatingCount           int                        `json:"rating_count"`
	ReleaseDates          []ReleaseDateWrapper       `json:"release_dates"`
	Remakes               []GameWrapper              `json:"remakes"`
	Remasters             []GameWrapper              `json:"remasters"`
	Screenshots           []ScreenshotWrapper        `json:"screenshots"`
	SimilarGames          []GameWrapper              `json:"similar_games"`
	Slug                  string                     `json:"slug"`
	StandaloneExpansions  []GameWrapper              `json:"standalone_expansions"`
	Status                GameStatus                 `json:"status"`
	Storyline             string                     `json:"storyline"`
	Summary               string                     `json:"summary"`
	Tags                  []Tag                      `json:"tags"`
	Themes                []ThemeWrapper             `json:"themes"`
	TotalRating           float64                    `json:"total_rating"`
	TotalRatingCount      int                        `json:"total_rating_count"`
	UpdatedAt             int                        `json:"updated_at"`
	URL                   string                     `json:"url"`
	VersionParent         *GameWrapper               `json:"version_parent"`
	VersionTitle          string                     `json:"version_title"`
	Videos                []GameVideoWrapper         `json:"videos"`
	Websites              []WebsiteWrapper           `json:"websites"`
}

// GameCategory specifies a type of game content.
type GameCategory int

//go:generate stringer -type=GameCategory,GameStatus

// Expected GameCategory enums from the IGDB.
const (
	MainGame GameCategory = iota
	DLCAddon
	Expansion
	Bundle
	StandaloneExpansion
	Mod
	Episode
	Season
	Remake
	Remaster
	ExpandedGame
	Port
	Fork
)

// GameStatus specifies the release status of a specific game.
type GameStatus int

// Expected GameStatus enums from the IGDB.
const (
	StatusReleased GameStatus = iota
	_
	StatusAlpha
	StatusBeta
	StatusEarlyAccess
	StatusOffline
	StatusCancelled
	StatusRumored
	StatusDelisted
)

// GameService handles all the API
// calls for the IGDB Game endpoint.
type GameService service

// Get returns a single Game identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Games, an error is returned.
func (gs *GameService) Get(id int, opts ...Option) (*Game, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var g []*Game

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.post(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Game with ID %v", id)
	}

	return g[0], nil
}

// List returns a list of Games identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Game is ignored. If none of the IDs
// match a Game, an error is returned.
func (gs *GameService) List(ids []int, opts ...Option) ([]*Game, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var g []*Game

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := gs.client.post(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Games with IDs %v", ids)
	}

	return g, nil
}

// Index returns an index of Games based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Games can
// be found using the provided options, an error is returned.
func (gs *GameService) Index(opts ...Option) ([]*Game, error) {
	var g []*Game

	err := gs.client.post(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Games")
	}

	return g, nil
}

// Search returns a list of Games found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Games are found using the provided query, an error is returned.
func (gs *GameService) Search(qry string, opts ...Option) ([]*Game, error) {
	var g []*Game

	opts = append(opts, setSearch(qry))
	err := gs.client.post(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Game with query %s", qry)
	}

	return g, nil
}

// Count returns the number of Games available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Games to count.
func (gs *GameService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Games")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Game object.
func (gs *GameService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Game fields")
	}

	return f, nil
}

func (g *GameWrapper) UnmarshalJSON(data []byte) error {
	if id, err := strconv.Atoi(string(data)); err == nil {
		g.ID = id
		return nil
	}
	return json.Unmarshal(data, &g.Game)
}
