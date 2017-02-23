# Backup restore docker persistent with cron

``` yml 
services:
  mongodb:
    image: mongo:3.4
  backups:
    image: mikeifomin/sweetheart-data
    envirnoment:
      EVERY = 1 day
      CRON = * 8 8 8 *
      B2_KEY = 
      B2_ACCOUNT = 
       
    volumes:
     - user_data:/bkp/dir
```

will backup data
into proj	
