-- postgres
-- \l+ -> to show show databases
-- sudo -u postgres psql postgres
-- ALTER USER postgres PASSWORD 'root';
--- \c userlogin; -. to select database
CREATE DATABASE userlogin;

ALTER TABLE userprofile ALTER COLUMN phonenumber TYPE varchar(14);
ALTER TABLE userprofile ADD isVerified boolean DEFAULT false;

UPDATE userprofile SET isUserVerified=true WHERE phonenumber='8896726484';


CREATE TABLE userprofile (fullname varchar(100), emailid varchar(100), phonenumber char(13) NOT NULL,PRIMARY KEY (phonenumber));


INSERT INTO userprofile(fullname, emailid, phonenumber) VALUES('amrendra','yamren00@gmail.com',8896726484);
INSERT INTO userprofile(fullname, emailid, phonenumber) VALUES('naveen','menaveenpal@gmail.com',88967264840);
