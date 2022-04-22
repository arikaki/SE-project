# Project - KORA

## Sprint 2
 

We have accomplished in the project according to the issue/stories conveyed (closed stories in sprint_2) 
#### ISSUE 1 - Added Signin and Signup Page
#### ISSUE 2 - Added Category page for New Users
#### ISSUE 3 - Created new collections and dummy data 
#### ISSUE 4 - Implemented User Session Management 
#### ISSUE 5 - Implemented Fetch User details for profile page 
#### ISSUE 6 - Added test cases for frontend
#### ISSUE 7 - Added test cases for backend


# Backend Development server
Run the below commands to start the development server.

> `cd ../backend`

> `go run main.go `

![image](https://user-images.githubusercontent.com/38340989/156869950-cc8f8cd0-4a3a-49c9-b4bf-9dbcd394a263.png)

 
> # Backend
API :: 
The below API's are used 
1. Implemented Insert users 
2. API to Fetch User Details by username
3. API to post dummy answers
4. API to post dummy questions
5. API to post dummy users
6. API to Login
7. API to Logout

## Added three collections

![image](https://user-images.githubusercontent.com/38340989/156870019-b477427b-c660-4e9a-a335-70b1a70dc5ae.png)


## 1. API to Fetch User Details by username
The /fetch-user API is used to fetch a particular user by username as requested from the frontend Webpage.

![image](https://user-images.githubusercontent.com/38340989/156868749-f2084c0a-ca34-4f00-83fe-a2f2dbf7a888.png)

Fetches the user details from cookies stored after the user has loggedin
## 2.  API to post dummy answers
The /dummyanswer API is used to post dummy answers to database

![image](https://user-images.githubusercontent.com/38340989/156868797-f6a4af88-8e7f-4e14-b4d6-370aafe7e9d5.png)

Inserted Dummy answers in Database
## 3. API to Insert user
The /api/user/insert API is used to insert users into database.

![image](https://user-images.githubusercontent.com/38340989/156868672-0ba42052-ca28-41f2-b936-453983d8c802.png)

The data is inserted in Database

![image](https://user-images.githubusercontent.com/38340989/156868704-28ac84b4-1f10-4abb-8c98-5e3a87574b99.png)

## 4.  API to post dummy questions
The /dummyquestion is used to insert questions into database.

![image](https://user-images.githubusercontent.com/38340989/156869262-ba802db0-fcd9-4f97-8f83-f5abc30ce487.png)


## 5.  API to post dummy users
The /dummyuser is used to insert users into database.

![image](https://user-images.githubusercontent.com/38340989/156869284-d05dccbb-7a7c-4a63-975e-975367bee2b7.png)

## 6.  API to Login
The /login is used to login using username and password.

![image](https://user-images.githubusercontent.com/38340989/156869409-de72ba78-a393-4ef2-90d5-b3b37cb154a0.png)

## 7.  API to Logout
The /logout is used to logout user.

![image](https://user-images.githubusercontent.com/38340989/156869643-428aabc7-8578-48ba-bfea-223a1b4f46db.png)

### Testcases for backend:
run `go test`

successfully ran the test cases for backend

![testcases](https://user-images.githubusercontent.com/38340989/156869740-80115618-a8d2-4f1a-bccf-97ed74fdc4f8.jpeg)

