# Backend ชั่วคราว
## Start
go run main.go

## Edit .env

ปรับเปรียบตาม server mysql ที่เปิด

## สร้าง container mysql ชั่วคราวได้จาก

```bash
docker pull mysql:latest
docker run --name aws-container -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=mydatabase -p 3307:3306 -d mysql:latest
```

ใช้ mysql:latest ชื่อ aws-container
user=root
password=rootpassword
database=mydatabase
PORT=-3307 (port ของ local นะใช้สำหรับ backend)

## มี Table อยู่ใน โฟลเดอร์  database
ใช้ extension vscode run ไฟล์ .sql ได้


