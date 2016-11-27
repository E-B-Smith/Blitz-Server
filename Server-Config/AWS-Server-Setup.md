
AWS Instance Set Up
===================
Root user is 'ubuntu'.

3.  Create 'sysadmin' sudo-capable user.

        sudo adduser sysadmin  --shell /bin/bash  --disabled-password
        sudo usermod -a -G adm sysadmin
        sudo usermod -a -G admin sysadmin

        sudo groupadd signers
        sudo usermod -a -G signers sysadmin

4.  Update su-doers file

        sysadmin   ALL=(ALL:ALL) NOPASSWD:ALL

6.  Update sshd settings

5.  Add key to ~/.ssh/authorized_keys

7.  Install time daemon:  sudo apt-get install ntp

8.  sudo  ln -fsv /usr/share/zoneinfo/US/Pacific /etc/localtime

2.  Set host name: blitzhere
    - Edit /etc/hostname
    - Edit /etc/hosts
    - `sudo service hostname restart`

1.  Set up external DNS: blitzhere.com / www.blitzhere.com.

11. Install postgres

        [Instal Postgres](http://tecadmin.net/install-postgresql-server-on-ubuntu/#)

        sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" >> /etc/apt/sources.list.d/pgdg.list'
        wget -q https://www.postgresql.org/media/keys/ACCC4CF8.asc -O - | sudo apt-key add -

        sudo apt-get update
        sudo apt-get install postgresql postgresql-contrib

    Configure postgres user info.

    pg_ident.conf:

        # MAPNAME       SYSTEM-USERNAME         PG-USERNAME
        adminmap        ubuntu                  postgres
        adminmap        sysadmin                postgres
        adminmap        blitzhere               postgres
        adminmap        blitzlabs               postgres

        blitzmap        blitzhere               blitzlabs
        blitzmap        blitzhere               blitzhere


    pg_hba.conf:

        # Database administrative login by Unix domain socket
        local   all             postgres                                peer map=adminmap

        # TYPE  DATABASE        USER            ADDRESS                 METHOD

        # "local" is for Unix domain socket connections only
        local   blitzlabs       blitzlabs                               peer map=blitzmap
        local   blitzhere       blitzhere                               peer map=blitzmap

9.  Install nginx:      sudo apt-get nginx

10. Configure nginx:

        sudo ln -svi /etc/nginx/sites-available/BlitzHere-nginx.conf  \
            /etc/nginx/sites-enabled/BlitzHere-nginx.conf

12. Create a blitzhere account as a regular user.

        sudo adduser blitzhere  --shell /bin/bash  --disabled-password

13. Deploy, files, make database, run build.
