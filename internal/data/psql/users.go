package psql

import (
	"context"
	"github.com/qdm12/go-template/internal/models"
)

//// CreateUser inserts a user in the database.
//func (db *Database) CreateUser(ctx context.Context, user models.Post) (err error) {
//	_, err = db.sql.ExecContext(ctx,
//		"INSERT INTO users(id, account, username, email) VALUES ($1,$2,$3,$4);",
//		user.ID,
//		user.Account,
//		user.Username,
//		user.Email,
//	)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// GetUserByID returns the user corresponding to a user ID from the database.
//func (db *Database) GetUserByID(ctx context.Context, id uint64) (user models.Post, err error) {
//	row := db.sql.QueryRowContext(ctx,
//		"SELECT account, email, username FROM users WHERE id = $1;",
//		id,
//	)
//	user.ID = id
//	err = row.Scan(&user.Account, &user.Email, &user.Username)
//	if errors.Is(err, sql.ErrNoRows) {
//		return user, fmt.Errorf("%w: for id %d", dataerrors.ErrUserNotFound, id)
//	} else if err != nil {
//		return user, fmt.Errorf("%w: for id %d", err, id)
//	}
//	return user, nil
//}

func (db *Database) GetPosts(ctx context.Context) ([]models.Post, error) {
	rows, err := db.sql.QueryContext(ctx,
		"SELECT id, content FROM posts;",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]models.Post, 0)
	// An album slice to hold data from returned rows.
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Content); err != nil {
			return nil, err
		}
		res = append(res, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
