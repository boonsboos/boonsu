package osu

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Beatmap struct {
	MapsetID     int     `json:"beatmapset_id"`
	StarRating   float64 `json:"difficulty_rating"`
	MapID        int     `json:"id"`
	Status       string  `json:"status"`
	Difficulty   string  `json:"version"`
	OD           float64 `json:"accuracy"`
	AR           float64 `json:"ar"`
	BPM          float64 `json:"bpm"`
	CS           float64 `json:"cs"`
	HP           float64 `json:"drain"`
	CircleCount  int     `json:"count_circles"`
	SliderCount  int     `json:"count_sliders"`
	SpinnerCount int     `json:"count_spinners"`
	Length       int     `json:"total_length"`
	PlayLength   int     `json:"hit_length"` // also known as Drain Time
	URL          string  `json:"url"`
	MaxCombo     int     `json:"max_combo"`
}

type BeatmapSet struct {
	Artist   string `json:"artist"`
	Mapper   string `json:"creator"`
	MapsetID int    `json:"id"`
	Title    string `json:"title"`
}

func GetBeatmap(beatmapID int) (Beatmap, error) {
	req, err := http.NewRequest("GET", OsuAPIURL+"/beatmaps/"+strconv.Itoa(beatmapID), nil)
	if err != nil {
		log.Println("failed to make request to get beatmap: " + err.Error())
		return Beatmap{}, errors.New("failed to make request")
	}

	req.Header.Add("Authorization", "Bearer "+osuAuthToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("failed to complete request for beatmap data: " + err.Error())
		return Beatmap{}, errors.New("failed to read response with beatmap data")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response with beatmap: " + err.Error())
		return Beatmap{}, errors.New("failed to read response with Beatmap data")
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("getting osu beatmap ID:", strconv.Itoa(beatmapID), string(bytes))
		return Beatmap{}, errors.New("bad request")
	}

	var beatmap Beatmap
	json.Unmarshal(bytes, &beatmap)
	return beatmap, nil
}
