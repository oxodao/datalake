package module_spotify

type SpotifyStream struct {
	EndTime        string `json:"endtime"`
	ArtistName     string `json:"artistName"`
	TrackName      string `json:"trackName"`
	DurationPlayed int64  `json:"msPlayed"`
}

type DatalakeArtist struct {
	Id        string `json:"id"`
	SpotifyId string `json:"spotify_id"`
	Name      string `json:"display_name"`
}
