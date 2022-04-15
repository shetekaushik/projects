create database StoreShortenURL;
use StoreShortenURL;

/*Person*/
create table MapShortAndLongURL(
	id int not null primary key AUTO_INCREMENT,
  	LongURL TEXT, 
	ShortURL TEXT
);


DROP PROCEDURE if EXISTS `GetLongURL`;
delimiter //
CREATE PROCEDURE `GetLongURL`(
	IN `ShortenURL` TEXT
)
BEGIN
	SELECT msl.LongURL
	FROM MapShortAndLongURL AS msl
	WHERE msl.ShortURL=ShortenURL;
END//
delimiter ;

-- CALL GetPersonByID(1);


DROP PROCEDURE if EXISTS `AddLongAndShortenURL`;
delimiter //
CREATE PROCEDURE `AddLongAndShortenURL`(
	IN `sURL` TEXT,
	IN `lURL` TEXT)
BEGIN
	DECLARE CODE CHAR(5) DEFAULT '00000';
	DECLARE msg TEXT;

	DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
	BEGIN
		GET DIAGNOSTICS CONDITION 1
		CODE = RETURNED_SQLSTATE, msg=MESSAGE_TEXT;
	END;

	START transaction;
		INSERT INTO MapShortAndLongURL(ShortURL, LongURL)
		VALUES(sURL,lURL);

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



DROP PROCEDURE if EXISTS `isExist`;
delimiter //
CREATE PROCEDURE `isExist`(
	IN `lURL` TEXT
)
BEGIN
	DECLARE longURLAvail int DEFAULT 0;
	select distinct 1 into longURLAvail
	from MapShortAndLongURL
	where LongURL=lURL;

	select longURLAvail;

	if longURLAvail=1 then
		select distinct ShortURL
		from MapShortAndLongURL
		where LongURL=lURL;
	end if;

END//
delimiter ;
