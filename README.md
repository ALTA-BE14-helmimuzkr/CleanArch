# Unit Test
Jalankan perintah ``` sh unit-test.sh ``` untuk melakukan unit test

![Coverage test image](https://github.com/ALTA-BE14-helmimuzkr/CleanArch/blob/main/cover-screenshot/image_2023-01-12_23-24-12.png)

# API Documentation 
## URL
API URL http://13.213.44.165
## USER
### Register
```
  POST /register
```
Body Request - JSON
```
{
	"nama": string, 
	"email": string,  
	"password": string, 
	"alamat": string,
	"hp": string
}
```
### Login
```
  POST /login
```
Body Request - JSON
```
    {
        "email": string,
        "password": string,
    }
```
### Profile
**Required** Token Bearer from login
```
  GET /users/profile
```
### Update
**Required** Token Bearer from login
```
  PUT /users
```
Body Request - JSON
```
{
	"nama": string,
	"email": string,
	"password": string,
	"alamat": string,
	"hp": string
}
```
### Deactive
**Required** Token Bearer from login
```
  DELETE /users
```
## Book
### Get all book
```
  GET /books
```
### Get my book
**Required** Token Bearer from login
````
  GET /books/mybook
````
Body Request - JSON
```
{
		"judul": string,
		"tahun_terbit": number,
		"penulis": string
}
```
### Add
**Required** Token Bearer from login
```
  POST /books
```
Body Request - JSON
```
{
		"judul": string,
		"tahun_terbit": number,
		"penulis": string
}
```
### Update
**Required** Token Bearer from login
```
  PUT /books/:id
```
Body Request - JSON
```
{
		"judul": string,
		"tahun_terbit": number,
		"penulis": string
}
```
| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------  |
| `id`      | `string` | **Required**. Id of book to update |

### Delete
**Required** Token Bearer from login
```
  DELETE /books/:id
```
| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------  |
| `id`      | `string` | **Required**. Id of book to delete |
