@echo off
echo ==========================================
echo    API REST de Produtos - Setup
echo ==========================================
echo.

echo 1. Iniciando banco de dados PostgreSQL...
docker-compose up -d

echo.
echo 2. Aguardando banco de dados...
timeout /t 5 /nobreak > nul

echo.
echo 3. Executando aplicacao Go...
go run main.go

pause
