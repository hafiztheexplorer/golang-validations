package golangvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

////////////////////////////////////////////////////////////////////////////////////////////////

func TestValidation(t *testing.T) {
	validkah := validator.New()
	if validkah == nil {
		t.Error("validkah adalah nil")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////

func TestValidasiMultiTag(t *testing.T) {
	validkah4 := validator.New()

	//	misal user harus angka, selain tidak boleh kosong
	var user string = "1234"

	err := validkah4.Var(user, "number,required")

	if err != nil {
		fmt.Println("berikut hasil validasinya = ", err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////

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

type DataPengguna struct {
	Id           string                 `validate:"required"`
	NamaPengguna string                 `validate:"required"`
	Alamat       []DataAlamat           `validate:"required,dive"`
	Hobi         []string               `validate:"required,dive,required,min=3"`
	Sekolah      map[string]DataSekolah `validate:"required,dive,keys,required,min=2,endkeys,dive"`
}

type DataAlamat struct {
	Kota   string `validate:"required"`
	Negara string `validate:"required"`
}

type DataSekolah struct {
	NamaSekolah string `validate:"required"`
}

func TestValidasiMap(t *testing.T) {

	validkah6data := validator.New()
	ContohInputan := DataPengguna{
		Id:           "",
		NamaPengguna: "",
		Alamat: []DataAlamat{
			{
				Kota:   "",
				Negara: "",
			},
		},
		Hobi: []string{
			"",
			"",
			"",
		},
		Sekolah: map[string]DataSekolah{
			"SD": {
				NamaSekolah: "",
			},
			"SMP": {
				NamaSekolah: "",
			},
			"SMA": {
				NamaSekolah: "",
			},
			"": {
				NamaSekolah: "",
			},
		},
	}

	err := validkah6data.Struct(ContohInputan)
	if err != nil {
		fmt.Println(err.Error())

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type DataPengguna2 struct {
	Id           string                  `validate:"required"`
	NamaPengguna string                  `validate:"required"`
	Alamat       []DataAlamat2           `validate:"required,dive"`
	Hobi         []string                `validate:"required,dive,required,min=3"`
	Sekolah      map[string]DataSekolah2 `validate:"required,dive,keys,required,min=2,endkeys,dive"`
	Wallet       map[string]int          `validate:"dive,keys,required,endkeys,required,gt=1000"`
}

type DataAlamat2 struct {
	Kota   string `validate:"required"`
	Negara string `validate:"required"`
}

type DataSekolah2 struct {
	NamaSekolah string `validate:"required"`
}

func TestValidasiBasicMap(t *testing.T) {

	validkah6data := validator.New()
	ContohInputan := DataPengguna2{
		Id:           "",
		NamaPengguna: "",
		Alamat: []DataAlamat2{
			{
				Kota:   "",
				Negara: "",
			},
		},
		Hobi: []string{
			"",
			"",
			"",
		},
		Sekolah: map[string]DataSekolah2{
			"SD": {
				NamaSekolah: "",
			},
			"SMP": {
				NamaSekolah: "",
			},
			"SMA": {
				NamaSekolah: "",
			},
			"": {
				NamaSekolah: "",
			},
		},
		Wallet: map[string]int{
			"BCA":  10000000,
			"BCPT": 0,
			"":     1000,
		},
	}

	err := validkah6data.Struct(ContohInputan)
	if err != nil {
		fmt.Println(err.Error())

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type LoginRequest2 struct {
	Username string `validate:"custom1,min=5,email"`
	Password string `validate:"custom1,min=5"`
}

func TestValidasiAliasTag(t *testing.T) {

	validkah6data := validator.New()
	validkah6data.RegisterAlias("custom1", "required,max=255")

	var loginrequest1 LoginRequest2

	// Required dan ini harus format email, karena sudah diset di atas
	loginrequest1.Username = "sdasdasdwdsad"
	loginrequest1.Password = ""

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

// misal ini membuat tag untuk memastikan username yang dibuat beneran
// dengan menggunakan parameter validator.FieldlLevel, dengan return value nya boolean
func MustValidUsername(field validator.FieldLevel) bool {

	//field itu balikannya reflection value, jadi nanti Field ini isinya bakalan semua baik itu string atau int makannya pakai interface, dan hasilnya itu dikoknversi menjadi string
	value, ok := field.Field().Interface().(string)
	if ok {

		//jika usernamenya tidak uppercase, direturn false
		if value != strings.ToUpper(value) {
			return false
		}

		//jika panjang tidak lebih dari 5 huruf, direturn false
		if len(value) < 5 {
			return false
		}
	}

	//ini kalau kondisi memenuhi, misal buat username, itu panjangnya lebih dari 5 huruf
	return true
}

//semua itu akan diproses di function testing di bawah

type LoginRequest3 struct {
	Username string `validate:"custom1,min=5,username"`
	Password string `validate:"custom1,min=5"`
}

func TestValidasiCustom(t *testing.T) {

	validkah6data := validator.New()
	validkah6data.RegisterAlias("custom1", "required,max=255")
	validkah6data.RegisterValidation("username", MustValidUsername)

	var loginrequest1 LoginRequest3

	// Required dan ini harus format username, karena sudah diset di atas
	//dan ini merupakan custom tag
	loginrequest1.Username = "DARTHSIDIOUS"
	loginrequest1.Password = ""

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err.Error())
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

type LoginRequestNew struct {
	Username string `validate:"custom1,email"`
	Password string `validate:"custom1,min=5"`
	Pin      string `validate:"required,Pincheck=6"`
}

func TestValidasiCustomParameter(t *testing.T) {

	validkah6data := validator.New()
	validkah6data.RegisterAlias("custom1", "required,max=255")
	validkah6data.RegisterValidation("Pincheck", MustValidPin)

	var loginrequest1 LoginRequestNew

	// Required dan ini harus format username, karena sudah diset di atas
	//dan ini merupakan custom tag
	loginrequest1.Username = "DARTHSIDIOUS"
	loginrequest1.Password = ""
	loginrequest1.Pin = "66666"

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type LoginRequestNew2 struct {
	Username string `validate:"custom1,email"`
	Password string `validate:"custom1,min=5"`
	Pin      string `validate:"required,Pincheck=6"`
}

func TestValidasiCustomParameter2(t *testing.T) {

	validkah6data := validator.New()
	validkah6data.RegisterAlias("custom1", "required,max=255")
	validkah6data.RegisterValidation("Pincheck", MustValidPin)

	var loginrequest1 LoginRequestNew2

	// Required dan ini harus format username, karena sudah diset di atas
	//dan ini merupakan custom tag
	loginrequest1.Username = "987867727" //misal ini bisa nomor telepon juga atau email
	loginrequest1.Password = ""
	loginrequest1.Pin = "66666"

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type LoginRequestNew3 struct {
	//perhatikan tag nya dalam kondisi OR atau |, dan tidak lagi terpisah
	//jadi salah satu kondisi bisa memenuhi, contoh pada username
	//boleh email, boleh numerical
	Username string `validate:"required,email|numeric"`
	Password string `validate:"required,min=5"`
	Pin      string `validate:"required,numeric|max=6"`
}

func TestValidasiTagOrRule(t *testing.T) {

	validkah6data := validator.New()

	var loginrequest1 LoginRequestNew3

	// Required dan ini harus format username, karena sudah diset di atas
	loginrequest1.Username = "08132344233" //misal ini bisa nomor telepon juga atau email
	loginrequest1.Password = "password"
	loginrequest1.Pin = "77886633"

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	rvalue, _, _, bool2 := field.GetStructFieldOK2()
	if !bool2 {
		panic("field not ok")
	}

	firstvalue := strings.ToUpper(field.Field().String())
	secondvalue := strings.ToUpper(rvalue.String())

	return firstvalue == secondvalue
}

type LoginRequestNew4 struct {
	Username     string `validate:"required,field_equal_ignore_case=Email|field_equal_ignore_case=NomorTelepon"`
	Email        string `validate:"required,email"`
	NomorTelepon string `validate:"required,numeric"`
	Nama         string `validate:"required"`
}

func TestValidasiCrossFieldCustom(t *testing.T) {

	validkah6data := validator.New()
	validkah6data.RegisterValidation("field_equal_ignore_case", MustEqualsIgnoreCase)

	loginrequest1 := LoginRequestNew4{
		Username:     "email@email.com",
		Email:        "Email@email.com",
		NomorTelepon: "081888222992",
		Nama:         "contoh nama user",
	}

	err := validkah6data.Struct(loginrequest1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

type RegisterReq struct {
	Username     string
	Email        string
	NomorTelepon string
	Nama         string
}

//kalau diamati di atas tanpa validasi tag seperti biasanya
//alih - alih validasinya tertanam pada function "RegisterSuksesValidasi"

func TestValidasiStructLevel(t *testing.T) {

	validkah6data := validator.New()
	//validasi diregistrasikan dengan RegisterStructValidation
	validkah6data.RegisterStructValidation(RegisterSuksesValidasi, RegisterReq{})

	RegReq := RegisterReq{
		Username:     "email@mail.com",
		Email:        "email@mail.com",
		NomorTelepon: "0887772211",
		Nama:         "contoh nama user",
	}

	err := validkah6data.Struct(RegReq)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func RegisterSuksesValidasi(level validator.StructLevel) {
	Register := level.Current().Interface().(RegisterReq)
	if Register.Username == Register.Email || Register.Username == Register.NomorTelepon {
		//success
	} else {
		//tag nya di sini "Username", "Username", "Username"
		level.ReportError(Register.Username, "Username", "Username", "Username", "")
	}

}
