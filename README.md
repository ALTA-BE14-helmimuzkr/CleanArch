# Unit Test
Jalankan perintah ``` sh unit-test.sh ``` untuk melakukan unit test

![Coverage test image](https://github.com/ALTA-BE14-helmimuzkr/CleanArch/blob/main/cover-screenshot/image_2023-01-12_23-24-12.png)

# API Documentation 
## URL
API URL http://13.213.44.165
## USER
### Register
```http
  POST /register
```
Body Request - JSON
```json
{
	"nama": string, 
	"email": string,  
	"password": string, 
	"alamat": string,
	"hp": string
}
```
### Login
```http
  POST /login
```
Body Request - JSON
```json
    {
        "email": string,
        "password": string,
    }
```
### Profile
**Required** Token Bearer from login
```http
  GET /users/profile
```
### Update
**Required** Token Bearer from login
```http
  PUT /users
```
Body Request - JSON
```json
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
```http
  DELETE /users
```
## Book
### Get all book
```http
  GET /books
```
### Get my book
**Required** Token Bearer from login
```http
  GET /books/mybook
```
Body Request - JSON
```json
{
		"judul": string,
		"tahun_terbit": number,
		"penulis": string
}
```
### Add
**Required** Token Bearer from login
```http
  POST /books
```
Body Request - JSON
```json
{
		"judul": string,
		"tahun_terbit": number,
		"penulis": string
}
```
### Update
**Required** Token Bearer from login
```http
  PUT /books/:id
```
Body Request - JSON
```json
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
```http
  DELETE /books/:id
```
| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------  |
| `id`      | `string` | **Required**. Id of book to delete |