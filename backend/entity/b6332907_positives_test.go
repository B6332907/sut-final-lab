package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

// `valid:"matches(^[L]//d{7}$)"`
type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name not blank"`
	Email      string
	CustomerID string `valid:"matches(^[L||M||H]\\d{7}$)~error ID customer"`
}

// func TestName(t *testing.T) {
// 	g := gomega.NewGomegaWithT(t)

// 	Customer := Customer{
// 		Name:       "",
// 		Email:      "sameaw13@gmail.com",
// 		CustomerID: "L1234567",
// 	}

// 	ok, err := govalidator.ValidateStruct(Customer)

// 	g.Expect(ok).NotTo(gomega.BeTrue())

// 	g.Expect(err).NotTo(gomega.BeNil())

// 	g.Expect(err.Error()).To(gomega.Equal("Name not blank"))
// }

func TestPositive(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	Customer := Customer{
		Name:       "Peem",
		Email:      "sameaw13@gmail.com",
		CustomerID: "M1234567",
	}

	ok, err := govalidator.ValidateStruct(Customer)

	g.Expect(ok).To(gomega.BeTrue())
	
	g.Expect(err).To(gomega.BeNil())

}

// func TestIDcustomer(t *testing.T) {
// 	g := gomega.NewGomegaWithT(t)

// 	Customer := Customer{
// 		Name:       "Peem",
// 		Email:      "sameaw13@gmail.com",
// 		CustomerID: "gjhj1234567",
// 	}

// 	ok, err := govalidator.ValidateStruct(Customer)

// 	g.Expect(ok).NotTo(gomega.BeTrue())

// 	g.Expect(err).NotTo(gomega.BeNil())

// 	g.Expect(err.Error()).To(gomega.Equal("error ID customer"))

// }
