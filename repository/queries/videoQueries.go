package queries

const (
	SaveVideoQuery = `INSERT INTO videos (KEY,VALUE) VALUES ($id, $video)`

	GetVideoByIdQuery = `SELECT * FROM videos WHERE meta().id = $id`

	GetVideoByTitleQuery = `SELECT videos.* FROM videos WHERE title = $title`

	GetVideoIdQuery = `SELECT videos.id FROM videos WHERE title = $1 `

	GetVideoCategoryQuery = `SELECT videos.category FROM videos WHERE meta().id = $1 `

	GetVideoSubCategoryQuery = `SELECT videos.sub_category FROM videos WHERE meta().id = $1 `

	GetVideosByCategoryQuery = `SELECT videos.* FROM videos WHERE category= $1`

	DeleteVideoQuery = `DELETE FROM videos WHERE meta().id = $1`

	UpdateVideoQuery = `UPDATE videos SET title=$title, content_link=$contentlink, status=$status, ` +
		`category= $category, sub_category=$subcategory, descriptions=$descriptions WHERE meta().id= $id`

	GetVideoStatusQuery = `SELECT videos.status FROM videos WHERE meta().id= $1`

	GetAllVideosQuery = `SELECT * FROM videos`

	GetVideoBySubCategoryQuery = `SELECT * FROM videos WHERE videos.category= $1 AND videos.sub_category=$2`

	IsVideoExistsQuery = `SELECT meta().id FROM videos WHERE title = $1`

	GetVideoTitleQuery = `SELECT videos.title FROM videos WHERE meta().id= $1`
)
