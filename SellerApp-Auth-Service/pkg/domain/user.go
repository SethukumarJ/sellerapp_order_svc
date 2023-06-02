package domain

type Users struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name" gorm:"unique" validate:"required,min=2,max=50"`
	FirstName string `json:"first_name" gorm:"not null;default:null" sql:"type:varchar(255);check:first_name <> 'rahul'"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"not null;unique" validate:"email,required"`
	Password  string `json:"password"`
}

func (u *Users) NotRahul() bool {
    return u.FirstName == "rahul"
}


func (s *Users) omitemptyName() bool {
	return s.FirstName == "rahul"
}
