# Gunakan image Golang versi terbaru
FROM golang:1.21-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum terlebih dahulu (untuk optimasi cache Docker)
COPY go.mod go.sum ./

RUN go mod download

# Copy seluruh kode sumber
COPY . .

# Build aplikasi
RUN go build -o main .

# Expose port (default: 8080)
EXPOSE $PORT

# Jalankan aplikasi
CMD ["./main"]