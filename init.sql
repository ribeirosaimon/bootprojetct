CREATE USER 'myuser'@'%' IDENTIFIED BY 'mypass';
GRANT ALL PRIVILEGES ON shortener.* TO 'myuser'@'%';
FLUSH PRIVILEGES;