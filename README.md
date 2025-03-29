
# Web Application: AWS (Amazing Work Space)

## วัถตุประสงค์
อำนวยความสะดวกในการจัดการและบริหารงานภายในทีม โดยเน้นการทำงานร่วมกันอย่างเป็นระบบ มีประสิทธิภาพ และโปร่งใส เพื่อเพิ่มประสิทธิภาพการทำงาน ลดความยุ่งยากในการติดตามสถานะงาน และสร้างพื้นที่ที่เหมาะสมสำหรับการสื่อสารและการทำงานร่วมกันในทีม

## จัดทำโดย
| ชื่อ | รหัสนิสิต | หมู่ | ชั้นปี |
| :- | -: | :-: | -: |
นาย ธนธัส สุวรรณ์ | 6510450429 | 202 | 3
น.ส. วัชราพร ภูวะนสุขสุนทร | 6510450933 | 200 | 3
นาย ศุภกฤต ปะมาคะมา | 6510450968 | 200 | 3
นาย สิทธิภัทท์ เทพสุธา | 6510451000 | 200 | 3

## Features (Backend Only)
@karnhao Feat: Communication, Listview, Calendar

@Peet555 Feat: Communication

@pickky007 Feat: Communication

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