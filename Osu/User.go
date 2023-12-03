package osu

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// don't need very detailed stats for showing an osu profile in discord
// lots of json is ignored
type UserExtended struct {
	AvatarURL    string         `json:"avatar_url"`
	Username     string         `json:"username"`
	ID           int            `json:"id"`
	ProfileColor string         `json:"profile_colour"`
	Country      string         `json:"country_code"`
	Statistics   UserStatistics `json:"statistics"`
}

// i love anonymous structs
type UserStatistics struct {
	Level struct {
		Current  int64 `json:"current"`
		Progress int32 `json:"progress"`
	} `json:"level"`

	PP         int64 `json:"pp"`
	GlobalRank int64 `json:"global_rank"`

	RankHighest struct {
		Rank int64 `json:"rank"`
	} `json:"rank_highest"`

	HitAccuracy float32 `json:"hit_accuracy"`
	PlayCount   int64   `json:"play_count"`

	// no clue why this is duplicated
	Rank struct {
		Global  int64 `json:"global"`
		Country int64 `json:"country"`
	} `json:"rank"`

	GradeCounts struct {
		SS  int `json:"ss"`
		SSH int `json:"ssh"`
		S   int `json:"s"`
		SH  int `json:"sh"`
		A   int `json:"a"`
	} `json:"grade_counts"`
}

func GetUserByUsername(osuUsername string) (UserExtended, error) {
	req, err := http.NewRequest("GET", OsuAPIURL+"user/"+osuUsername+"/osu?key=user", nil)
	if err != nil {
		log.Println("failed to make request to get user: " + err.Error())
		return UserExtended{}, errors.New("failed to make request")
	}

	req.Header.Add("Authorization", "Bearer "+osuAuthToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("failed to complete request for user data: " + err.Error())
		return UserExtended{}, errors.New("failed to read response with user data")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response with user data: " + err.Error())
		return UserExtended{}, errors.New("failed to read response with user data")
	}

	user := UserExtended{}
	json.Unmarshal(bytes, &user)

	log.Println(user, string(bytes))
	return user, nil
}
