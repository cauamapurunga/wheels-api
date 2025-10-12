-- Remove as tabelas na ordem correta de dependência.
DROP TABLE IF EXISTS ordens_servico;

-- Remove a tabela se ela já existir, para garantir um começo limpo.
DROP TABLE IF EXISTS veiculos;

-- Remove a tabela de usuários se ela já existir.
DROP TABLE IF EXISTS users;

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

-- Cria a tabela de usuários.
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
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

-- Cria a tabela para Ordens de Serviço.
CREATE TABLE ordens_servico (
    id SERIAL PRIMARY KEY,
    descricao_servico TEXT NOT NULL,
    custo DECIMAL(10, 2) NOT NULL,
    data_servico DATE NOT NULL,
    veiculo_placa VARCHAR(10) NOT NULL,
    CONSTRAINT fk_veiculo
        FOREIGN KEY(veiculo_placa) 
        REFERENCES veiculos(placa)
        ON DELETE CASCADE
);

-- Insere algumas ordens de serviço de exemplo.
INSERT INTO ordens_servico (descricao_servico, custo, data_servico, veiculo_placa) VALUES
('Troca de óleo e filtro do motor', 350.50, '2024-09-25', 'BRA-2E19'),
('Alinhamento e balanceamento', 180.00, '2024-09-20', 'RGE-5A22'),
('Troca das pastilhas de freio', 450.75, '2024-08-15', 'BRA-2E19');