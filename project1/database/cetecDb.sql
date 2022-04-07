create database cetec;
use cetec;

/*Person*/
create table person(
	id int not null key AUTO_INCREMENT,
  	name varchar(255), 
	age int
);

insert into person(id, name, age) 
VALUES 
(1, "mike", 31),
(2, "John", 20), 
(3, "Joseph", 20);


/*phone*/
create table phone
(
	id int not null key AUTO_INCREMENT,  
	number varchar(255), 
	person_id INT
);

insert into phone(id, person_id, NUMBER) 
VALUES 
(1,1, "444-444-4444"), 
(8,2, "123-444-7777"), 
(3,3, "445-222-1234");


/*address*/
create table address(
	id int not null key auto_increment,  
	city varchar(255), 
	state varchar(255), 
	street1 varchar(255), 
	street2 varchar(255), 
	zip_code varchar(255)
);

insert into address(id ,  city , state , street1 , street2 , zip_code ) 
VALUES 
(1,"Eugene", "OR", "111 Main St", "", "98765"),
(2, "Sacramento", "CA", "432 First St", "Apt 1", "22221"),
(3, "Austin", "TX", "213 South 1st St", "", "78704");


/*address_join*/
create table address_join(
	id int not null key auto_increment,  
	person_id int, 
	address_id int
);
insert into address_join(id, person_id, address_id) 
VALUES 
(1,1,3),
(2,2,1),
(3,3,2);



DROP PROCEDURE if EXISTS `GetPersonByID`;
delimiter //
CREATE PROCEDURE `GetPersonByID`(
	IN `Person_ID` int
)
BEGIN
	SELECT per.NAME, ph.number, adr.city, adr.state, adr.street1,
	adr.street2, adr.zip_code 
	FROM person AS per
	INNER JOIN phone AS ph ON ph.person_id=per.id
	INNER JOIN address_join AS adj ON adj.person_id=per.id
	INNER JOIN address AS adr ON adr.id=adj.address_id
	WHERE per.id=Person_ID;
END//
delimiter ;

-- CALL GetPersonByID(1);


DROP PROCEDURE if EXISTS `AddPersonInfo`;
delimiter //
CREATE PROCEDURE `AddPersonInfo`(
    IN `Person_Name` VARCHAR(255),
    IN `Person_PhoneNumber`  VARCHAR(255),
    IN `Person_City` VARCHAR(255),
    IN `Person_State` VARCHAR(255),
    IN `Person_Street1` VARCHAR(255),
    IN `Person_Street2` VARCHAR(255),
    IN `Person_ZipCode` VARCHAR(255)
)
BEGIN
	DECLARE CODE CHAR(5) DEFAULT '00000';
	DECLARE msg TEXT;
	DECLARE personID,addressID INT;

	DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
	BEGIN
		GET DIAGNOSTICS CONDITION 1
		CODE = RETURNED_SQLSTATE, msg=MESSAGE_TEXT;
	END;

	START transaction;

		INSERT INTO person(NAME)
		VALUES(Person_Name);
		
		SELECT MAX(id) INTO personID
		FROM person;
				
		INSERT INTO phone(person_id, number)
		VALUES(personID, Person_PhoneNumber);
		
		INSERT INTO address(city, state, street1, street2, zip_code)
		VALUES(Person_City, Person_State, Person_Street1, Person_Street2, Person_ZipCode);

		SELECT MAX(id) INTO addressID
		FROM address;
		
		INSERT INTO address_join(person_id, address_id)
		VALUES(personID, addressID);
	
	if CODE = '00000' then
		SELECT 1;
		COMMIT;
	else
		SELECT 0;
		SELECT msg;
		ROLLBACK;
	END if;

END//
delimiter ;
