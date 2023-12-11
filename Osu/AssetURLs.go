package osu

import "strconv"

func BeatmapImageLink(id int) string {
	return "https://assets.ppy.sh/beatmaps/" + strconv.Itoa(id) + "/covers/list.jpg"
}
