package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

// //////////////////////////////////////////////////////////////////////////////////////////////
func TestValidation(t *testing.T) {
	validkah := validator.New()
	if validkah == nil {
		t.Error("validkah adalah nil")
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////
func TestValidationField(t *testing.T) {
	validkah2 := validator.New()

	// variabel yang akan dicek, divalidasikan
	// coba di sini kasih inputan kosong, jadi tanda petik 2 aja x2, pasti keluar errnya apa
	// kalau diisi pasti beda lagi, passed dan tanpa warning
	//tetapi kalau kosong pasti passed juga, tapi terjadi error
	var user string = "username example"

	// proses validasi di sini
	err := validkah2.Var(user, "required")

	if err != nil {
		fmt.Println("berikut hasil validasinya = ", err.Error())
	}

}

// //////////////////////////////////////////////////////////////////////////////////////////////
func TestValidasiDuaVariabel(t *testing.T) {
	validkah3 := validator.New()

	//misal contoh inputan ganti password yang biasanya disuruh isi 2x
	inputanpasswordpertama := "sesuatu"
	inputanpasswordkonfirmasi := "sesuatu"

	err := validkah3.VarWithValue(inputanpasswordpertama, inputanpasswordkonfirmasi, "eqfield")
	if err != nil {
		fmt.Println("berikut hasil validasinya = ", err.Error())
	}

}

// //////////////////////////////////////////////////////////////////////////////////////////////
func TestValidasiMultiTag(t *testing.T) {
	validkah4 := validator.New()

	//	misal user harus angka, selain tidak boleh kosong
	var user string = "1234"

	err := validkah4.Var(user, "number,required")

	if err != nil {
		fmt.Println("berikut hasil validasinya = ", err.Error())
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////
func TestValidasiDenganTagParameter(t *testing.T) {
	validkah5 := validator.New()

	//	misal user harus angka, selain tidak boleh kosong, harus nomor, panjangnya harus antara 5-10 karakter
	user := "99999999999"
	fmt.Println("panjang varnya = ", len(user))

	err := validkah5.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////
type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

func TestValidasiTagParameterStruct(t *testing.T) {

	validkah6data := validator.New()

	var loginrequest1 LoginRequest

	// Required dan ini harus format email, karena sudah diset di atas
	loginrequest1.Username = "contoh@mail.com"
	loginrequest1.Password = "contohpasword"

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		fmt.Println(err.Error())

	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////
func TestValidasiError(t *testing.T) {

	validkah6data := validator.New()

	var loginrequest1 LoginRequest

	// Required dan ini harus format email, karena sudah diset di atas
	loginrequest1.Username = "contoh@mail.com"
	loginrequest1.Password = "contohpasword"

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		ValidationError := err.(validator.ValidationErrors)
		for i, FieldError := range ValidationError {
			// untuk bentuk for seperti ini coba lihat https://www.w3schools.com/go/go_loops.php
			fmt.Println(i, "error = ", FieldError.Field(), "on tag = ", FieldError.Tag(), "with error = ", FieldError.Error())
		}

	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////
type RegUser struct {
	Username        string `validate:"required,email"`
	Password        string `validate:"required,min=5"`
	ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
}

func TestValidasiCrossField(t *testing.T) {

	validkah6data := validator.New()

	var loginrequest2 RegUser

	// Required dan ini harus format email, karena sudah diset di atas
	loginrequest2.Username = "contoh@mail.com"
	loginrequest2.Password = "contohpassword"
	loginrequest2.ConfirmPassword = "contohpasword"

	err := validkah6data.Struct(loginrequest2)
	if err != nil {
		fmt.Println(err.Error())

	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////
type DataDiri struct {
	Nama     string   `validate:"required,max=50"`
	Umur     int      `validate:"required,max=3"`
	Alamat   string   `validate:"required,max=200"`
	Email    string   `validate:"required,email"`
	DataUser DataUser //`validate:"required"` ini nested struct,
}

//sepertinya tidak perlu validate di nested struct nya karena di bawah ini
//masing masing sudah divalidate sendiri per item nya

type DataUser struct {
	Username   string `validate:"required"`
	Password   string `validate:"required,min=5"`
	RePassword string `validate:"required,min=5,eqfield=Password"`
}

func TestValidasiNestedStruct(t *testing.T) {

	validkah6data := validator.New()
	Request := DataDiri{
		Nama:   "",
		Umur:   20,
		Alamat: "",
		Email:  "",
		DataUser: DataUser{
			Username:   "",
			Password:   "",
			RePassword: "",
		},
	}

	err := validkah6data.Struct(Request)
	if err != nil {
		fmt.Println(err.Error())

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type DataDiri2 struct {
	Nama      string     `validate:"required,max=50"`
	Umur      int        `validate:"required,max=3"`
	Alamat    string     `validate:"required,max=200"`
	Email     string     `validate:"required,email"`
	DataUsers []DataUser `validate:"required,dive"`
}

//sepertinya tidak perlu validate di nested struct nya karena di bawah ini
//masing masing sudah divalidate sendiri per item nya

type DataUser2 struct {
	Username   string `validate:"required"`
	Password   string `validate:"required,min=5"`
	RePassword string `validate:"required,min=5,eqfield=Password"`
}

func TestValidasiCollectionStruct(t *testing.T) {

	validkah6data := validator.New()
	Request := DataDiri2{
		Nama:   "",
		Umur:   0,
		Alamat: "",
		Email:  "",
		DataUsers: []DataUser{
			{
				Username:   "",
				Password:   "",
				RePassword: "",
			},
			{
				Username:   "",
				Password:   "",
				RePassword: "",
			},
		},
	}

	err := validkah6data.Struct(Request)
	if err != nil {
		fmt.Println(err.Error())

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type DataDiri3 struct {
	Nama      string     `validate:"required,max=50"`
	Umur      int        `validate:"required,max=3"`
	Alamat    string     `validate:"required,max=200"`
	Email     string     `validate:"required,email"`
	DataUsers []DataUser `validate:"required,dive"`
	Hobbies   []string   `validate:"required,dive,required,min=3"` //agar data hobbies seacara keseluruhan dan di dalamnya itu sama sama required
}

//sepertinya tidak perlu validate di nested struct nya karena di bawah ini
//masing masing sudah divalidate sendiri per item nya

type DataUser3 struct {
	Username   string `validate:"required"`
	Password   string `validate:"required,min=5"`
	RePassword string `validate:"required,min=5,eqfield=Password"`
}

func TestValidasiBasicCollection(t *testing.T) {

	validkah6data := validator.New()
	Request := DataDiri3{
		Nama:   "",
		Umur:   0,
		Alamat: "",
		Email:  "",
		DataUsers: []DataUser{
			{
				Username:   "",
				Password:   "",
				RePassword: "",
			},
			{
				Username:   "",
				Password:   "",
				RePassword: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Codings",
			"",
			"X",
		},
	}

	err := validkah6data.Struct(Request)
	if err != nil {
		fmt.Println(err.Error())

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type SchoolData struct {
	Nama string `validate:"required"`
}

type User struct {
	Id       string                `validate:"required"`
	NamaUser string                `validate:"required"`
	Address  []Address             `validate:"required,dive"`
	Hobbies  []string              `validate:"required,required,min=1"`
	Schools  map[string]SchoolData `validate:"required,dive,keys,required,min=2,endkeys,dive"`
}

func TestValidasiMap(t *testing.T) {

	validkah6data := validator.New()
	Request := DataDiri3{
		Nama:   "",
		Umur:   0,
		Alamat: "",
		Email:  "",
		DataUsers: []DataUser{
			{
				Username:   "",
				Password:   "",
				RePassword: "",
			},
			{
				Username:   "",
				Password:   "",
				RePassword: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Codings",
			"",
			"X",
		},
	}

	err := validkah6data.Struct(Request)
	if err != nil {
		fmt.Println(err.Error())

	}
}
