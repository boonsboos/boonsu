package osu

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Score struct {
	Accuracy  float64  `json:"accuracy"`
	MaxCombo  int      `json:"max_combo"`
	Mods      []string `json:"mods"`
	FullCombo bool     `json:"perfect"`
	PP        float64  `json:"pp"`
	Rank      string   `json:"rank"`
	HitInfo   struct {
		Count100  int `json:"count_100"`
		Count300  int `json:"count_300"`
		Count50   int `json:"count_50"`
		CountMiss int `json:"count_miss"`
	} `json:"statistics"`
	Map    Beatmap      `json:"beatmap"`
	Set    BeatmapSet   `json:"beatmapset"`
	Player UserExtended `json:"user"`
}

func GetMostRecentScoreFromUser(osuID int) (Score, error) {
	req, err := http.NewRequest("GET", OsuAPIURL+"/users/"+strconv.Itoa(osuID)+"/scores/recent?mode=osu&limit=1&include_fails=1", nil)
	if err != nil {
		log.Println("failed to make request to get user: " + err.Error())
		return Score{}, errors.New("failed to make request")
	}

	req.Header.Add("Authorization", "Bearer "+osuAuthToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("failed to complete request for user data: " + err.Error())
		return Score{}, errors.New("failed to read response with user data")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response with user data: " + err.Error())
		return Score{}, errors.New("failed to read response with score data")
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("getting osu score for:", strconv.Itoa(osuID), string(bytes))
		return Score{}, errors.New("bad request")
	}

	var recent []Score = make([]Score, 1)
	json.Unmarshal(bytes, &recent)
	return recent[0], nil
}

func GetBestScoreFromUser(osuID int) (Score, error) {
	req, err := http.NewRequest("GET", OsuAPIURL+"/users/"+strconv.Itoa(osuID)+"/scores/best?mode=osu&limit=1&include_fails=1", nil)
	if err != nil {
		log.Println("failed to make request to get score: " + err.Error())
		return Score{}, errors.New("failed to make request")
	}

	req.Header.Add("Authorization", "Bearer "+osuAuthToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("failed to complete request for score: " + err.Error())
		return Score{}, errors.New("failed to read response with user data")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response with score: " + err.Error())
		return Score{}, errors.New("failed to read response with score data")
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("getting osu score for:", strconv.Itoa(osuID), string(bytes))
		return Score{}, errors.New("bad request")
	}

	var recent []Score = make([]Score, 1)
	json.Unmarshal(bytes, &recent)
	return recent[0], nil
}

func (s Score) GetMapInfo() string {
	return s.Set.Title + " [" + s.Map.Difficulty + "] " + strconv.FormatFloat(s.Map.StarRating, 'f', 2, 64) + `★`
}

func (s Score) GetHitInfo() string {
	return strconv.Itoa(s.HitInfo.Count300) + "/" +
		strconv.Itoa(s.HitInfo.Count100) + "/" +
		strconv.Itoa(s.HitInfo.Count50) + "/" +
		strconv.Itoa(s.HitInfo.CountMiss)
}
