# Welcome to AltaPay API
<p align="center">
  <img width="170" height="170" src="https://raw.githubusercontent.com/Sistem-Manajemen-E-Wallet/FE-E-Wallet/3eb9a7c5ac3da0ee2f7f0caecc1689fbd89376ac/src/assets/logo/logo.svg">
</p>

##  📓 About The Project

Selamat datang di projek kami, projek ini bertujuan untuk mereplikasi beberapa fitur e-wallet seperti bertransaksi ke food court, melakukan transaksi, melakukan top-up yang terintegrasi dengan payment gateway, dan mencatat historiesnya. Aplikasi ini dirancang untuk memberikan kemudahan dan kenyamanan dengan tampilan yang ciamik serta intuitif.


## 🎴 Prototype
* [Figma](https://www.figma.com/design/YXwBncVF4lN75jgXaPPh36/AltaPay---Group-1?node-id=5-54&t=IzA74d3i6IAdJJSL-0)
 
## 🤝 Collaboration
* [Github Organazitation](https://github.com/Sistem-Manajemen-E-Wallet)
* [Discord](https://discord.com)

## ⚔ Backend
* [Github Repository For Backend](https://github.com/Sistem-Manajemen-E-Wallet/BE-E-Wallet)
* [GitHub Actions](https://github.com/Sistem-Manajemen-E-Wallet/BE-E-Wallet/actions)
* [Google Cloud Platform](https://cloud.google.com/)
* [Swagger Api](https://app.swaggerhub.com/apis-docs/EZABINTANGRAMADHAN_1/e-wallet-API/1.0.0)
* [Postman Collection](https://e-wallet-altera.postman.co/workspace/e-wallet-ALTERA-Workspace~7c75c907-2738-42cd-8c2d-ae851fa6663b/collection/34459708-17dc0709-35f6-4f9b-b35f-bd691a3a4ffa?action=share&creator=34459708&active-environment=34459708-cd08b619-beed-44a8-ac52-af08a29be5ec)

## 👩‍🌾 ERD
<p align="center">
  <img width="500" height="500" src="https://raw.githubusercontent.com/Sistem-Manajemen-E-Wallet/BE-E-Wallet/main/final-erd.png">
</p>

## 🎮 Features
- Berikut beberapa fitur dari AltaPay:

    - Pendaftaran dan Autentikasi
        - Registrasi pengguna: Formulir pendaftaran dengan memasukkan nomor telepon, email, alamat, pin, pin konfirmasi.
        - Login: Sistem login menggunakan nomor telepon dan pin
    - Pengelolaan Akun
        - Profil pengguna: Halaman profil untuk melihat dan mengedit informasi pribadi.
        - Saldo dompet: Tampilan saldo saat ini.
    - TopUp Dana
        - Melakukan pengisian dana: Integrasi dengan metode pembayaran seperti Virtual Account, Kartu Kredit/Debit, dan Transfer Bank
    - Transaksi
        - Pembayaran: Melakukan pengirim uang ke merchant.
        - Riwayat transaksi: Daftar semua transaksi yang pernah dilakukan berdasarkan tanggal melakukan transaksi dengan merchant maupun topup dana.
    - Keamanan
        - Keamanan transaksi: Enkripsi data.

## ✔ Unit Test
<p align="center">
  <img src="https://raw.githubusercontent.com/Sistem-Manajemen-E-Wallet/BE-E-Wallet/main/test-result.png">
</p>


## 📑 List Endpoints
| Tag | Endpoint |
| --- | --- |
| 👤 User | `POST /login` |
| 👤 User | `POST /users/customer` |
| 👤 User | `POST /users/merchant` |
| 👤 User | `GET /users` |
| 👤 User | `PUT /users` |
| 👤 User | `POST /users/changeprofilepicture` |
| 💰 Wallet | `GET /wallets` |
| 🛍️ Product | `GET /products` |
| 🛍️ Product | `POST /products` |
| 🛍️ Product | `GET /products/:id` |
| 🛍️ Product | `PUT /products/:id` |
| 🛍️ Product | `DELETE /products/:id` |
| 🛍️ Product | `GET /users/:id/products` |
| 🛍️ Product | `POST /products/:id/images` |
| 💳 Topup | `POST /topups` |
| 💳 Topup | `GET /topups` |
| 💳 Topup | `GET /topups/:id` |
| 💵 Transaction | `POST /transactions` |
| 💵 Transaction | `GET /transactions` |
| 💵 Transaction | `POST /transactions/verify` |
| 💵 Transaction | `GET /transactions/:id` |
| 💵 Transaction | `PUT /transactions/:id` |
| 📝 History | `GET /histories` |

## 🧵 Tech Stack
- **Golang**
    - -> Golang adalah bahasa pemrograman yang digunakan untuk membangun backend aplikasi.
- **Echo**
    - -> Echo adalah sebuah framework web yang bersifat minimalis, cepat, dan ekspresif, yang memudahkan pengembangan aplikasi web menggunakan bahasa pemrograman Go.
- **GORM**
    - -> GORM adalah sebuah library yang memudahkan pengembang dalam melakukan interaksi dengan database menggunakan konsep Object-Relational Mapping (ORM).
- **MySQL**
    - -> MySQL adalah Relational Database Management System (RDBMS) yang digunakan untuk membuat basis data relasional.
- **JWT (JSON Web Token)**
    - -> Standar industri untuk token akses yang digunakan untuk otentikasi dan otorisasi.
- **GCP (Google Cloud Platform)**
    - -> Layanan cloud yang digunakan untuk hosting dan layanan lainnya.
- **Cloudinary**
    - -> Platform media cloud yang digunakan untuk manajemen dan penyimpanan gambar.
- **Midtrans**
    - -> Gateway pembayaran yang digunakan untuk memproses transaksi pembayaran.
- **Docker**
    - -> Platform yang digunakan untuk mengemas aplikasi dan dependensinya dalam bentuk kontainer.
- **Postman**
    - -> Alat pengujian API yang digunakan untuk menguji endpoint API.
- **GitHub**
    - -> Platform pengembangan perangkat lunak yang digunakan untuk kontrol versi dan kolaborasi.

## 🛠️ Installation
Untuk menjalankan proyek ini kamu membutuhkan beberapa environment variable yang dapat kamu contoh di .env.example setelah itu kamu dapat  mengeksport dengan menggunakan perintah source .env.

Berikut adalah beberapa environment variabel yang diperlukan:

Untuk mengetahaui environment terkait cloudinary kamu dapat mengujungi ini

* <https://cloudinary.com/documentation/go_quick_start>

Untuk membuat dan folder agar bisa diassign ke environment `CLOUDINARY_UPLOAD_FOLDER` kamu dapat mengikuti ini

* <https://cloudinary.com/documentation/dam_folders_collections_sharing>

Terkait dengan midtrans kamu dapat mengunjungi link berikut

* <https://github.com/Midtrans/midtrans-go>

Untuk menjalankan program ini pertama kamu harus mengclone repository ini dengan menggunakan perintah

```bash
git clone https://github.com/Sistem-Manajemen-E-Wallet/BE-E-Wallet.git
```

masuk ke folder

```bash
cd BE-E-Wallet
```

pastikan golang dan mysql kamu sudah terinstall, jika belum silahkan kunjungi :

* <https://go.dev/doc/install> <br />
* <https://dev.mysql.com/doc/refman/8.3/en/windows-installation.html>

jika sudah silahkan jalankan

```bash
go mod tidy
go run .
```



## 🙋‍♂️ Authors

- [@anggraanutomo](https://www.github.com/anggraanutomo)
- [@ezabintangr](https://github.com/ezabintangr)



