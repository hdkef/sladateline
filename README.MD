## launch program

> $ go run main.go

## explanation

1. inisiasi variabel datePointer, hourOfWork, dateLine50, dateLine75, dateLine100

2. datePointer adalah pointer terhadap date yang telah diubah ubah oleh loop,
hourOfWork menampung increment dari jam kerja pada loop, dateLine50 adalah jumlah jam kerja minimal yang harus dicapai hingga 50%, dateLine75 adalah jumlah jam kerja minimal yang harus dicapai hingga 75%, dateLine100 adalah jumlah jam kerja untuk 100%

3. pada loop pertama dengan kondisi hourOfWork < hour SLA,ada kondisi cek apakah weekday atau weekend

4. jika weekend, dan hari sabtu, maka date pointer ditambah 2 hari, jika weekend dan hari minggu maka date pointer ditambah 1 hari

5. jika weekday, pertama cek dulu jam apakah di bawah 09.00 ? jika iya set jam dari datePointer ke jam 09.00

6. kemudian terdapat loop kedua dengan 8 kali loop sesuai jam kerja per hari (8 jam (09.00 - 18.00 dan istirahat 1 jam)) 

7. pada loop tersebut akan dicek :
 jika datePointer jam 12, maka datePointer ditambah 1 jam dan break loop
 jika datePointer jam > 18, maka break loop

8. kemudian, hourOfWork dan datePointer akan diincrement 1 jam

9. Tiap selesai increment, ada switch case untuk cek apakah hourOfWork telah sama dengan dateLine50, dateLine75, ataupun dateLine100

10. Jika sama, maka append ke array untuk mendapatkan date per percentage nya

11. Selesai loop, return array