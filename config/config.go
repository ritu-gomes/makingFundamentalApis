package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host string
	Port int
	Name string
	User string
	Password string
	EnableSSLMode bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int 
	JwtSecretKey string
	DB *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variable", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == ""{
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be a number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("jwt secret key is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("HOST")
	if dbHost == "" {
		fmt.Println("db Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("PORT")
	if dbPort == "" {
		fmt.Println("db Port is required")
		os.Exit(1)
	}

	dbport, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be a number")
		os.Exit(1)
	}

	dbName := os.Getenv("NAME")
	if dbName == "" {
		fmt.Println("db Name is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("PASSWORD")
	if dbPassword == "" {
		fmt.Println("db Password is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("USER")
	if dbUser == "" {
		fmt.Println("db user is required")
		os.Exit(1)
	}

	enableSSLMode := os.Getenv("ENABLE_SSL_MODE")
	enableSSLmode, err := strconv.ParseBool(enableSSLMode)
	if err != nil {
		fmt.Println("invalid enable ssl mode value")
		os.Exit(1)
	}

	dbConfiguration := &DBConfig{
		Host: dbHost,
		Port: int(dbport),
		Name: dbName,
		User: dbUser,
		Password: dbPassword,
		EnableSSLMode: enableSSLmode,
	}

	configurations  = &Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: int(port),
		JwtSecretKey: jwtSecretKey,
		DB: dbConfiguration,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}