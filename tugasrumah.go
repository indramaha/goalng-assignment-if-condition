// no 1. Buatlah program dalam bahasa Go untuk mencari bilangan terbesar dari 3 bilangan bulat.
// package main

// import "fmt"

// func main (){
//     var a, b, c int

//     fmt.Scan(&a, &b, &c)
//     if (a >= b) && (a >= c) {
//         fmt.Println(a)
//     } else if (b >= a) && (b >= c) {
//         fmt.Println(b)
//     } else if (c >= a) && (c >= b) {
//         fmt.Println(c)
//     } else if {
//         fmt.Println("gatau")
//     }
// }

// soal nomor 2
// package main

// import "fmt"

// func main (){
//     var input string

//     fmt.Scan(&input)
//     if (input == "tinggi") {
//         fmt.Println("macet")
//     } else if (input == "rendah") {
//         fmt.Println("tidak macet")
//     } else{
//         fmt.Println("gatau")
//     }
// }

// soal nomor 3
// package main

// import "fmt"

// func main(){
//     var k, p, hasil int

//     fmt.Scan(&k, &p)
//     if ((p%k) == 0){
//         hasil = p/k
//         fmt.Println(hasil)
//     } else if ((p%k) != 0){
//         hasil = (p/k) + 1
//         fmt.Println(hasil)
//     } else {
//         fmt.Println("Not found")
//     }
// }

// sial nomor 4
// package main

// import "fmt"

// func main() {
//     var a, b, c int

//     fmt.Scan(&a, &b, &c)
//     if (a<b && b< c){
//         fmt.Println(c, b, a)
//     } else if (a<c && c< b){
//         fmt.Println(b, c, a)
//     } else if (b<c && c< a){
//         fmt.Println(a, c, b)
//     } else if (c<b && b< a){
//         fmt.Println(a, b, c)
//     } else if (c<a && a< b){
//         fmt.Println(b, a, c)
//     } else if (b<a && a< c){
//         fmt.Println(c, a, b)
//     } else {
//         fmt.Println("Not found")
//     }
// }

// soal nomor 5
// package main

// import "fmt"

// func main (){
//     var input string

//     fmt.Scan(&input)
//     if (input == "A") {
//         fmt.Println(5)
//     } else if (input == "B") {
//         fmt.Println(4)
//     } else if (input == "C") {
//         fmt.Println(3)
//     } else if (input == "D") {
//         fmt.Println(2)
//     } else if (input == "E") {
//         fmt.Println(1)
//     } else if (input == "T") {
//         fmt.Println(0)
//     } else {
//         fmt.Println("input berupa huruf A, B, C, D, E, T")
//     }
// }

// soal nomor 6
// package main

// import "fmt"

// func main (){
//     var juri1, juri2, juri3 string

//     fmt.Scan(&juri1, &juri2, &juri3)
//     if (juri1 == "yes") && (juri2 == "yes") && (juri3 == "yes") {
//         fmt.Println("lolos")
//     } else if ((juri1 == "yes") && (juri2 == "yes") && (juri3 == "no")) || ((juri1 == "yes") && (juri2 == "no") && (juri3 == "yes")) || ((juri1 == "no") && (juri2 == "yes") && (juri3 == "yes")) {
//         fmt.Println("lolos")
//     } else {
//         fmt.Println("tidak lolos")
//     }
// }

// soal nomor 6
// package main

// import "fmt"

// func main (){
//     var inputHours, pCarPrice, pMotorCyclePrice, result int
//     var inputTransport string

//     fmt.Scan(&inputTransport, &inputHours)
//     pCarPrice = 3000
//     pMotorCyclePrice = 2000
//     if (inputTransport == "mobil"){
//         result = pCarPrice + (1000 *(inputHours-1))
//         fmt.Println(result)
//     } else if (inputTransport == "motor"){
//         result = pMotorCyclePrice + (500 *(inputHours-1))
//         fmt.Println(result)
//     } else {
//         fmt.Println("diluar konteks")
//     }
// }