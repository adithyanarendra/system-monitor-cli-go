# 📊 System Resource Monitor CLI

A terminal-based system resource monitor built with Go and `tview`, displaying CPU, memory, disk, and network usage with ASCII bar graphs.

---

## 🚀 Features
- **Real-time system monitoring** (updates every 2 seconds)
- **CPU, Memory, Disk, and Network stats**
- **ASCII bar graphs** for visual representation
- **Minimal dependencies**

---

## 🛠 Installation & Usage

### 1️⃣ **Clone the Repository**
```sh
git clone https://github.com/adithyanarendra/system-monitor-cli-go.git
cd system-monitor-cli-go
```

### 2️⃣ **Run the Program**
#### ✅ Direct Execution (Recommended)
```sh
go run main.go
```

#### ✅ If You Have Multiple Files (e.g., `stats.go`)
```sh
go run .
```

---

## 🔨 Building a Distributable Binary

### **For Your OS**
```sh
go build -o system-monitor
./system-monitor  # Run the binary
```

### **Cross-Compiling for Other Platforms**

#### 🔹 Linux (64-bit)
```sh
GOOS=linux GOARCH=amd64 go build -o system-monitor-linux
```

#### 🔹 Windows (64-bit)
```sh
GOOS=windows GOARCH=amd64 go build -o system-monitor.exe
```

#### 🔹 macOS (Apple Silicon & Intel)
```sh
GOOS=darwin GOARCH=arm64 go build -o system-monitor-mac
GOOS=darwin GOARCH=amd64 go build -o system-monitor-mac-intel
```

---

## 📦 Dependencies
- **[tview](https://github.com/rivo/tview)** (Terminal UI)
- **[gopsutil](https://github.com/shirou/gopsutil)** (System stats, if used in `stats.go`)

Install missing dependencies with:
```sh
go mod tidy
```

---

## 🎯 Future Enhancements
- ✅ More detailed graphs
- ✅ Support for multiple themes
- ✅ Export logs to file

---

## 📝 License
MIT License - Feel free to modify and distribute!



