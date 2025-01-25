package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	// ข้อมูลการเชื่อมต่อ
	server := "ingestprogram.thaipbs.or.th:21" // พอร์ต FTP ปกติคือ 21
	user := "124"
	pass := "1234"

	// เชื่อมต่อกับ FTP เซิร์ฟเวอร์
	c, err := ftp.Dial(server, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์: %v", err)
	}

	// ล็อกอินด้วยชื่อผู้ใช้และรหัสผ่าน
	err = c.Login(user, pass)
	if err != nil {
		log.Fatalf("ล็อกอินล้มเหลว: %v", err)
	}

	fmt.Println("เชื่อมต่อและล็อกอินสำเร็จ")

	// ดึงรายการไดเรกทอรีหลัก
	entries, err := c.List("/")
	if err != nil {
		log.Fatalf("ไม่สามารถดึงรายการโฟลเดอร์ได้: %v", err)
	}

	fmt.Println("รายการโฟลเดอร์:")
	for _, entry := range entries {
		fmt.Printf(" - %s (Type: %v)\n", entry.Name, entry.Type)
	}

	// ล็อกเอาท์ออกจากเซิร์ฟเวอร์
	err = c.Logout()
	if err != nil {
		log.Fatalf("การล็อกเอาท์ล้มเหลว: %v", err)
	}

	fmt.Println("ล็อกเอาท์สำเร็จ")
}
