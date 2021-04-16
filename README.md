# latihan go

## 1. Soal looping

```js
soal berada di looping.go
```

Buatlah sebuah looping dari 1 - 100, jika angkanya dapat dibagi dengan 2 tampilkan "Buzz", jika angkanya dapat dibagi dengan 3 tampilkan "Fuzz", jika angkanya dapat dibagi dengan 5 maka tampilkan "BuzzFuzz", sisanya jangan ditampilkan

## 2. Simple Calculator

```js
soal berada di simpleCalc.go
```

Mari membuat calculator sederhana ...
Jadi kamu disuruh membuat kalkulator yang dapat menampung 2 angka dengan tipe data float32 , dan dapat menerima operasi + (penambahan), - (pengurangan), / (pembagian) , \* (perkalian), % (modulus). Kalian harus memunculkan hasil eksekusi sesuai dengan operasi yang dimasukkan, ketika operasi yang dimasukkan salah , tampilkan "error operasi yang dimasukkan salah!"

## 3. create Password

```js
soal berada di createPassword.go
```

Buatlah sebuah eksekusi code dimana dapat mengubah nama kalian menjadi password. Password yang dibuat akan otomatis dengan step sebagai berikut :

1. pertama lakukan pengecekan, nama yang dikonversi menjadi password harus minimal 5 karakter, jika kurang kembalikan error
2. Kedua, Balik karakter dari urutan paling belakang menjadi urutan paling depan
3. Ketiga, tambah karakter pertama dengan huruf Besar dari nama awal yang diinput, dan tmabah karakter terakhir dengan huruf besar dari nama akhir yang diinput.
4. Terakhir, tambahkan karakter di belakang sendiri dengan panjang dari karakter yang diinput

```js
NB :
1. teman - teman bisa menggunakan strings.ToUpper(string) untuk mengubah dari huruf kecil ke huruf besar, dan strings.ToLower(string) jika huruf kecil
2. Kalian juga bisa menggunakan package strconv.Itoa(int) untuk mengkonversi dari int ke string
```
