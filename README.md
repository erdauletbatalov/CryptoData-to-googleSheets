# This is job interview task

Google Spreadsheet with parsed data (everyone-reedAccess): https://docs.google.com/spreadsheets/d/1mqBbMNEonPIPozZ9EA-kzzon81uXP9MKsmEzim0URHA/edit?usp=sharing

Just run ```go run main.go``` to fill the table with info  
You can choose ```how many rows``` you want to Read from each API and ```where to write``` into the table  

## Парсинг Cryptorank:  
Ресурс: [```Cryptorank```](https://cryptorank.io/price/bitcoin)  
Данные для парсинга: ```Теги нескольких валют (первых трех)```  

Метод парсинга: ```Любой```  
Метод хранения полученных результатов: ```Запись в гугл таблицы (api) по запуску Количество столбцов 3: Наименование, Теги, Timestamp.```  
Время на выполнение: ```Решает исполнитель```  



## Парсинг CoinGecko  
Ресурс [```Coingecko```](https://www.coingecko.com/en)  
Данные для парсинга: ```Валюты, их стоимость относительно доллара```  
Метод парсинга: ```Любой```  
Метод хранения полученных результатов: ```Запись в гугл таблицы(api) по запуску. Количество столбцов: Наименование, Цена, Timestamp. (Должно выводиться за один запрос 65! валют, вместе с ценами).```  
Время на выполнение: ```Решает исполнитель```  
