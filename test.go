package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "log"
)

func main() {
    unmarshal()
    marshal()
}

func unmarshal() {
    // 待解析数据
    yamlContent := `
        field1: 小韩说课
        field2:
          field3: value
          field4: [42, 1024]
        `
    // 存储解析数据
    result := make(map[string]interface{})
    // 执行解析
    err := yaml.Unmarshal([]byte(yamlContent), &result)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    fmt.Println(result)
    // result:
    // map[field1:小韩说课 field2:map[field3:value field4:[42 1024]]]
}

func marshal() {
    // map数据
    data := map[string]interface{}{
        "field1": "小韩说课",
        "field2": map[string]interface{}{
            "field3": "value",
            "field4": []int{42, 1024},
        },
    }
    text, err := yaml.Marshal(&data)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    fmt.Println(string(text))
    // result:
    /*
       field1: 小韩说课
       field2:
         field3: value
         field4:
         - 42
         - 1024
    */
}
