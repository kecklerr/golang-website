# golang-website
golang web site and test
trying to get my hands about golang and proper setup.

docker directory
- Dockerfile to build the image
- basic image
- uses centos 7 to run the go app executable

openshift  directory
- config yaml file to deploy the docker build  to openshift
- currently using  CRC (OKD)

main directory
- main.go for app 
- using gorilla mux to do routing

Futrure 
- put in CICD steps for all this so I dont have to do it manually