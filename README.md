# Menu API

## การติดตั้งและการใช้งาน

### ข้อกำหนดระบบ

โปรเจกต์นี้ต้องการ:

- Go 1.19 ขึ้นไป
- MySQL หรือ MariaDB
- Git (ถ้าต้องการ Clone โปรเจกต์จาก Git repository)

### ขั้นตอนการติดตั้ง

1. Clone โปรเจกต์จาก Git repository

   `bash
git clone https://github.com/teeratchapat/golang-beer-example`

2. ติดตั้ง dependencies ใช้คำสั่งต่อไปนี้เพื่อติดตั้ง dependencies

`go mod tidy`

3. ตั้งค่า docker

`docker-compose up --build`

5. รันโปรเจกต์

`go run main.go`
