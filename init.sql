-- Create products table
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

-- Insert seed data
INSERT INTO products (name, description, price, category, stock_quantity) VALUES
('Smartphone Samsung Galaxy S23', 'Smartphone Android com 256GB de armazenamento', 2999.99, 'Eletrônicos', 50),
('Notebook Dell XPS 13', 'Laptop ultrafino com processador Intel i7', 4999.99, 'Computadores', 25),
('Fone de Ouvido Sony WH-1000XM4', 'Fone com cancelamento de ruído ativo', 799.99, 'Áudio', 100),
('Tablet Apple iPad Air', 'Tablet com tela de 10.9 polegadas', 2499.99, 'Eletrônicos', 30),
('Teclado Mecânico Logitech MX Keys', 'Teclado sem fio para produtividade', 349.99, 'Acessórios', 75),
('Mouse Gamer Razer DeathAdder V3', 'Mouse ergonômico para jogos', 299.99, 'Acessórios', 120),
('Monitor LG UltraWide 34"', 'Monitor curvo para trabalho e jogos', 1999.99, 'Monitores', 15),
('Câmera Canon EOS R6', 'Câmera mirrorless profissional', 8999.99, 'Fotografia', 10),
('Headset Gamer HyperX Cloud II', 'Headset com som surround 7.1', 399.99, 'Áudio', 60),
('SSD Samsung 1TB NVMe', 'Armazenamento de alta velocidade', 599.99, 'Armazenamento', 40);
