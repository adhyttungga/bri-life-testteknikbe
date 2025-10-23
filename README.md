Go Test Teknis BE
===

### Description
Microservice with Go, RESTful API, and implement Authorization and Authentication.

### Prerequisites
- Go (version 1.25 or newer is recommended)
- MySQL
- Docker

### Getting Started
These instructions will get a copy of the project up and running on your local machine for development and testing purpose

#### Install Dependencies
First, clone the repository and navigate into the project directory:
```bash
git clone https://github.com/adhyttungga/bri-life-testteknisbe.git
cd bri-life-testteknisbe
```

### Running the Application
#### Run the Server
To start the gRPC server, execute the following command:
```bash
./scripts/start.sh
```
The server will start as docker container and listen on a specified port (e.g., ":8080")

===

### *Note
1. Alasan menggunakan clean architecture:
Saya menggunakan Clean Architecture karena memudahkan dalam membangun aplikasi yang maintainable, testable, dan scalable. Hal tersebut diperoleh dengan memisahkan masalah ke dalam beberapa layer. Clean Architecture menjadikan logika bisnis independen dari faktor lain seperti database atau framework UI.  Pemisahan tersebut mengurangi dampam modifikasi di satu area dengan area lain dan meningkatkan kemampuan adaptasi terhadap kebutuhan bisnis.

2. Alasan menggunakan RESTful API:
Alasan dibalik penggunaam RESTful API karena popularitasnya dalam pengembangan web services dan beberapa keuntungan lain sebagai berikut:
- Scalability and Flexibility
- Simplicity and Ease of Use
- Statelessness (server tidak menyimpan context dari client)
- Caching Capabilities
- Security