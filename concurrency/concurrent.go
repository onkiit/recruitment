package main

import (
	"encoding/json"
	"encoding/csv"
	"io/ioutil"
	"net/http"
	"bytes"
	"log"
	"os"
)

type Region struct {
	Kode	string `json:"kode_wilayah"`
	Nama	string `json:"nama"`
	Mst		string `json:"mst_kode_wilayah"`
}

type Museum struct{
	ID				string `json:"museum_id"`
	Kode_kelola		string `json:"kode_pengelolaan"`
	Nama			string `json:"nama"`
	Sdm				string `json:"sdm"`
	Alamat			string `json:"alamat_jalan"`
	Kelurahan		string `json:"desa_keluarhan"`
	Kecamatan		string `json:"kecamatan"`
	Kabupaten		string `json:"kabupaten_kota"`
	Provinsi		string `json:"propinsi"`
	Lintang			string `json:"lintang"`
	Bujur			string `json:"bujur"`
	Koleksi			string `json:"koleksi"`
	Sumber_dana		string `json:"sumber_dana"`
	Pengelola		string `json:"pengelola"`
	Tipe			string `json:"tipe"`
	Standar			string `json:"standar"`
	Tahun			string `json:"tahun_berdiri"`
	Bangunan		string `json:"bangunan"`
	Luas			string `json:"luas_tanah"`
	Status			string `json:"status_kepemilikan"`
}

type ProvinceResponse struct{
	Data	[]Region `json:"data"`
}

type DistrictResponse struct{
	Data	[]Region `json:"data"`
}

type MuseumResponse struct{
	Data	[]Museum `json:"data"`
}

var district DistrictResponse

//global request 
func request(url string, param string, target interface{}) interface{}{
	if param != "" {
		url += param
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	body, errr := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // cause error invalid character 'ï' looking for beginning of value json response
    if errr != nil {
		log.Println(errr)
	}

	erro := json.Unmarshal(body, target)
	if erro != nil{
		log.Fatal(erro)
	}

	return target //return as empty interface
}

func getProvince() []Region {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
	province := ProvinceResponse{}
	res := request(url, "", &province) // &province: json unmarshal must be pointer param
	prov_response := res.(*ProvinceResponse) //type assertion, cannot range type of interface{}, must be cast to struct
	return prov_response.Data //return array
}

func getDistrict() []Region {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET?mst_kode_wilayah="
	district := DistrictResponse{}
	var district_arr []Region
	provinces := getProvince()
	for _, province := range provinces {
		res := request(url, province.Kode, &district)
		for _, dist := range res.(*DistrictResponse).Data {
			district_arr = append(district_arr, dist)
		}
	}
	return district_arr
}

func getMuseumByCity(code string) []Museum {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota="
	museum := MuseumResponse{}
	res := request(url, code, &museum)
	museum_response := res.(*MuseumResponse)
	return museum_response.Data
}

func createCSV(district_name string, museums []Museum){
	header := []string{
		"ID Museum",
		"Kode Pengelolaan",
		"Nama",
		"SDM",
		"Alamat",
		"Desa/Kelurahan",
		"Kecamatan",
		"Kabupaten/Kota",
		"Provinsi",
		"Lintang",
		"Bujur",
		"Koleksi",
		"Sumber Dana",
		"Pengelola",
		"Tipe",
		"Standar",
		"Tahun Berdiri",
		"Bangunan",
		"Luas Tanah",
		"Status Kepemilikan" }
	f, err := os.Create(district_name + ".csv")
    if err != nil {
        log.Println(err)
    }
	defer f.Close()
	
    w := csv.NewWriter(f)
	w.Write(header)
	for _, museum := range museums {
		var data []string
		data = append(data, museum.ID, museum.Kode_kelola, museum.Nama, museum.Sdm, museum.Alamat, museum.Kelurahan, museum.Kecamatan, museum.Kabupaten, museum.Provinsi, museum.Lintang, museum.Bujur, museum.Koleksi, museum.Sumber_dana, museum.Pengelola, museum.Tipe, museum.Standar, museum.Tahun, museum.Bangunan, museum.Luas, museum.Status)
		w.Write(data)
	}
	w.Flush()
	log.Println(district_name + ".csv => File Created")
}

func createFile(district Region){
	museum := getMuseumByCity(district.Kode)
	if len(museum) > 0 {
		createCSV(district.Nama, museum)
	}
}

func main(){
	districts := getDistrict()
	for _, district := range districts {
		createFile(district)
	}
	
}