func main() {

m := map[string]int{

"James":           32,
"Miss Moneypenny": 27,
}
fmt.Println(m)
fmt.Println(m["James"])
fmt.Println(m["Arun"])

v, ok := m["Arun"]
fmt.Println(v)
fmt.Println(ok)

m["todd"] = 33
fmt.Println(m)

// range from map
for k, v := range m {
fmt.Println(k, v)
}

v1 := []int{1, 2, 3}
fmt.Println(v1)

// range from slice
for k, v := range v1 {
fmt.Println(k, v)
}

}

