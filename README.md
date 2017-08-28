Данный репозиторий представляет собой дополнительный уровень абстракции
между конкретными реализациями парсеров для разных устройств:
https://github.com/doctor-fate/mskix-parsers 
и самой программой для рисования диаграмм: 
https://github.com/doctor-fate/mskix-drawer

Пакет `mskix` определяет интерфейс:
```go
type Parser interface {
    Parse([]byte) (device.Data, error)
}
```
Данный интерфейс позволяет взаимодействовать клиентскому коду и 
конкретным реализациям парсеров между собой. 

Минимальный пример использования данного пакета:
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/doctor-fate/mskix"
    _ "github.com/doctor-fate/mskix-parsers"
)

func main() {
    input := /* acquire input bytes... */
    
    data, err := mskix.Parse(input)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(data)
}
    
```

Как видно из примера клиентский код ничего не знает про конкретные реализации парсеров. Также стоит отметить пустой импорт `_ "github.com/doctor-fate/mskix-parsers"`. Он позволяет различным парсерам из пакета `github.com/doctor-fate/mskix-parsers` 
зарегестрировать себя, для того чтобы они стали доступны для использования в функции `Parse`.

API данного репозитория должно гарантировать обратную совместимость, для того, чтобы клиентский код продолжал работать без изменений.