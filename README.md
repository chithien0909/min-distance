# min-distance

## MinDistance1
1. Một vòng lặp qua các phần tử của biến "arr" để lưu lại vị trí các chuỗi trong mảng
2. Kiểm tra xem phần tử a và b có tồn tại trong mảng không, nếu không tồn tại sẽ trả về -1
3. Gán minDist là số lượng phần tử trong mảng (vị trí lớn nhất có thể)
4. Lặp qua danh sách vị trí của của phần tử a và b
   - Tính khoảng cách từng phần vị trí của a so với b
   - So sánh khoảng cách với minDist nếu nhỏ hơn sẽ được gán lại cho biến minDist
5. Trả về minDist


## MinDistance2 

1. Một vòng lặp qua các phần tử của biến "arr" để đế đếm số lượng xuất hiện của từng phần tử trong mảng
2. Kiểm tra xem phần tử a và b có tồn tại trong mảng không, nếu không tồn tại sẽ trả về -1
3. Khởi tạo minDist là số lượng phần tử trong mảng (vị trí lớn nhất có thể)
4. Khởi tạo aIndex và bIndex là vị trí đầu tiên xuất hiện trong mảng
5. Vòng lặp for cho đến khi aIndex, bIndex đến vị trí cuối cùng trong mảng
   - 5.1 Nếu (phần tử thứ aIndex = a và bIndex  = b) hoặc (phần tử thứ aIndex = b và bIndex  = a) thì sẽ so sánh khoảng cách giữa a và b với minDist
   nếu nhỏ hơn sẽ gán cho minDist. Tăng aIndex và bIndex thêm 1 Thoả điều kiện sẽ quay lại bước 5 ngược lại sẽ đi sang 5.2
   - 5.2 Kiểm tra xem tại vị trí bIndex có xuất hiện "a" cần tìm không, nếu có sẽ gán aIndex  = bIndex (Giải quyết trường hợp "a" xuất hiện nhiều lần so với a tại vị trí trước đó )
   - 5.3 Kiểm tra xem phần tử tại aIndex có bằng "a" hoặc "b" hay không, nếu có sẽ tăng bIndex thêm 1 và thực hiện bước 5, ngược lại sẽ sang 5.4
   - 5.4 Kiểm tra xem phần tử tại bIndex có bằng "b" hay không nếu bằng thì sẽ tăng aIndex thêm 1 và thực hiện bước 5, ngược lại sẽ sang 5.5
   - 5.5 Khi tại vị trí aIndex và bIndex không tìm thấy vị trí nào bằng "a" hoặc "b" sẽ tăng aIndex và bIndex lên 1 đơn vị
    => Giống cách quét cạn kiểm tra hết lần lượt từng phần tử.
6. Đảo vị trí tìm a, b = b, a. Do một số test case thì b sẽ đúng trước a, nên cần đổi chỗ 2 giá trị tìm theo chiều ngược lại.
   Có thể đảo ngược lại arr nhưng sẽ tốn nhiều chi phí hơn khi reverse 
7. So sánh giá trị của 2 lần gọi GetMin trả về minDist nhỏ nhất.
=> Có thể set thêm Goroutines để chạy đồng thời GetMin
