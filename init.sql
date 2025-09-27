-- Remove a tabela se ela já existir, para garantir um começo limpo.
DROP TABLE IF EXISTS veiculos;

-- Cria a tabela com a estrutura correta.
CREATE TABLE veiculos (
    id SERIAL PRIMARY KEY,
    placa VARCHAR(10) NOT NULL UNIQUE,
    marca VARCHAR(50) NOT NULL,
    modelo VARCHAR(100) NOT NULL,
    ano_fabricacao INT NOT NULL,
    cor VARCHAR(30) NOT NULL,
    nome_proprietario VARCHAR(255) NOT NULL
);

-- Insere os 11 registros da sua imagem.
-- O 'id' não é especificado aqui porque a coluna SERIAL o gera automaticamente.
INSERT INTO veiculos (placa, marca, modelo, ano_fabricacao, cor, nome_proprietario) VALUES
('BRA-2E19', 'Volkswagen', 'Gol', 2020, 'Branco', 'Carlos Eduardo Souza'),
('RGE-5A22', 'Fiat', 'Mobi', 2023, 'Vermelho', 'Ana Clara Ferreira'),
('JKL-1234', 'Chevrolet', 'Onix', 2022, 'Prata', 'Pedro Henrique Lima'),
('XYZ-0F05', 'Toyota', 'Corolla', 2024, 'Preto', 'Mariana Costa Alves'),
('MNO-5678', 'Hyundai', 'HB20', 2021, 'Cinza', 'Lucas Gabriel Martins'),
('PQR-1G33', 'Jeep', 'Renegade', 2024, 'Preto', 'Cauã Mapurunga'),
('STU-9012', 'Honda', 'Civic', 2019, 'Branco', 'Rafael Oliveira Santos'),
('VWX-3H44', 'Ford', 'Ka', 2018, 'Prata', 'Beatriz Almeida'),
('YZA-3456', 'Renault', 'Kwid', 2022, 'Laranja', 'Guilherme Pereira'),
('BCD-7I88', 'Nissan', 'Kicks', 2024, 'Preto', 'Laura Fernandes Rocha'),
('RIO-2A18', 'Chevrolet', 'Onix', 2024, 'Preto', 'Irineu Martins');