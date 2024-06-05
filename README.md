
##  📓 Deskripsi Singkat

Selamat datang di projek kami, projek ini bertujuan untuk mereplikasi beberapa fitur e-wallet seperti bertransaksi ke food court, melakukan transaksi, melakukan top-up yang terintegrasi dengan payment gateway, dan mencatat historiesnya. Aplikasi ini dirancang untuk memberikan kemudahan dan kenyamanan dengan tampilan yang ciamik serta intuitif.

## 📜 Swagger

[SwaggerApi](https://app.swaggerhub.com/apis-docs/EZABINTANGRAMADHAN_1/e-wallet-API/1.0.0)

## 👀 Environment Variables

Untuk menjalankan proyek ini kamu membutuhkan beberapa environment variable yang dapat kamu contoh di .env.example setelah itu kamu dapat  mengeksport dengan menggunakan perintah source .env.

Berikut adalah beberapa environment variabel yang diperlukan:

Untuk mengetahaui environment terkait cloudinary kamu dapat mengujungi ini

<https://cloudinary.com/documentation/go_quick_start>

Untuk membuat dan folder agar bisa diassign ke environment `CLOUDINARY_UPLOAD_FOLDER` kamu dapat mengikuti ini

<https://cloudinary.com/documentation/dam_folders_collections_sharing>

Terkait dengan midtrans kamu dapat mengunjungi link berikut

<https://github.com/Midtrans/midtrans-go>

## 🛠️ Installation

Untuk menjalankan program ini pertama kamu harus mengclone repository ini dengan menggunakan perintah

```bash
git clone https://github.com/Sistem-Manajemen-E-Wallet/BE-E-Wallet.git
```

masuk ke folder

```bash
cd BE-E-Wallet
```

pastikan golang dan mysql kamu sudah terinstall, jika belum silahkan kunjungi :

<https://go.dev/doc/install> <br />
<https://dev.mysql.com/doc/refman/8.3/en/windows-installation.html>

jika sudah silahkan jalankan

```bash
go mod tidy
go run .
```

Dan selamat kamu berhasil menjalankan projek ini 🎊🎊.
Untuk mengetahui endpoint-endpoint yang dapat digunakan kamu dapat melihat pada folder docs.
## 🙋‍♂️ Authors

- [@anggraanutomo](https://www.github.com/anggraanutomo)
- [@ezabintangr](https://github.com/ezabintangr)


## 👨‍💻 Tech Stack

**Server:** Golang, Echo, MySQL


