### Description

Aplikasi Kanban App yang bisa digunakan untuk membuat sebuah _task_ (tugas) dan juga bisa mengelompokkan _task_ tersebut menjadi beberapa bagian.

Fitur-fitur baik dari sisi Rest API dan tampilan (_template_) web adalah sebagai berikut:

- Register user
- Login user
- Logout
- Create Category (Kategori)
- Delete Category
- Create Task (tugas)
- Update Task
- Delete Task
- Move Task (dari satu kategori ke kategori lain)

### Penting untuk mengubah koneksi database lokal menjadi milik anda :

```go
os.Setenv("DATABASE_URL", "postgres://<username>:<password>@localhost:5432/<database_name>") // Ubah dengan credential database postgres di localhost.
```