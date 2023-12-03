package osu

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type BeatmapDifficultyAttributes struct {
	MaxCombo             int64   `json:"max_combo"`
	StarRating           float64 `json:"star_rating"`
	AimDifficulty        float64 `json:"aim_difficulty"`
	ApproachRate         float64 `json:"approach_rate"`
	FlashlightDifficulty float64 `json:"flashlight_difficulty"`
	OverallDifficulty    float64 `json:"overall_difficulty"`
	SliderFactor         float64 `json:"slider_factor"`
	SpeedDifficulty      float64 `json:"speed_difficulty"`
	SpeedNoteCount       int64   `json:"speed_note_count"`
}

func GetBeatmapAttributes(beatmapID string) BeatmapDifficultyAttributes {
	req, err := http.NewRequest("GET", OsuAPIURL+"/beatmaps/"+beatmapID+"/attributes", nil)
	if err != nil {
		log.Fatal("failed to make request to get beatmap attributes: " + err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+osuAuthToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to complete request for beatmap attributes: " + err.Error())
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("failed to read response with beatmap attributes: " + err.Error())
	}

	a := struct { // temporary wrapper struct
		Attributes BeatmapDifficultyAttributes `json:"attributes"`
	}{BeatmapDifficultyAttributes{}} // containing BDA

	json.Unmarshal(bytes, &a)
	return a.Attributes
}
