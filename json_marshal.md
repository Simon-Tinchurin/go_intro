# JSON Marshal/Unmarshal

```
import "encoding/json"
```

Marshal - сериализация, возвращает json объект
```
func main() {
    v:=map[string]interface{}{
        "field1": "some str",
        "field2": true,
        "field3": 123,
        }
    json.Marshal(sv)
}
```
Unmarshal - десериализация
```
func main() {
    b:=[]byte(`{"field1": "some str", "field2": true, "field3": 123}`)
    var data map[string]interface{}
    if err := json.Unmarshal(byt, &data); err != nil {
        panic(err)
    }
}
```