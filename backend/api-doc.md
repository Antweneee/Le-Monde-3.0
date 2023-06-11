# API DOCUMENTATION

## DATABASE ROUTES

**Article**
Route : /articles
Type : POST
Body : File uploaded
Description : Crée un article

Route : /articles/:cid
Type : GET
Description : Récupère un article

Route : /articles/:cid
Type : DELETE
Description : Supprime un article

**USER**:

Route : /database/User
Type : POST
Body: {
    Email string
    Username string
    Password string
}
Description: Ajoute un user

Route : /database/User:email
Type : GET
Query: {
    email string
}
Description: Récupère un user grace à son email

Route : /database/User
Type : DELETE
Body: {
    Email string
    Username string
    Password string
}
Description: Supprime un user

**Articles**:

Route : /database/Article
Type : POST
Body: {
    UserId uint
    Title string
    Cid string
    CreatedAt string
    TopicId int
}
Description: Créer un Article

Route : /database/Article:cid
Query: {
    cid string
}
Type : GET
Description: Récupère un Article grace à son cid

Route : /database/Article
Type : DELETE
Body: {
    email string
    username string
    password string
}
Description: Supprime un Article

**LIKE**:
Route : /database/Like
Type : POST
Body: {
    UserId uint
	ArticleId uint
}
Description: Ajoute un like

**Bookmark**:
Route : /database/Bookmark
Type : POST
Body: {
    Name string
	Description string
	UserId uint
}
Description: Ajoute un bookmark

**Topic**:
Route : /database/Topic
Type : POST
Body: {
    Name string
}
Description: Ajoute un Topic

