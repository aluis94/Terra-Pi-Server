# README - Terra-Pi Server #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

* Summary of set up
* Configuration
* Dependencies
* Database configuration
* How to run tests
* Deployment instructions

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact

## Raspberry Pi - install ###

### Install latest version of GoLang on the Raspberyy pi
Tutorial here:
https://www.jeremymorgan.com/tutorials/raspberry-pi/install-go-raspberry-pi/

### Required Golang packages
* github.com/urfave/negroni v1.0.0

* .com/gorilla/mux v1.8.0

* .com/go-co-op/gocron v1.19.0

* .com/jinzhu/gorm v1.9.16

* .com/mattn/go-sqlite3 v1.14.16

* golang.org/x/text v0.8.0

* .com/robfig/cron/v3 v3.0.1

* golang.org/x/sync v0.1.0

* .com/jinzhu/inflection v1.0.0

#### Install twilio
```
sudo pip3 install twilio
```

#### Install SMTP server
```
sudo apt update
sudo apt install ssmtp
sudo apt-get install mailutils

```
#### Configure SMTP server
```
sudo nano /etc/ssmtp/ssmtp.conf
```
root=postmaster

mailhub=smtp.gmail.com:587

hostname=raspberrypi

AuthUser={{email here}}@gmail.com

AuthPass={{TheGmailPassword}}

FromLineOverride=YES

UseSTARTTLS=YES

 pip3 install Adafruit_Python_DHT



