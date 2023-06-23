package dbRequests

const (
	// Forum
	FORUM_EXIST       = `SELECT EXISTING(SELECT ) `
	CREATE_FORUM      = `INSERT INTO forums (title, user_, slug) values ($1, $2, $3);`
	GET_FORUM_BY_SLUG = `SELECT title, user_, slug, posts, threads FROM forums WHERE slug = $1`

	DEFAULT_GET_USERS   = `SELECT U.nickname, U.fullname, U.about, U.email FROM users as U `
	DEFAULT_GET_THREADS = `SELECT id, title, author, forum, message, votes, slug, created FROM threads where forum = $1`

	// Posts
	GET_POST_BY_ID = `SELECT id, COALESCE(parent, 0), author, message, is_edited, forum, thread, created FROM posts WHERE id = $1`
	UPDATE_POST    = `UPDATE posts SET message = $1, is_edited = $2 WHERE id = $3;`

	// Service
	CLEAR_ALL  = `TRUNCATE TABLE forums, posts, threads, user_forum, users, votes CASCADE;`
	GET_STATUS = `SELECT (SELECT count(*) FROM users) AS users, ` +
		`(SELECT count(*) FROM forums) AS forums, ` +
		`(SELECT count(*) FROM threads) AS threads, ` +
		`(SELECT count(*) FROM posts) AS posts;`

	// Threads
	INSERT_IN_THREAD         = `INSERT INTO threads (title, author, forum, message, slug, created) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created;`
	GET_THREAD_BY_ID         = `SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE id = $1`
	GET_THREAD_BY_SLUG       = `SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE slug = $1;`
	GET_THREAD_BY_SLUG_OR_ID = `SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE id = $1 OR slug = $2;`

	GET_VOTES      = `SELECT votes FROM threads WHERE id = $1;`
	UPDATE_THREADS = `UPDATE threads SET title = $1, message = $2 WHERE id = $3;`

	CREATE_PART_POSTS   = `INSERT INTO posts (parent, author, message, forum, thread, created) VALUES `
	GET_POST_TREE       = `SELECT id, COALESCE(parent, 0), author, message, is_edited, forum, thread, created FROM posts WHERE thread = $1`
	DEFAULT_FLAT_TREE   = `SELECT id, COALESCE(parent, 0), author, message, is_edited, forum, thread, created FROM posts WHERE thread = $1`
	DEFAULT_PARENT_TREE = `SELECT id, COALESCE(parent, 0), author, message, is_edited, forum, thread, created FROM posts WHERE path[1] IN `

	// USER
	CREATE_USER = `INSERT INTO users VALUES ($1, $2, $3, $4);`
	UPDATE_USER = "UPDATE users SET " +
		"fullname = COALESCE(NULLIF(TRIM($1), ''), fullname), " +
		"about = COALESCE(NULLIF(TRIM($2), ''), about), " +
		"email = COALESCE(NULLIF(TRIM($3), ''), email) " +
		"WHERE nickname = $4 RETURNING fullname, about, email;"
	GET_USER_BY_NICKNAME = `SELECT nickname, fullname, about, email FROM users WHERE nickname = $1;`
	GET_MATCHED_USERS    = `SELECT nickname, fullname, about, email FROM users WHERE nickname = $1 OR email = $2;`

	// VOTE
	CREATE_VOTE = `INSERT INTO votes (nickname, thread, voice) VALUES ($1, $2, $3) ON CONFLICT (nickname, thread) DO UPDATE SET voice = EXCLUDED.voice;`
)
