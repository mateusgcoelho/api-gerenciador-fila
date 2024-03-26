CREATE TABLE IF NOT EXISTS pessoas (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(255) NOT NULL,
  telefone VARCHAR(255) NULL UNIQUE,
  cpf VARCHAR(255) NULL UNIQUE,
  data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS usuarios (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  senha VARCHAR(255) NOT NULL,
  codigo_registro VARCHAR(255) UNIQUE,
  permissoes INT NOT NULL DEFAULT 0,
  pessoa_id INT NOT NULL,
  data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE usuarios
ADD CONSTRAINT fk_usuario_pessoa FOREIGN KEY (pessoa_id) REFERENCES pessoas(id);
CREATE TABLE IF NOT EXISTS filas (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(255) NOT NULL,
  senha_atual INT NOT NULL,
  data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS atendimentos (
  id SERIAL PRIMARY KEY,
  senha INT NOT NULL,
  pessoa_id INT NULL,
  responsavel_id INT NOT NULL,
  data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  data_finalizacao TIMESTAMP NULL,
  data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE atendimentos
ADD CONSTRAINT fk_atendimento_pessoa FOREIGN KEY (pessoa_id) REFERENCES pessoas(id);
ALTER TABLE atendimentos
ADD CONSTRAINT fk_atendimento_responsavel FOREIGN KEY (responsavel_id) REFERENCES pessoas(id);
