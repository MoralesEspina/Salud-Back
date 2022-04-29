-- MySQL Script generated by MySQL Workbench
-- Thu May 27 19:40:16 2021
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema dassystem
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema dassystem
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `dassystem` DEFAULT CHARACTER SET utf8mb4 ;
USE dassystem;

-- -----------------------------------------------------
-- Table `dassystem`.`job`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dassystem`.`job` (
  `uuid` VARCHAR(36) NOT NULL,
  `name` VARCHAR(100) NULL DEFAULT NULL,
  `description` VARCHAR(100) NULL DEFAULT NULL,
  `isJob` TINYINT NULL DEFAULT '0',
  `isWorkDependency` TINYINT NULL DEFAULT '0',
  `isEspeciality` TINYINT NULL DEFAULT '0',
  PRIMARY KEY (`uuid`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `dassystem`.`person`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dassystem`.`person` (
  `uuid` VARCHAR(36) NOT NULL,
  `fullname` VARCHAR(100) NOT NULL,
  `cui` VARCHAR(70) NULL DEFAULT NULL,
  `partida` VARCHAR(100) NULL DEFAULT NULL,
  `sueldo` DOUBLE NULL DEFAULT NULL,
  `admissionDate` DATE NULL DEFAULT NULL,
  `job` VARCHAR(36) NULL DEFAULT NULL,
  `workdependency` VARCHAR(36) NULL DEFAULT NULL,
  `especiality` VARCHAR(36) NULL DEFAULT NULL,
  `reubication` VARCHAR(36) NULL DEFAULT NULL,
  `collective` VARCHAR(5) NULL DEFAULT NULL,
  `isPublicServer` TINYINT NULL DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  INDEX `person_job` (`job` ASC) ,
  INDEX `fk_person_job1_idx` (`workdependency` ASC) ,
  INDEX `fk_person_job2_idx` (`especiality` ASC) ,
  INDEX `fk_person_job3_idx` (`reubication` ASC) ,
  CONSTRAINT `fk_person_job1`
    FOREIGN KEY (`workdependency`)
    REFERENCES `dassystem`.`job` (`uuid`),
  CONSTRAINT `fk_person_job2`
    FOREIGN KEY (`especiality`)
    REFERENCES `dassystem`.`job` (`uuid`),
  CONSTRAINT `fk_person_job3`
    FOREIGN KEY (`reubication`)
    REFERENCES `dassystem`.`job` (`uuid`),
  CONSTRAINT `person_job`
    FOREIGN KEY (`job`)
    REFERENCES `dassystem`.`job` (`uuid`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `dassystem`.`rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dassystem`.`rol` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `role` VARCHAR(50) NULL DEFAULT NULL,
  `description` VARCHAR(50) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 15
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `dassystem`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dassystem`.`user` (
  `uuid` VARCHAR(36) NOT NULL,
  `username` VARCHAR(100) NOT NULL,
  `password` VARCHAR(200) NOT NULL,
  `rol_id` INT NOT NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE INDEX `username` (`username` ASC) ,
  INDEX `user_rol` (`rol_id` ASC) ,
  CONSTRAINT `user_rol`
    FOREIGN KEY (`rol_id`)
    REFERENCES `dassystem`.`rol` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `dassystem`.`autorization`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dassystem`.`autorization` (
  `uuid` VARCHAR(36) NOT NULL,
  `register` INT NULL DEFAULT NULL,
  `submittedAt` DATETIME NULL DEFAULT NULL,
  `modifiedAt` DATETIME NULL DEFAULT NULL,
  `startdate` DATE NULL DEFAULT NULL,
  `enddate` DATE NULL DEFAULT NULL,
  `resumework` DATE NULL DEFAULT NULL,
  `holidays` INT NULL DEFAULT NULL,
  `totaldays` INT NULL DEFAULT NULL,
  `pendingdays` INT NULL DEFAULT NULL,
  `observation` TEXT NULL DEFAULT NULL,
  `authorizationyear` VARCHAR(4) NULL DEFAULT NULL,
  `person_idperson` VARCHAR(36) NOT NULL,
  `user_uuid` VARCHAR(36) NOT NULL,
  `personnelOfficer` VARCHAR(60) NULL DEFAULT NULL,
  `personnelOfficerPosition` VARCHAR(45) NULL DEFAULT NULL,
  `personnelOfficerArea` VARCHAR(50) NULL DEFAULT NULL,
  `executiveDirector` VARCHAR(60) NULL DEFAULT NULL,
  `executiveDirectorPosition` VARCHAR(45) NULL DEFAULT NULL,
  `executiveDirectorArea` VARCHAR(70) NULL DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  INDEX `autorization_person` (`person_idperson` ASC) ,
  INDEX `autorization_user` (`user_uuid` ASC) ,
  CONSTRAINT `autorization_person`
    FOREIGN KEY (`person_idperson`)
    REFERENCES `dassystem`.`person` (`uuid`),
  CONSTRAINT `autorization_user`
    FOREIGN KEY (`user_uuid`)
    REFERENCES `dassystem`.`user` (`uuid`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `dassystem`.`vacationrequest`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dassystem`.`vacationrequest` (
  `uuid` VARCHAR(36) NOT NULL,
  `register` VARCHAR(45) NULL DEFAULT NULL,
  `submittedAt` DATETIME NULL DEFAULT NULL,
  `modifiedAt` DATETIME NULL DEFAULT NULL,
  `lastYearVacation` INT NULl DEFAULT NULL,
  `vacationYearRequest` INT NULL DEFAULT NULL,
  `lastVacationFrom` DATE NULL DEFAULT NULL,
  `lastVacationTo` DATE NULL DEFAULT NULL,
  `vacationFromDate` DATE NULL DEFAULT NULL,
  `vacationToDate` DATE NULL DEFAULT NULL,
  `hasVacationDay` TINYINT NULL DEFAULT NULL,
  `daysQuantity` INT NULL DEFAULT NULL,
  `observations` VARCHAR(500) NULL DEFAULT NULL,
  `person_uuid` VARCHAR(36) NOT NULL,
  `publicServer_uuid` VARCHAR(36) NOT NULL,
  `user_uuid` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`uuid`),
  INDEX `fk_vacationrequest_person_idx` (`person_uuid` ASC) ,
  INDEX `fk_vacationrequest_user1_idx` (`user_uuid` ASC) ,
  INDEX `fk_vacationrequest_person1_idx` (`publicServer_uuid` ASC) ,
  CONSTRAINT `fk_vacationrequest_person`
    FOREIGN KEY (`person_uuid`)
    REFERENCES `dassystem`.`person` (`uuid`),
  CONSTRAINT `fk_vacationrequest_person1`
    FOREIGN KEY (`publicServer_uuid`)
    REFERENCES `dassystem`.`person` (`uuid`),
  CONSTRAINT `fk_vacationrequest_user1`
    FOREIGN KEY (`user_uuid`)
    REFERENCES `dassystem`.`user` (`uuid`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

--
--ALTER TABLE vacationrequest ADD lastYearVacation int;
--ALTER TABLE vacationrequest ADD vacationYearRequest int;