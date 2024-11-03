# Backend ชั่วคราว
สามารถศึกษาแนวทางการ implement ได้จากตัวอย่างที่มีให้เลย
1. Controller
หน้าที่:
controllogic: Controllers รับผิดชอบในการจัดการกับการเรียกข้อมูลที่เข้ามาจากผู้ใช้ (เช่น ผ่าน HTTP requests) และดำเนินการตอบสนองต่อผู้ใช้ (responses) โดยทั่วไปแล้ว Controller จะเป็นตัวกลางระหว่าง Routes และ Services
ประมวลผลคำขอ: เมื่อได้รับคำขอจาก Routes Controller จะเรียกใช้ Services เพื่อดึงข้อมูลหรือดำเนินการตามที่ร้องขอ

2. Route หน้าที่: กำหนดเส้นทาง (Routes): Routes ใช้สำหรับกำหนด URL ที่ผู้ใช้สามารถเข้าถึงได้ และเชื่อมโยงกับ Controllers ที่เหมาะสม
จัดการคำขอ HTTP: Routes จะระบุว่าคำขอ HTTP ประเภทใด (GET, POST, PUT, DELETE) จะเรียกใช้ฟังก์ชันใดใน Controller

3. Service หน้าที่:ประมวลผลข้อมูล: Services จะรับผิดชอบในการจัดการกับข้อมูลที่เกี่ยวข้องกับการดำเนินการในฐานข้อมูล เช่น การดึงข้อมูล การบันทึกข้อมูล การอัปเดตข้อมูล และการลบข้อมูล
แยกความรับผิดชอบ: Services แยกการจัดการข้อมูลออกจากลอจิกของแอพพลิเคชัน ทำให้โค้ดมีความชัดเจนและดูแลรักษาได้ง่าย
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


