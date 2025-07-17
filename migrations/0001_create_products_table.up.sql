-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    category VARCHAR(100),
    stock_quantity INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, description, price, category, stock_quantity) VALUES
('Smartphone Samsung Galaxy S23', 'Smartphone Android com 256GB de armazenamento', 2999.99, 'Eletrônicos', 50),
('Notebook Dell XPS 13', 'Laptop ultrafino com processador Intel i7', 4999.99, 'Computadores', 25),
('Smart TV LG OLED 55"', 'TV 4K com tecnologia OLED e suporte a HDR', 3999.99, 'Eletrônicos', 30),
('Fone de Ouvido Sony WH-1000XM4', 'Fones de ouvido com cancelamento de ruído', 1499.99, 'Acessórios', 100),
('Cadeira Gamer DXRacer', 'Cadeira ergonômica para jogos com apoio lombar', 899.99, 'Móveis', 20),
('Console PlayStation 5', 'Console de videogame com jogos exclusivos', 4999.99, 'Jogos', 15),
('Monitor Dell UltraSharp 27"', 'Monitor 4K com alta precisão de cores', 2499.99, 'Computadores', 40),
('Impressora HP LaserJet Pro M404dn', 'Impressora a laser rápida e eficiente', 1299.99, 'Periféricos', 10),
('Mouse Logitech MX Master 3', 'Mouse sem fio com várias funcionalidades', 499.99, 'Acessórios', 200),
('Teclado Mecânico Corsair K95 RGB', 'Teclado mecânico com retroiluminação RGB', 799.99, 'Acessórios', 150),
('Câmera Canon EOS M50', 'Câmera mirrorless com lente intercambiável', 3499.99, 'Fotografia', 30),
('Smartwatch Apple Watch Series 7', 'Relógio inteligente com monitoramento de saúde', 3999.99, 'Wearables', 60),
('Caixa de Som JBL Flip 5', 'Caixa de som Bluetooth portátil e resistente à água', 699.99, 'Áudio', 80),
('Tablet Samsung Galaxy Tab S7', 'Tablet Android com S Pen inclusa', 2799.99, 'Tablets', 45),
('Roteador TP-Link Archer AX50', 'Roteador Wi-Fi 6 com alta velocidade e cobertura', 599.99, 'Redes', 70);