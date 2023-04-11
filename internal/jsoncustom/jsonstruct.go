package jsoncustom

type Result struct {
	Gender  string  `jsoncustom:"gender"`
	Name    Name    `jsoncustom:"name"`
	Email   string  `jsoncustom:"email"`
	Login   Login   `jsoncustom:"login"`
	DOB     DOB     `jsoncustom:"dob"`
	Phone   string  `jsoncustom:"phone"`
	Cell    string  `jsoncustom:"cell"`
	ID      ID      `jsoncustom:"id"`
	Picture Picture `jsoncustom:"picture"`
}

type Name struct {
	Title string `jsoncustom:"title"`
	First string `jsoncustom:"first"`
	Last  string `jsoncustom:"last"`
}

type Login struct {
	UUID     string `jsoncustom:"uuid"`
	Username string `jsoncustom:"username"`
	Password string `jsoncustom:"password"`
	Salt     string `jsoncustom:"salt"`
	MD5      string `jsoncustom:"md5"`
	SHA1     string `jsoncustom:"sha1"`
	SHA256   string `jsoncustom:"sha256"`
}

type DOB struct {
	Date string `jsoncustom:"date"`
	Age  int    `jsoncustom:"age"`
}

type ID struct {
	Name  string `jsoncustom:"name"`
	Value string `jsoncustom:"value"`
}

type Picture struct {
	Large     string `jsoncustom:"large"`
	Medium    string `jsoncustom:"medium"`
	Thumbnail string `jsoncustom:"thumbnail"`
}

type Info struct {
	Seed    string `jsoncustom:"seed"`
	Result  int    `jsoncustom:"result"`
	Page    int    `jsoncustom:"page"`
	Version string `jsoncustom:"version"`
}

type Response struct {
	Results []Result `jsoncustom:"results"`
	Info    Info     `jsoncustom:"info"`
}
