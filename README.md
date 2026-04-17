# 📝 Go Task Manager CLI  

A simple and efficient Command-Line Task Manager built using Golang.  
This project helps users manage daily tasks directly from the terminal with persistent storage using a JSON file.

---

## 🚀 Features  

- ➕ Add new tasks  
- 📋 View all tasks  
- ❌ Delete tasks  
- 💾 Persistent storage using JSON file  
- ⚡ Fast and lightweight CLI application  

---

## 📁 Project Structure  
```
go-task-manager/
│── main.go        # Entry point of the application  
│── task.go        # Task model/structure  
│── storage.go     # File handling & business logic  
│── go.mod         # Go module file  
│── tasks.json     # Auto-created database file  
```
---

## 🛠️ Tech Stack  

- Language: Golang  
- Storage: JSON File  

Concepts Used:
- Structs  
- File Handling  
- JSON Encoding/Decoding  
- CLI Arguments  

---

## ⚙️ Installation & Setup  

### 1. Clone the Repository  
```bash
git clone https://github.com/your-username/go-task-manager.git
cd go-task-manager

```
## 2. Install Dependencies
```
go mod tidy
```
## 3. Run the Application
```
go run .
```
##▶️ Usage
➕ Add Task
```
go run . add "Learn Go"
```

##📋 List Tasks
```
go run . list
```
##❌ Delete Task
```
go run . delete 1
```
##📌 Example Output
```
1. Learn Go [❌]
2. Build Project [❌]
```
## 📈 Future Enhancements
✅ Mark tasks as completed
📅 Add deadlines & priorities
🌐 Convert into REST API using Gin
🗄️ Integrate database (MongoDB / MySQL)
🎨 Build frontend (React / Next.js)

## 🤝 Contributing
Contributions are welcome!

Steps:

- Fork the repository
- Create your feature branch
- Commit your changes
- Push to the branch
- Open a Pull Request

## 📄 License
This project is open-source and available under the MIT License.

## ⭐ Support
If you like this project, give it a star on GitHub!
```
If you want next level 🔥  
I can give:
- :contentReference[oaicite:0]{index=0}  
- :contentReference[oaicite:1]{index=1}  
- Or :contentReference[oaicite:2]{index=2}
```
