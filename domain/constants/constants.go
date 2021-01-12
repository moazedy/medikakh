package constants

// general constants for casbin
const (
	SaveAction   = "save"
	ReadAction   = "read"
	DeleteAction = "delete"
	UpdateAction = "update"
)

// consts related to articles
const (
	ArticleObject = "article"
)

// related to video
const (
	VideoObject = "video"
)

// related to user
const (
	UserObject = "user"
)

// related to dd
const (
	DDobject = "dd"
)

// related to category
const (
	CategoryObject = "category"
)

// related to role
const (
	BronzeUserObject = "bronze"
	SilverUserObject = "silver"
	GoldUserObject   = "gold"
	GuestUserObject  = "guest"
	SystemRoleObject = "system"
	AdminRoleObject  = "admin"
)

const (
	JwtSecretKey = "secret_key"
)
