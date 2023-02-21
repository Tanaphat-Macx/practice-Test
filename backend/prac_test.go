package backend

import(
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name 	string `valid:"required~Name can not  blank"`
	Url 	string `gorm:"uniqueIndex" valid:"url"` 
	Email	string `valid:"email"`
	Age     int `valid:"range(10|21)~invalid"`
}

// func TestUser(t *testing.T){
// 	g := NewGomegaWithT(t)

// 	t.Run("Name can not blank", func(t *testing.T) {
// 		U := User{
// 				Name: "",
// 				Url: "https://www.google.co.th/?hl=th",
// 			}

// 		ok, err := govalidator.ValidateStruct(U)

// 			// คาดหหวังไม่ให้ค่าที่ถูกต้องออกมาถูก ต้องออกมาผิด
// 		g.Expect(ok).NotTo(BeTrue())

// 		g.Expect(err).ToNot(BeNil())

// 		g.Expect(err.Error()).To(Equal("Name can not  blank"))
// 	})
// }

// func TestUrl(t *testing.T){
// 	g := NewGomegaWithT(t)

// 	t.Run("URL: test", func(t *testing.T) {
// 		U := User{
// 				Name: "Macx",
// 				Url: "Test",
// 			}

// 		ok, err := govalidator.ValidateStruct(U)

// 			// คาดหหวังไม่ให้ค่าที่ถูกต้องออกมาถูก ต้องออกมาผิด
// 		g.Expect(ok).NotTo(BeTrue())

// 		g.Expect(err).ToNot(BeNil())

// 		g.Expect(err.Error()).To(Equal("Url: Test does not validate as url"))
// 	})
// }

func TestEmail(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("URL: email", func(t *testing.T) {
		U := User{
				Name: "Macx",
				Url: "https://www.google.co.th/?hl=th",
				Email: "Test@hotmail.com",
				Age: 12,
			}

		ok, err := govalidator.ValidateStruct(U)

			// คาดหหวังไม่ให้ค่าที่ถูกต้องออกมาถูก ต้องออกมาผิด
		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("invalid"))
	})
}


