# GoXploitDB

A simple GoLang microservice to query and get data from [Exploit-DB](https://www.exploit-db.com/). The microservice achieves this goal by initiating asynchronous scans on the [Exploit-DB GitHub repository](https://github.com/offensive-security/exploitdb) and by populating an Sqlite local DB. Thereby enabling the users to search for CVE's based on CVE ID, Platform, Type or Year. 

## Prerequisites

- GoLang (version >= 1.17.8)
- Sqlite (Optional)
- Docker (Optional)


## Run GoXploitDB [using Go]

GoXploitDB requires  **go1.17 or higher**  to run successfully. Run the following commands to build the latest version
```sh
git clone https://github.com/vishnusomank/GoXploitDB.git
cd GoXploitDB
go mod tidy
go build -o GoXploitDB main.go 
```
To run the program use
```sh
./GoXploitDB
```

## Run GoXploitDB [using Docker]

### 1. Local Build

```sh
git clone https://github.com/vishnusomank/GoXploitDB.git
cd GoXploitDB
docker build -t goxploitdb . 
```
To run the program use
```sh
docker run -d -p 8080:8080 goxploitdb
```

### 2. Pre-Built Docker image

Run the microservice using
```sh
docker run -d -p 8080:8080 knoxuser/goxploitdb
```
> **Note:** Use `knoxuser/goxploitdb:stable` to avoid initial DB loading time 

## Run GoXploitDB on Kubernetes
1. Using Local copy
    ```sh
    kubectl apply -f ./k8s-goxploitdb.yaml
    ```
2. Using Raw content from GitHub
    ```sh
    kubectl apply -f https://raw.githubusercontent.com/vishnusomank/GoXploitDB/main/k8s-goxploitdb.yaml
    ```


## Features
The server exposes 5 APIs 

- `/goxploitdb/api/v1/cve/:cve`
- `/goxploitdb/api/v1/platform/:platform`
- `/goxploitdb/api/v1/type/:type` 
- `/goxploitdb/api/v1/platforms`
- `/goxploitdb/api/v1/types` 


| ENDPOINT | TYPE  | DATA | EXPLANATION |
|--|--|--|--|
| `/goxploitdb/api/v1/cve/:cve` | GET  | cve_id | Returns details on the CVE as key-value pairs of `AUTHOR` , `CVE` , `EDB-ID` , `ID` , `PLATFORM` , `TITLE` , `TYPE` and `URL` |
|`/goxploitdb/api/v1/platform/:platform`| GET  | platform_name | Displays exploit details for a specific platform (eg, windows, linux etc)|
|`/goxploitdb/api/v1/type/:type` | GET  | type_name | Displays exploit details for a specific type (eg, webapps, remote etc)|
|`/goxploitdb/api/v1/platforms`| GET  | NIL | Displays usable values for platforms|
|`/goxploitdb/api/v1/types`| GET | NIL | Displays usable values for types|


## API Usage

 1. List all usable platforms

	```json
	curl -s "http://localhost:8080/goxploitdb/api/v1/platforms" | jq
    {
    "Total usable platform values": [
        "AIX",
        "ALPHA",
        "ANDROID",
        "ASP",
        "ASPX",
        "BSD",
        "CFM",
        "CGI",
        "FREEBSD",
        "HARDWARE",
        "HP-UX",
        "IOS",
        "JAVA",
        "JSON",
        "JSP",
        "LINUX",
        "LINUX_X86",
        "LUA",
        "MACOS",
        "MULTIPLE",
        "NODEJS",
        "OPENBSD",
        "OSX",
        "PERL",
        "PHP",
        "PYTHON",
        "RUBY",
        "SOLARIS",
        "UNIX",
        "WATCHOS",
        "WINDOWS",
        "WINDOWS_X86",
        "WINDOWS_X86-64",
        "XML"
    ]
    }
	```
2. List all usable types

    ```json
    curl -s "http://localhost:8080/goxploitdb/api/v1/types" | jq
    {
    "Total usable type values": [
        "LOCAL",
        "WEBAPPS",
        "DOS",
        "REMOTE"
    ]
    }
    ```	                    

3. List CVE data
    - List a specific CVE detail
	
        ```json
        curl -s "http://localhost:8080/goxploitdb/api/v1/cve/CVE-2022-29548" | jq
        {
        "AUTHOR": "cxosmo",
        "CVE": "CVE-2022-29548",
        "EDB-ID": "50970",
        "ID": 2203,
        "PLATFORM": "PHP",
        "TITLE": "WSO2 Management Console (Multiple Products) - Unauthenticated Reflected Cross-Site Scripting (XSS)",
        "TYPE": "WEBAPPS",
        "URL": "https://www.exploit-db.com/exploits/50970"
        }
        {
        "Total Records": 1
        }
        ``` 
    - List all CVE data for a year
        ```json
        curl -s "http://localhost:8080/goxploitdb/api/v1/cve/2022-" | jq
        {
        "AUTHOR": "Stephen Chavez & Robert Willis",
        "CVE": "CVE-2022-27226",
        "EDB-ID": "50832",
        "ID": 160,
        "PLATFORM": "HARDWARE",
        "TITLE": "iRZ Mobile Router - CSRF to RCE",
        "TYPE": "REMOTE",
        "URL": "https://www.exploit-db.com/exploits/50832"
        }
        <------------------------------>SNIP<------------------------------>
        {
        "AUTHOR": "Tomer Peled",
        "CVE": "CVE-2022-24562",
        "EDB-ID": "50974",
        "ID": 2582,
        "PLATFORM": "WINDOWS",
        "TITLE": "IOTransfer V4 – Remote Code Execution (RCE)",
        "TYPE": "REMOTE",
        "URL": "https://www.exploit-db.com/exploits/50974"
        }
        {
        "Total Records": 63
        }
        ```
    - List all CVEs present in the DB
        ```json
        curl -s "http://localhost:8080/goxploitdb/api/v1/cve/cve" | jq
        {
        "AUTHOR": "Kristian Erik Hermansen <kristian.hermansen@gmail.com>",
        "CVE": "CVE-2013-4011",
        "EDB-ID": "28507",
        "ID": 1,
        "PLATFORM": "AIX",
        "TITLE": "IBM AIX 6.1 / 7.1 local root privilege escalation",
        "TYPE": "LOCAL",
        "URL": "https://www.exploit-db.com/exploits/28507"
        }
        <------------------------------>SNIP<------------------------------>
        {
        "AUTHOR": "Jonas Lejon",
        "CVE": "CVE-2018-10653",
        "EDB-ID": "47951",
        "ID": 2688,
        "PLATFORM": "XML",
        "TITLE": "Citrix XenMobile Server 10.8 - XML External Entity Injection",
        "TYPE": "WEBAPPS",
        "URL": "https://www.exploit-db.com/exploits/47951"
        }
        {
        "Total Records": 1916
        }
        ```
4. List all exploit details for a platform
	
    ```json
    curl -s "http://localhost:8080/goxploitdb/api/v1/platform/linux" | jq
    
    {
    "AUTHOR": "c0ntex",
    "CVE": "CVE-2011-2702",
    "EDB-ID": "20167",
    "ID": 561,
    "PLATFORM": "LINUX",
    "TITLE": "eGlibc Signedness Vulnerability",
    "TYPE": "DOS",
    "URL": "https://www.exploit-db.com/exploits/20167"
    }
    <------------------------------>SNIP<------------------------------>
    {
    "AUTHOR": "d7x",
    "CVE": "N/A",
    "EDB-ID": "46249",
    "ID": 754,
    "PLATFORM": "LINUX_X86",
    "TITLE": "MySQL User-Defined (Linux) x32 / x86_64 sys_exec function local privilege escalation exploit",
    "TYPE": "LOCAL",
    "URL": "https://www.exploit-db.com/exploits/46249"
    }
    {
    "Total Records": 194
    }
    ``` 
5. List all exploit details for a type
	
    ```json
    curl -s "http:/localhost:8080/goxploitdb/api/v1/type/webapps" | jq
    
    {
    "AUTHOR": "modpr0be (modpr0be[at]spentera.com)",
    "CVE": "CVE-2012-2995, CVE-2012-2996",
    "EDB-ID": "21319",
    "ID": 4,
    "PLATFORM": "AIX",
    "TITLE": "Trend Micro InterScan Messaging Security Suite Stored XSS and CSRF",
    "TYPE": "WEBAPPS",
    "URL": "https://www.exploit-db.com/exploits/21319"
    }
    <------------------------------>SNIP<------------------------------>
    {
    "AUTHOR": "Trent Gordon",
    "CVE": "N/A",
    "EDB-ID": "48026",
    "ID": 2689,
    "PLATFORM": "XML",
    "TITLE": "ExpertGPS 6.38 - XML External Entity Injection",
    "TYPE": "WEBAPPS",
    "URL": "https://www.exploit-db.com/exploits/48026"
    }
    {
    "Total Records": 1951
    }
    ``` 
