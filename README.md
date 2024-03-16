POST /create-news (создание новости) :
{
  "date" : "",
    "name" : "",
    "small_description" : "",
    "full_description" : "",
    "image_path" : ""
}
успешное создание:
   "id": ,
  "status": "success"
неккоректная дата:
 "error": "pq: invalid input syntax for type date: \"\""
пустое поле name или small_description или full_description:
{
  error:"invalid data"
}
 
DELETE /delete-news (удаление новости) :
  {
    "id":2
  }
  успешное удаление:
    "status": "delete success"
  нет записей с таким id:
    "error": "sql: no rows in result set"
    
POST /get-newsById (получить запись по id):
  {
    "id":1
  }
  успешное получение:
   "news": {
        "id":,
        "date": "",
        "name": "",
        "small_description": "",
        "full_description": "",
        "image_path": ""
    }
    нет записей с таким id:
    "error": "record not found"
    
  GET /get-allNews(получить список всех новостей):
    успешное получение:
     "news": 
        {
            "id": 8,
            "date": "2024-01-01T00:00:00Z",
            "name": "",
            "small_description": "",
            "full_description": "",
            "image_path": "на серваке"
        },
        {
            "id": 9,
            "date": "2024-01-01T00:00:00Z",
            "name": "",
            "small_description": "",
            "full_description": "",
            "image_path": "на серваке"
        }
    
      нет новостей:
       "news": []
       
  POST /update-news (изменить новость):
   "id":,
    "date" : "",
    "name" : "",
    "small_description" : "",
    "full_description" : "",
    "image_path" : ""

    если успешное изменение новости: 
     "status": "success"

    PUT /upload-photo(загрузить фото):
    key newsPhoto и само фото прикрепляешь
    ответ:
      "file path": "static/news/uoGOH9f.png",
      "status": "successfully create"
      
    DELETE /delete-photo:
    {
    "name":""
    }
    успешное удаление фото:
      "file path": "./static/news/uoGOH9f.png",
      "status": "successfully delete"
    если нет фото:
     "error": "remove ./static/news/uoGOH9f.png: no such file or directory"
      
     
       
      
  
    
