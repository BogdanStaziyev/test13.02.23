package response

type UserResponse struct {
	ID    string `json:"id" bson:"_id"`
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
}
