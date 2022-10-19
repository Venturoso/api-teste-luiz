-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 19-Out-2022 às 21:25
-- Versão do servidor: 10.4.22-MariaDB-log
-- versão do PHP: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `api_luiz`
--
CREATE DATABASE IF NOT EXISTS `api_luiz` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
USE `api_luiz`;

-- --------------------------------------------------------

--
-- Estrutura da tabela `produtos`
--

CREATE TABLE `produtos` (
  `id` int(11) NOT NULL,
  `nome` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `descricao` text COLLATE utf8_unicode_ci NOT NULL,
  `valor` float NOT NULL,
  `ativo` varchar(1) COLLATE utf8_unicode_ci DEFAULT 'S'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Extraindo dados da tabela `produtos`
--

INSERT INTO `produtos` (`id`, `nome`, `descricao`, `valor`, `ativo`) VALUES
(1, 'hello', 'Ola', 22, '1'),
(3, 'sadss', 'Anti ferrugem', 33.7, '2'),
(5, 'WD40', 'Anti ferrugem', 33.5, '1'),
(7, 'hello', 'Ola', 22, '1'),
(9, 'sadss', 'Anti ferrugem', 33.5, '1'),
(11, 'sadss', 'Anti ferrugem', 33.5, '1'),
(13, 'sadss', 'Anti ferrugem', 33.5, '1'),
(15, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(17, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(19, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(21, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(23, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(25, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(27, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(29, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(31, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(37, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(39, 'ds', 'Anti ferrugem', 33.5, '1'),
(41, '', 'Anti ferrugemss', 33.5, '1'),
(43, '', 'Anti ferrugemss', 33.5, '1'),
(45, '', 'Anti ferrugemss', 33.5, '1'),
(47, '', 'Anti ferrugemss', 33.5, '1'),
(49, 'sadss', 'Anti ferrugemss', 33.5, '1'),
(51, 'sadssss', 'Anti ferrugemss', 33.5, '1'),
(53, 'sadssss', 'Anti ferrugemss', 33.5, '1'),
(55, 'ds', 'Anti ferrugem', 33.5, '1'),
(57, 'sadssss', 'Anti ferrugemss', 33.5, '1'),
(59, 'sadssss', 'Anti ss', 33, 'N'),
(61, 'sadssss', 'Anti ferrugemss', 33.5, '0'),
(63, 'sadssss', 'Anti ferrugemss', 33.5, '0'),
(65, 'sadssss', 'Anti ferrugemss', 33.5, 'N'),
(67, 'sadssss', 'Anti ferrugemss', 33.5, 'N'),
(69, 'sadssss', 'Anti ferrugemss', 33.5, 'N'),
(71, 'sadssss', 'Anti ferrugemss', 33.5, 'N'),
(73, 'sadssss', 'Anti ferrugemss', 33.5, 'N'),
(75, 'sadssss', 'Anti ferrugemss', 33.5, 'N');

--
-- Índices para tabelas despejadas
--

--
-- Índices para tabela `produtos`
--
ALTER TABLE `produtos`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT de tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `produtos`
--
ALTER TABLE `produtos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=77;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
