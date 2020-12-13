package queries

const (
	SaveVideoQuery = `INSERT INTO videos (KEY,VALUE) VALUES ($id, $video)`

	GetVideoByIdQuery = `SELECT * FROM videos WHERE meta().id = $id`

	GetVideoByTitleQuery = `SELECT * FROM videos WHERE videos.title = $title`
)
