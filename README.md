# pgxCopy

CopyDatabase는 PostgreSQL 데이터베이스와 상호작용하는 Go 라이브러리로, 구조체나 맵에서 데이터를 빠르고 효율적으로 COPY 방식으로 데이터베이스에 삽입할 수 있는 기능을 제공합니다. 이 라이브러리는 ORM처럼 작동하며, 데이터베이스에 데이터를 삽입할 때 매우 빠른 성능을 자랑합니다. 또한, 필드 태그를 사용하여 데이터를 추출하고 삽입할 수 있어 구조체의 필드를 유연하게 다룰 수 있습니다.