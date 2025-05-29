
### ✅ Alasan Penggunaan Layering Pattern

1. **Separation of Concerns**  
   Karena terdapat 3 layer pada layering patern yaitu:
   - `presentation` yang diatur oleh `handler` cuma ngurusin request & response
   - `business` yang diatur oleh `services` fokus pada aturan bisnis
   - `persistance` yang diatur oleh `repository` cuma urus database  
   Ini bikin kode Mudah untuk dipahami dan mudah dibaca.

2. **Mudah Dites (Testable)**  
   Karena dependency bisa dimock per-layer, unit test bisa dibuat tanpa ngetes keseluruhan sistem (misal mock repo saat ngetes service).

3. **Maintainable & Scalable**  
   Ketika aplikasi bertambah kompleks, kita tinggal split layer jadi lebih modular tanpa refactor besar-besaran.

4. **Reusable**  
   Fungsi-fungsi di `service` bisa dipanggil di banyak handler tanpa duplikasi logic.

## 🚀 Fitur Utama

- ✅ Register & Login (JWT)
- ✅ CRUD resource (contoh: Product, User, dll)
- ✅ Validasi request menggunakan `go-playground/validator`
- ✅ Middleware Auth
- ✅ end-to-end test menggunakan `testify`

## 🧪 Testing

Pengujian End-to-End dilakukan menggunakan library `testify` untuk mempermudah penulisan assertion. Pada tahap ini, tidak digunakan mock object karena pengujian diarahkan langsung ke database nyata, sehingga dapat merepresentasikan alur kerja aplikasi secara utuh dari sisi client hingga ke layer penyimpanan data.

## 📝 Dokumentasi Api
Dokumentasi API dapat diakses melalui Swagger setelah project dijalankan pada local environment
```bash
  http://127.0.0.1:1323/swagger/index.html
```
Swagger ini memuat seluruh endpoint yang tersedia beserta detail request dan response-nya, sehingga memudahkan dalam proses integrasi maupun pengujian.