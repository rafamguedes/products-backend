#!/usr/bin/env pwsh

Write-Host "==========================================" -ForegroundColor Green
Write-Host "    API REST de Produtos - Setup" -ForegroundColor Green
Write-Host "==========================================" -ForegroundColor Green
Write-Host ""

Write-Host "1. Iniciando banco de dados PostgreSQL..." -ForegroundColor Yellow
docker-compose up -d

Write-Host ""
Write-Host "2. Aguardando banco de dados..." -ForegroundColor Yellow
Start-Sleep -Seconds 5

Write-Host ""
Write-Host "3. Executando aplicação Go..." -ForegroundColor Yellow
go run main.go
