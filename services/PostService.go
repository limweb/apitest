package services 
 
import ( 
	"apitest/db" 
	"apitest/models" 
	"fmt" 
) 
 
func GetAllPost(b *[]models.Post) (err error) { 
	if err = db.GetDB().Find(b).Error; err != nil {  // select * from posts
		return err 
	} 
	return nil 
} 
 
func GetAllIdDescPost(b *[]models.Post) (err error) { 
	if err = db.GetDB().Order("id desc").Find(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func AddNewPost(b *models.Post) (err error) { 
	if err = db.GetDB().Create(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func GetOnePost(b *models.Post, id string) (err error) { 
	if err := db.GetDB().Where("id = ?", id).First(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func PutOnePost(b *models.Post, id string) (err error) { 
	fmt.Println(b) 
	db.GetDB().Save(b) 
	return nil 
} 
 
func DeletePost(b *models.Post, id string) (err error) { 
	db.GetDB().Where("id = ?", id).Delete(b) 
	return nil 
} 
 