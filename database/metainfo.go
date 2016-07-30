package database

import ()

type Tablename string
type Rowname string

//Названия таблиц
const (
	TableUser Tablename = "forumuser"
	TableForum Tablename = "forum"
	TablePost Tablename = "post"
	TableThread Tablename = "thread"
	TableFollow Tablename = "follow"
)

//Названия строк юзера
const (
	User_id Rowname = "id"
	User_name Rowname = "name"
	User_username Rowname = "username"
	User_email Rowname = "email"
	User_isAnonymous Rowname = "isanonymous"
)

//Названия строк форума
const (
	Forum_ID Rowname = "id"
	Forum_name Rowname = "name"
	Forum_shortname Rowname = "short_name"
	Forum_user_ Rowname = "user"
)

//Названия строк подписчика
const (
	Follow_follower_ Rowname = "follower"
	Follow_following_ Rowname = "following"
)

//Названия строк треда
const (
	Thread_id Rowname = "id"
	Thread_date Rowname = "date"
	Thread_dislikes Rowname = "dislikes"
	Thread_likes Rowname = "likes"
	Thread_slug Rowname = "slug"
	Thread_title Rowname = "title"
	Thread_forum_ Rowname = "forum"
	Thread_message Rowname = "message"
	Thread_isClosed Rowname = "isclosed"
	Thread_isDeleted Rowname = "isdeleted"
	Thread_user_ Rowname = "user"
)

type Forum struct {
	id          int
	name        string
	short_name  string
	email       string
	about       string
	isAnonymous bool
}

type Post struct {
	id          int
	date        string
	dislikes    int
	isApproved  bool
	about       string
	isAnonymous bool
}