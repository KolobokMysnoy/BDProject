# DB_PROJECT
Индивидуальный проект для курса СУБД 2-го семестра Технопарка.
Запросы в большинстве хранятся в app/requests

docker build -t inka .
docker run -p 5000:5000 --name inka -t inka
./technopark-dbms-forum func -u http://localhost:5000/api -r report.html

// заполнение:
./technopark-dbms-forum fill --url=http://localhost:5000/api --timeout=900

// тестирование:
./technopark-dbms-forum perf --url=http://localhost:5000/api --duration=600 --step=60

РПС 1531
