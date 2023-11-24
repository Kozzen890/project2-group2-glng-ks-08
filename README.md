# project2-group2-glng-ks-08


## PROJECT 2 - KELOMPOK 2 - HACKTIV8 - MBKM - GOLANG FOR BACK-END




## Team 2-KS08 Contributors

- Hafidzurrohman Saifullah (GLNG-KS-08-02) - GitHub@Hafidzurr - Golang For Back-End - Universitas Gunadarma
- Sherly Fauziyah Syaharani (GLNG-KS-08-018) - GitHub@Sherlyfauz - Golang For Back-End - Universitas Merdeka Malang
- Timotius Winsen Bastian (GLNG-KS-08-016) - GitHub@Kozzen890 - Golang For Back-End - Universitas Dian Nuswantoro


## API URL 

https://project2-group2-glng-ks-08-production.up.railway.app/

## Postman Documentation
https://documenter.getpostman.com/view/27920824/2s9YeD7Czk

## System Requirement
- Golang.
- Postgres SQL.
- Gorm.
## Installation Local

-  Open terminal or command prompt

```bash
git clone https://github.com/Kozzen890/project2-group2-glng-ks-08
cd project2-group2-glng-ks-08
go mod tidy
```

- Setting Database

    a. Create database in postgres SQL with name mygram or you can change whats name you like, but coution here you must change database name in db.go too.

    b. Go to db.go, comment line code from dns = fmt.Sprintf - dbname, dbPort) and uncomment line code dsn = "host=host....

    c. Change to your own db credential in db.go.



- Run
```bash
go run main.go
```

## Installation & Deployment to Railway

-  Open terminal or command prompt
```bash
git clone https://github.com/Kozzen890/project2-group2-glng-ks-08
cd project2-group2-glng-ks-08
go mod tidy
```

-  Push into Your New Repo
    a. Create a New Repository in Your Github Account

    b. Change the Remote URL
    ```bash
    git remote set-url origin https://github.com/new_user/new_repo.git
    ```

    c. Push to the New Repository
    ```bash
    git push -u origin master or your name for repo banch
    ```

- Create Account Railway using your github Account and Login

    ###### Create `New Project` -> Choose `Deploy from github Repo` -> Choose `Your Repo Name` -> Wait Deploying Untill Getting Error

- Adding Postgres SQL into Your Project
    ###### Choose `New` -> Choose `Database` -> Choose `Postgres SQL` -> Wait Deploying Untill Getting Error


- env & .gitignore
    ###### a. Edit `.env` in local or there is no, you can create `.env` and adding : 

    ```
        PGHOST=**your_db_host**
        PGPORT=**your_db_port**
        PGUSER=**your_db_user**
        PGPASSWORD=**your_db_password**
        PGDATABASE=**your_db_name**
    ```

    ######  b. Change Variable with your own variable getting from Railway, to see your variable, you can see them in your `postgres SQL` and go to `variables`.

    ######  c. Edit `.gitignore` in local or there is no, you can create `.gitignore` and adding :

    ```
    .env
    ```

    ###### d. Push your changes
    
-  Adding `.env` Variables
    ###### a. Adding whole variable on `.env`, into `your project` and go to `variables`, adding in one by one.
    ###### b. Now wait deploying and after that you can create your own domain.
