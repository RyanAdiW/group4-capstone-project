
<!-- ABOUT THE PROJECT -->
## 💻 &nbsp;About The Project
E-Assets merupakan suatu project Capstone untuk membangun sebuah Rest API Employee Assets Management App menggunakan Golang. App ini menjadi sebuah wadah bagi employee perusahaan untuk melakukan pengajuan peminjaman asset perusahaan. 

### Features
- Employee
    - Login/Logout
    - Melihat semua aset perusahaan
    - Melihat aset perusahaan berdasarkan kategori dan/atau ketersediaan barang
    - Mengajukan perminataan peminjaman aset
    - Mengembalikan aset
    - Melihat histori peminjaman aset
    - Melihat status peminjaman aset
- Admin
    - Login/Logout
    - Menambah aset baru
    - Mengupdate aset perusahaan
    - Melihat dan/atau mengupdate status peminjaman aset
    - Melihat permintaan peminjaman aset berdasarkan kategori aset, status peminjaman, dan tanggal permintaan
    - Menyetujui dan menolak permintaan peminjaman aset
    - Meng-assign aset ke employee
- Manager
    - Login/Logout
    - Melihat list permintaan peminjaman aset berdasarkan tanggal permintaan dan/atau tanggal pengembalian
    - Melihat list permintaan peminjaman aset berdasarkan status peminjaman
    - Melakukan persetujuan/penolakan perminataan peminjaman

### &nbsp;Images
<details>
<summary>&nbsp;🖼 ERD</summary>
<img src="image/ERD.png">
</details>

### 🕮 &nbsp;OpenAPI Documentation

See documentation [here](https://app.swaggerhub.com/apis-docs/RyanAdiW/Employee-Assets-Management/1.0.0).

### 🛠 &nbsp;Built With

![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![AWS](https://img.shields.io/badge/-AWS-05122A?style=flat&logo=amazon)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;
![Ubuntu](https://img.shields.io/badge/-Ubuntu-05122A?style=flat&logo=ubuntu)&nbsp;

## How to Use
### * Running on Local Server
- Install Golang, Postman, MySQL Management System (ex. MySQL Workbench)
- Clone repository with HTTPS:
```
git clone https://github.com/RyanAdiW/group4-capstone-project.git
```
* Create File `.env`:
```
export DB_USERNAME=[username db]
export DB_PASSWORD=[password db]
export DB_ADDRESS=[ip addres db]
export DB_NAME=[name db]
export S3_REGION=[S3 region]
export S3_KEY_ID=[S3 key id]
export S3_ACCESS_KEY=[S3 access key]
export S3_BUCKET_NAME=[S3 bucket name]
```
* Run `main.go` on local terminal
```
$ source .env && go run app/main.go
```
* Run the endpoint according to the OpenAPI Documentation (Swagger) via Postman 

<!-- CONTACT -->
## 📮 &nbsp;Contact

[![GitHub Ryan](https://img.shields.io/badge/-Ryan-white?style=flat&logo=github&logoColor=black)](https://github.com/ryanadiw)
[![GitHub Hilmi](https://img.shields.io/badge/-Hilmi-white?style=flat&logo=github&logoColor=black)](https://github.com/hilmihi)

<p align="center">:copyright: 2022</p>
</h3>