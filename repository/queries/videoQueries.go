package queries

const (
	SaveVideoQuery = `INSERT INTO videos (KEY,VALUE) VALUES ($id, $video)`

	GetVideoByIdQuery = `SELECT videos.* FROM videos WHERE meta().id = $id`

	GetVideoByTitleQuery = `SELECT videos.* FROM videos WHERE title = $title`

	GetVideoIdQuery = `SELECT videos.id FROM videos WHERE title = $1 `

	GetVideoCategoryQuery = `SELECT videos.category FROM videos WHERE meta().id = $1 `

	GetVideoSubCategoryQuery = `SELECT videos.sub_category FROM videos WHERE meta().id = $1 `

	GetVideosByCategoryQuery = `SELECT videos.* FROM videos WHERE category= $1`

	DeleteVideoQuery = `DELETE FROM videos WHERE meta().id = $1`

	UpdateVideoQuery = `UPDATE videos SET title=$1, content_link=$2, status=$3, ` +
		`category= $4, sub_category=$5, descriptions=$6 WHERE meta().id= $7`

	GetVideoStatusQuery = `SELECT videos.status FROM videos WHERE meta().id= $1`

	GetAllVideosQuery = `SELECT * FROM videos`

	GetVideoBySubCategoryQuery = `SELECT * FROM videos WHERE videos.category= $1 AND videos.sub_category=$2`

	IsVideoExistsQuery = `SELECT meta().id FROM videos WHERE title = $1`

	GetVideoTitleQuery = `SELECT videos.title FROM videos WHERE meta().id= $1`
)
