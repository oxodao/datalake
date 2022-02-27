package module_spotify

import "github.com/jmoiron/sqlx"

type ORM struct {
	db *sqlx.DB
}

func newORM(db *sqlx.DB) *ORM {
	return &ORM{db: db}
}

func (o ORM) GetArtist(name string) (DatalakeArtist, error) {
	artist := DatalakeArtist{}
	err := o.db.Get(&artist, `
		SELECT id, spotify_id, display_name
		FROM module_spotify_artist
		WHERE display_name = $1
	`, name)

	return artist, err
}

func (o ORM) UpsertArtist(name string) (string, error) {
	row := o.db.QueryRowx(`
		INSERT INTO module_spotify_artist (display_name)
		VALUES ($1)
		ON CONFLICT (display_name) DO UPDATE SET display_name = EXCLUDED.display_name
		RETURNING id
	`, name)

	if row.Err() != nil {
		return "", row.Err()
	}

	var id string
	err := row.Scan(&id)

	return id, err
}

func (o ORM) UpsertTrack(idArtist, track string) (string, error) {
	row := o.db.QueryRowx(`
		INSERT INTO module_spotify_track (artist_id, display_name)
		VALUES ($1, $2)
		ON CONFLICT (display_name) DO UPDATE SET display_name = EXCLUDED.display_name
		RETURNING id
	`, idArtist, track)

	if row.Err() != nil {
		return "", row.Err()
	}

	var id string
	err := row.Scan(&id)

	return id, err
}

func (o ORM) UpsertPlayedTrack(idTrack, playedAt string, duration int64) (string, error) {
	row := o.db.QueryRowx(`
		INSERT INTO module_spotify_played_track (track_id, played_at, spotify_played_at, duration_played)
		VALUES ($1, TO_TIMESTAMP('YYYY-MM-DD HH24:MI', $2), $2, $3)
		ON CONFLICT (track_id, spotify_played_at) DO UPDATE SET spotify_played_at = EXCLUDED.spotify_played_at
		RETURNING id
	`, idTrack, playedAt, duration)

	if row.Err() != nil {
		return "", row.Err()
	}

	var id string
	err := row.Scan(&id)

	return id, err
}
