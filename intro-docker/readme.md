## Intro-Docker

### Setup Docker di AWS

Silahkan konek ke AWS via SSH dari komputer masing-masing. Setelah konek, lakukan perintah dibawah ini.

```bash
    sudo apt install docker.io #Instalasi docker
    sudo chmod 777 /usr/var/docker.sock #berikan akses pada 
    docker ps -a #Validasi instalasi docker
```
Proses instalasi docker selesai

---

### Docker Cheatsheet

Beberapa fungsi yang sering digunakan pada docker


| Command | Fungsi |
| ------- | ------ |
| docker images | Menampilkan seluruh images yang dimiliki |
| docker ps -a | Menampilkan seluruh container yang dimiliki |
| docker pull <image_name> | Download image dari docker hub |
| docker run [option] <image_name> | Membuat dan menjalankan container berdasarkan image tertentu menggunakan batasan sesuai dengan option yang diberikan |
|option | . --name <nama_container> : penamaan container |
|| -p <host_port:container_port> : konfigurasi port yang akan digunakan oleh container |
|| -d : daemon mode (setelah program berhasil dibuat, akan masuk pada fase tertentu sehingga user dapat melakukan aktivitas lain) |
|| -e [env: value] : menambahkan environment variable untuk container|
| docker start <container_ID> | Mengaktifkan container  |
| docker stop <container_ID> | Menghentikan aktivitas container |
| docker rm <container_ID> | Menghapus container yang telah dispesifikasikan, hanya dapat digunakan pada container yang non aktif |
| docker rmi <container_ID> | Menghapus image yang dimiliki |
---

### Membuat Docker Image

Buatlah sebuah `Dockerfile` pada project anda. 
Berikut contoh kode minimum yang bisa anda gunakan untuk mengubah project menjadi sebuah docker image.

```bash
FROM golang:1.17

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["/app/main"]
```

Pastikan anda sedang beroperasi dalam folder project anda, jalankan perintah ini
```bash
    docker build -t <nama_image> .
```
Perintah tersebut akan membuat sebuah image berdasarkan kode pada `Dockerfile` anda.

Jika ingin upload ke Docke Hub, silahkan login docker account pada CLI anda lalu ubah format nama image menjadi `docker.io/username/nama_image`, lalu jalankan perintah 

```bash
    docker push <username/nama_image>
```