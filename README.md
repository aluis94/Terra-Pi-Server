# README #

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

### Raspberry Pi - install ###

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
