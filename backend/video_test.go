package backend

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}

func TestCorrectValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Correct data", func(t *testing.T) {
		v := Video{
			Name: "Macx",
			Url:  "https://www.google.co.th/?hl=th",
		}

		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).To(BeTrue())

		g.Expect(err).To(BeNil())

	})
}


func TestNameValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Name cannot be blank", func(t *testing.T) {
		v := Video{
			Name: "", //ผิด
			Url:  "https://www.google.co.th/?hl=th",
		}

		ok, err := govalidator.ValidateStruct(v)

		//คาดหวังว่าตัวแปรที่ OK จะไม่แสดงค่า True
		g.Expect(ok).NotTo(BeTrue())

		//คาดหวังว่าตัวแปรที่ OK จะไม่แสดงค่า Nil
		g.Expect(err).ToNot(BeNil())

		//คาดหวังว่าตัวแปรที่ OK จะแสดงค่าที่กำหนดใน()
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})
}

func TestUrlValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Url: Test does not validate as url", func(t *testing.T) {
		v := Video{
			Name: "Tanaphat",
			Url:  "Test", //ผิด
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("Url: Test does not validate as url"))

	})
}
