## **Dev Plan**

---

### **project desc**

the cloud server bash on minio.

### **back-end stack**

- **storge stack**
  
  minio

### **api**

- **login**

  >method: POST

  >uri: auth/login

  **request**

  parm | type | desc
  :--: | :--: |:--:
  account | string | Xcloud account
  password | string | Xcloud password

  **response**
  ```json
  {
      "code":200,
      "token":"bear...",
      "msg":"success"
  }
  ```

- get-user-info
  
  >method: GET

  >uri: auth/get-user-info

  **request**

  parm | type | desc
  :--: | :--: :--:
  token | string | user token

  **response**
  ```json
  {
      "code":200,
      "nick":"....",
      "avatar":"....",
      "msg":"success"
  }
  ```

- get-storge-info
  
  >method: GET

  >uri: auth/get-storge-info

  **request**

  parm | type | desc
  :--: | :--: | :--:
  token | string | xlcloud user token

  **response**
  ```json
  {
      "code":200,
      "used":"...",
      "limit":"...",
      "msg":"success"
  }
  ```

- get-server-info

    >method: GET

  >uri: auth/get-server-info

  **request**

  parm | type | desc
  :--: | :--: | :--:
  token | string | xlcloud user token

  **response**
  ```json
  {
      "code":200,
      "data":[
          {"ip":"127.0.0.1","stroge":"1000G",..."},
          {},
          ...
      ],
      "msg":"success"
  }
  ```

- get-source
  
  >method: GET

  >uri: auth/get-source

  **request**

  parm | type | desc
  :--: | :--: | :--:
  token | string | xlcloud user token
  type | string | the type of get source (download, view)
  target | []string | source target

  **response**
  ```json
  {
      "code":200,
      "data":[
          ...
      ],
      "msg":"success"
  }
  ```

- upload-source
  
  >method: POST

  >uri: auth/upload-source

  **request**

  parm | type | desc
  :--: | :--: | :--:
  token | string | xlcloud user token
  target | []string | upload source

  **response**
  ```json
  {
      "code":200,
      "msg":"success"
  }
  ```

- delete-source
  
  >method: DELETE

  >uri: auth/delete-source

  **request**

  parm | type | desc
  :--: | :--: | :--:
  token | string | xlcloud user token
  target | []string | upload source

  **response**
  ```json
  {
      "code":200,
      "msg":"success"
  }
  ```