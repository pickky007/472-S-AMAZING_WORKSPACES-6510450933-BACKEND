## Web Application: AWS (Amazing Work Space)

### Features
1. **Login**: ยืนยัน user เพื่อเข้าใช้งาน
2. **Register**: สร้าง new user
3. **ระบบ Workspace**  
   Workspace เป็นระบบที่ผู้ใช้งานสามารถทำงานร่วมกันได้ โดยผู้ใช้ทุกคนสามารถสร้าง Workspace และเข้าร่วม Workspace ของผู้อื่นได้
4. **Kanban Board**  
   Kanban board จะอยู่ภายใน Workspace โดยจะมี section ย่อยอยู่ภายใน
   - **Section**  
     สามารถสร้าง section ที่ต้องการได้เพื่อกำหนดว่า activity นี้อยู่ใน section ไหน เช่น activity Ux/Ui อยู่ใน section Doing หมายถึงกำลังทำอยู่
   - **Activity**  
     แต่ละ activity จะมีชื่อ คำอธิบาย วันที่เริ่ม และวันที่คาดว่าจะเสร็จ  
     ตัวอย่าง:
     - **Name**: Ux/Ui
     - **Description**: ทำการออกแบบ interface
     - **Start Date**: วันที่และเวลาที่สร้าง activity
     - **End Date**: วันที่และเวลาที่ทำเสร็จ
## Backend 
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

## จัดทำโดย  
นาย ธนโชติ งามคณะ  6510450411  
นาย ธนธัส สุวรรณ์  6510450429  
น.ส. วัชราพร ภูวะนสุขสุนทร  6510450933  
นาย ศุภกฤต ปะมาคะมา  6510450968  
นาย สิทธิภัทท์ เทพสุธา  6510451000 
