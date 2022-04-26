## Intro to AWS ##

Silahkan login ke console aws

> https://signin.aws.amazon.com

Silahkan search menu `EC2`

Setelah masuk ke menu `EC2` silahkan create instance berdasarkan rekomendasi yang ada.
Untuk pertama kali pembuatan silahkan membuat key-pair pada kolom yang telah disediakan.

Masuk ke menu `Connect` -> `SSH` -> Copy command ssh yang telah disediakan

Pada komputer anda, buka `gitbash`/`terminal`(linux) arahkan ke folder yang menyimpan file key-pair anda. Paste command yang telah anda copy lalu jalankan.

Setelah berhasil konek ke aws, lakukan beberapa perintah ini

```sh
    sudo apt update #Update list library
    sudo apt upgrade #Upgrade komponen yang diperlukan 
    sudo apt install docker.io #install docker
    sudo chmod 777 /var/run/docker.sock #untuk memberikan akses menggunakan docker
    docker ps #coba eksekusi docker

    #untuk keperluan docker
    sudo apt install pyhton3-pip
    sudo pip3 install docker-compose
```