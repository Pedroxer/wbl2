package main

import "fmt"

// handler
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
	isIll             bool
}

func NewPatient(s string) *Patient {
	return &Patient{Name: s}
}

// concrete product 1
type reception struct {
	next Department
}

func NewReception() *reception {
	return &reception{}
}

func (r *reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *reception) SetNext(next Department) {
	r.next = next
}

// concrete product 2
type doctor struct {
	next Department
}

func NewDoctor() *doctor {
	return &doctor{}
}

func (d *doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		if p.isIll {
			d.next.Execute(p)
		} else {
			fmt.Println("Patient went home")
		}
		return
	}

	fmt.Println("Doctor checking patient")

	p.DoctorCheckUpDone = true
	if p.isIll {
		d.next.Execute(p)
	} else {
		fmt.Println("Patient went home")
	}
}

func (d *doctor) SetNext(next Department) {
	d.next = next
}

// concrete handler 1
type medical struct {
	next Department
}

func NewMedical() *medical {
	return &medical{}
}

func (m *medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *medical) SetNext(next Department) {
	m.next = next
}

// concrete handler 2
type cashier struct {
	next Department
}

func NewCashier() *cashier {
	return &cashier{}
}

func (c *cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *cashier) SetNext(next Department) {
	c.next = next
}
func main() {
	hospitalPatient := NewPatient("abc")
	hospitalPatient.isIll = false
	cashier := NewCashier()
	// Определяем куда отправить пациент после комнаты выдачи медикаментов
	medical := NewMedical()
	medical.SetNext(cashier)
	// Определяем куда отправить пациент после кабинета врача
	doctor := NewDoctor()
	doctor.SetNext(medical)
	// Определяем куда отправить пациент после регистратуры
	reception := NewReception()
	reception.SetNext(doctor)

	// Пациент приходит в больницу
	reception.Execute(hospitalPatient)
}
