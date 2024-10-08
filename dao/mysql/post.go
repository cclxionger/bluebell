package mysql

import (
	"bluebell/global"
	"bluebell/models"
	"database/sql"
	"fmt"
	"strings"
)

func CreatePost(p *models.Post) error {
	sqlStr := ` insert into post(
	post_id,title,content,user_id,community_id
	) values (?,?,?,?,?)
	`
	_, err := global.DB.Exec(sqlStr, p.PostID, p.Title, p.Content, p.UserID, p.CommunityID)
	return err
}

func GetPostDetail(p *models.Post) error {
	sqlStr := ` select post_id, user_id, community_id, title, content, create_time
	from post
	where post_id = ?
	`
	return global.DB.QueryRow(sqlStr, p.PostID).Scan(
		&p.PostID,
		&p.UserID,
		&p.CommunityID,
		&p.Title,
		&p.Content,
		&p.CreateTime,
	)
}
func CheckPIDExist(pid int) (bool, error) {
	var temp int
	sqlStr := `select post_id from post where post_id = ?`
	err := global.DB.QueryRow(sqlStr, pid).Scan(&temp)
	if err != nil {
		if err == sql.ErrNoRows { // 或者根据您的数据库驱动检查相应的“没有找到行”的错误
			return false, nil // 或者返回特定的错误，例如 errors.New("post not found")
		}
		return false, err // 返回其他类型的错误
	}
	return true, nil
}

func GetPostList(idList []string, ppl *models.ParamPostList) ([]*models.Post, error) {
	var posts []*models.Post
	if len(idList) == 0 {
		return nil, nil
	}
	// 使用逗号拼接所有元素
	idInStr := strings.Join(idList, "','")
	idOrderStr := strings.Join(idList, ",")
	// 构建 SQL 查询
	// communtiyID 为 0 的时候，在 mysql 里面限制 page 和 size
	var sqlStr string
	if ppl.ID != 0 {
		sqlStr = fmt.Sprintf(`select post_id, user_id, community_id, title, content, create_time
		from post
		where post_id in ('%s')
		order by FIND_IN_SET(post_id,'%s')
		limit %d, %d
		`, idInStr, idOrderStr, (ppl.Page-1)*ppl.Size, ppl.Size)
	} else {
		sqlStr = fmt.Sprintf(`select post_id, user_id, community_id, title, content, create_time
	from post
	where post_id in ('%s')
	order by FIND_IN_SET(post_id,'%s') 
	`, idInStr, idOrderStr)
	}
	stmt, err := global.DB.Prepare(sqlStr)
	if err != nil {
		global.Log.Errorln("db prepare error", err.Error())
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		global.Log.Errorln("db query error", err.Error())
		return nil, err
	}
	global.Log.Debugln("idstr为", idOrderStr)
	defer rows.Close()
	for rows.Next() {
		p := &models.Post{}
		err := rows.Scan(&p.PostID, &p.UserID, &p.CommunityID, &p.Title, &p.Content, &p.CreateTime)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
